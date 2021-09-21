package handlers

import (
	"log"
	"tree-web-server/controllers"
	"tree-web-server/models"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	// response struct
	resp := models.ResponseMessage{Status: 200, Message: "OK"}
	return c.Status(200).JSON(resp)
}

func Login(c *fiber.Ctx) error {
	user := new(models.UserAuth)

	if err := c.BodyParser(user); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("username: ", user.Username)
	log.Println("password: ", user.Password)

	// query
	res, err := controllers.FindUser(user.Username)
	if err != nil { // err = record not found
		return c.Status(404).SendString(err.Error())
	}

	err = controllers.ComparePassword(user.Password, []byte(res.Password))
	if err != nil {
		return c.Status(404).SendString("invalid username or password")
	}
	return c.Status(200).SendString(res.Uuid)
}

func GetTreeList(c *fiber.Ctx) error {
	// get tree list form db
	// t := []string{"tree1", "tree2", "tree3", "tree4"}
	// tree := models.Item{Tree: t}
	uuid := new(models.Uuid)

	if err := c.BodyParser(uuid); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("uuid: ", uuid.Uuid)

	// query
	res, err := controllers.FindTree(uuid.Uuid)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	var resJSON models.Item
	for _, v := range res {
		resJSON.Tree = append(resJSON.Tree, v.TreeName)
	}

	return c.Status(200).JSON(resJSON)
}

func AddUser(c *fiber.Ctx) error {
	newuser := new(models.ManageUser)

	if err := c.BodyParser(newuser); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("uuid: ", newuser.Uuid)

	if newuser.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	password, err := controllers.HashPassword(newuser.Password)
	if err != nil {
		log.Println(err)
	}
	user := models.User{Username: newuser.Username, Password: string(password), Uuid: controllers.GenUUID()}
	res := controllers.Insert("users", &user)
	if res.Error != nil {
		log.Println("insert err:", res.Error.Error())
		return c.Status(500).SendString(res.Error.Error())
	} else if res.RowsAffected != 1 {
		log.Println("RowsAffected:", res.RowsAffected)
		return c.Status(500).SendString(res.Error.Error())
	}

	return c.Status(200).SendString("user: " + newuser.Username + " is created")
}

func UpdateUser(c *fiber.Ctx) error {
	updateUser := new(models.ManageUser)

	if err := c.BodyParser(updateUser); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("uuid: ", updateUser.Uuid)

	if updateUser.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	newPassword, err := controllers.HashPassword(updateUser.Password)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString(err.Error())
	}

	user := models.User{Password: string(newPassword)}
	res := controllers.Update("users", "username", updateUser.Username, user)
	if res.Error != nil {
		log.Println("updateUser err:", res.Error.Error())
		return c.SendStatus(500)
	} else if res.RowsAffected != 1 {
		log.Println("RowsAffected:", res.RowsAffected)
		return c.SendStatus(500)
	}

	return c.Status(200).SendString("user: " + updateUser.Username + " is updated")
}

func DeleteUser(c *fiber.Ctx) error {
	deleteUser := new(models.ManageUser)

	if err := c.BodyParser(deleteUser); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("uuid: ", deleteUser.Uuid)

	if deleteUser.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	var user models.User
	res := controllers.Delete("users", "username", deleteUser.Username, &user)
	if res.Error != nil {
		log.Println("deleteUser err:", res.Error.Error())
		return c.Status(500).SendString(res.Error.Error())
	} else if res.RowsAffected != 1 {
		log.Println("RowsAffected:", res.RowsAffected)
		return c.Status(500).SendString(res.Error.Error())
	}

	return c.Status(200).SendString("user: " + deleteUser.Username + " is deleted")
}
