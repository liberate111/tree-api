package models

type User struct {
	Username string `gorm:"primaryKey;autoIncrement:false"`
	Password string
	Uuid     string `gorm:"unique"`
}

type Tree struct {
	TreeName  string `gorm:"primaryKey;autoIncrement:false"`
	Owner     string
	Level     int
	State     string
	StartTime int
	StopTime  int
}
