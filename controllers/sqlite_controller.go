package controllers

import (
	"fmt"
	"log"
	"tree-web-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteClient struct {
	Con *gorm.DB
}

var (
	DB *SqliteClient
)

func ConnectSqlite(dbName string) (*SqliteClient, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	AutoMigrate(db)
	client := SqliteClient{Con: db}
	return &client, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.Table("users").AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	err = db.Table("trees").AutoMigrate(&models.Tree{})
	if err != nil {
		return err
	}
	return nil
}

func FindAllUser() ([]models.User, error) {
	var userSqlite []models.User
	res := DB.Con.Table("users").Find(&userSqlite)
	if res.Error != nil {
		return userSqlite, fmt.Errorf(res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return userSqlite, fmt.Errorf("data not found")
	}
	return userSqlite, nil
}

func FindAllTree() ([]models.Tree, error) {
	var trees []models.Tree
	res := DB.Con.Table("trees").Find(&trees)
	if res.Error != nil {
		return trees, res.Error
	}
	if res.RowsAffected == 0 {
		return trees, fmt.Errorf("data not found")
	}
	return trees, nil
}

func FindUser(username string) (models.User, error) {
	var userSqlite models.User
	res := DB.Con.Table("users").Where("username = ?", username).First(&userSqlite)
	if res.Error != nil {
		return userSqlite, fmt.Errorf("query user err: %s", res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return userSqlite, fmt.Errorf("user: %s not found", username)
	}
	return userSqlite, nil
}

func FindTree(field string, query string) ([]models.Tree, error) {
	var trees []models.Tree
	con := fmt.Sprintf("%s = ?", field)
	res := DB.Con.Table("trees").Where(con, query).Find(&trees)
	if res.Error != nil {
		return trees, res.Error
	}
	if res.RowsAffected == 0 {
		return trees, fmt.Errorf("tree not found")
	}
	return trees, nil
}

func Insert(tableName string, data interface{}) *gorm.DB {
	res := DB.Con.Table(tableName).Create(data)
	return res
}

func Update(tableName string, fieldPrimaryKey string, primaryKey string, data interface{}) *gorm.DB {
	con := fmt.Sprintf("%s = ?", fieldPrimaryKey)
	res := DB.Con.Table(tableName).Where(con, primaryKey).Updates(data)
	return res
}

func Delete(tableName string, fieldPrimaryKey string, primaryKey string, data interface{}) *gorm.DB {
	con := fmt.Sprintf("%s = ?", fieldPrimaryKey)
	res := DB.Con.Table(tableName).Where(con, primaryKey).Delete(data)
	return res
}

func ConnectDB() {
	var err error
	DB, err = ConnectSqlite("tree-wal.db")
	if err != nil {
		log.Panicln("connect to sqlite err:", err.Error())
	}
	// CreateUserTest()
}

// func CreateUserTest() {
// 	// for test
// 	password, err := HashPassword("password")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// table users
// 	// test insert
// 	user := models.User{Username: "moo", Password: string(password), Uuid: GenUUID()}
// 	res := Insert("users", &user)
// 	if res.Error != nil {
// 		log.Println("insert err:", res.Error.Error())
// 	} else if res.RowsAffected != 1 {
// 		log.Println("RowsAffected:", res.RowsAffected)
// 	}

// 	// test query
// 	// userTest, err := FindUser("moo")
// 	// if err != nil {
// 	// 	log.Println("query err:", res.Error.Error())
// 	// }
// 	// uuid := userTest.Uuid
// 	// log.Println("uuid:", uuid)

// 	//table trees
// 	tree := []models.Tree{
// 		// {TreeName: "Tree0001", Owner: uuid, Level: 1, State: "dry", StartTime: 1632335714, StopTime: 1632422114},
// 		// {TreeName: "Tree0002", Owner: uuid, Level: 2, State: "dry", StartTime: 1632335714, StopTime: 1632422114},
// 		// {TreeName: "Tree0003", Owner: uuid, Level: 3, State: "dry", StartTime: 1632335714, StopTime: 1632422114},
// 		// {TreeName: "Tree0004", Owner: uuid, Level: 4, State: "wet", StartTime: 1632335714, StopTime: 1632422114},
// 		// {TreeName: "Tree0005", Owner: uuid, Level: 5, State: "grow", StartTime: 1632335714, StopTime: 1632422114},
// 		{TreeName: "Tree0001"},
// 		{TreeName: "Tree0002"},
// 		{TreeName: "Tree0003"},
// 		{TreeName: "Tree0004"},
// 		{TreeName: "Tree0005"},
// 		{TreeName: "Tree0006"},
// 		{TreeName: "Tree0007"},
// 		{TreeName: "Tree0008"},
// 		{TreeName: "Tree0009"},
// 		{TreeName: "Tree0010"},
// 	}

// 	for _, v := range tree {
// 		res := Insert("trees", &v)
// 		if res.Error != nil {
// 			log.Println("insert err:", res.Error.Error())
// 		} else if res.RowsAffected != 1 {
// 			log.Println("RowsAffected:", res.RowsAffected)
// 		}
// 	}
// }

// func CreateUserAdmin() {
// 	password, err := HashPassword("9how,hlug-up;") // ต้นไม้สีเขียว
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	user := models.User{Username: "admin", Password: string(password), Uuid: "uuid-9how,hlug-up;"}
// 	res := Insert("users", &user)
// 	if res.Error != nil {
// 		log.Println("insert err:", res.Error.Error())
// 	} else if res.RowsAffected != 1 {
// 		log.Println("RowsAffected:", res.RowsAffected)
// 	}
// }
