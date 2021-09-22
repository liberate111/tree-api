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
		return c.Status(400).SendString(err.Error())
	}

	log.Println("username: ", user.Username)
	log.Println("password: ", user.Password)

	// query
	res, err := controllers.FindUser(user.Username)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = controllers.ComparePassword(user.Password, []byte(res.Password))
	if err != nil {
		return c.Status(400).SendString("invalid username or password")
	}
	return c.Status(200).SendString(res.Uuid)
}

func ChangePassword(c *fiber.Ctx) error {
	user := new(models.ChangePassword)

	if err := c.BodyParser(user); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(400).SendString(err.Error())
	}

	// query
	resQuery, err := controllers.FindUser(user.Username)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = controllers.ComparePassword(user.OldPassword, []byte(resQuery.Password))
	if err != nil {
		return c.Status(400).SendString("invalid username or password")
	}

	newPassword, err := controllers.HashPassword(user.NewPassword)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	updateUser := models.User{Password: string(newPassword)}
	resUpdate := controllers.Update("users", "username", user.Username, updateUser)
	if resUpdate.Error != nil {
		log.Println("updateUser err:", resUpdate.Error.Error())
		return c.SendStatus(500)
	} else if resUpdate.RowsAffected != 1 {
		log.Println("RowsAffected:", resUpdate.RowsAffected)
		return c.SendStatus(500)
	}

	return c.Status(200).SendString("user: " + updateUser.Username + " is updated")
}
