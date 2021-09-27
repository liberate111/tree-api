package handlers

import (
	"fmt"
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
	var resJSON models.ResponseMessage

	if err := c.BodyParser(user); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	// query
	res, err := controllers.FindUser(user.Username)
	if err != nil {
		resJSON = models.ResponseMessage{Status: 404, Message: err.Error()}
		return c.Status(404).JSON(resJSON)
	}

	err = controllers.ComparePassword(user.Password, []byte(res.Password))
	if err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: "invalid username or password"}
		return c.Status(400).JSON(resJSON)
	}
	resJSON = models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Uuid: res.Uuid}}
	return c.Status(200).JSON(resJSON)
}

func ChangePassword(c *fiber.Ctx) error {
	user := new(models.ChangePassword)
	var resJSON models.ResponseMessage

	if err := c.BodyParser(user); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	// query
	resQuery, err := controllers.FindUser(user.Username)
	if err != nil {
		resJSON = models.ResponseMessage{Status: 404, Message: err.Error()}
		return c.Status(404).JSON(resJSON)
	}

	err = controllers.ComparePassword(user.OldPassword, []byte(resQuery.Password))
	if err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: "invalid username or password"}
		return c.Status(400).JSON(resJSON)
	}

	newPassword, err := controllers.HashPassword(user.NewPassword)
	if err != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: err.Error()}
		return c.Status(500).JSON(resJSON)
	}

	updateUser := models.User{Password: string(newPassword)}
	resUpdate := controllers.Update("users", "username", user.Username, updateUser)
	if resUpdate.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: resUpdate.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if resUpdate.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 500, Message: fmt.Sprintf("RowsAffected: %d", resUpdate.RowsAffected)}
		return c.Status(500).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success"}
	return c.Status(200).JSON(resJSON)
}
