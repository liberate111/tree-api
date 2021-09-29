package handlers

import (
	"fmt"
	"tree-web-server/controllers"
	"tree-web-server/models"

	"github.com/gofiber/fiber/v2"
)

func FindAllUser(c *fiber.Ctx) error {
	user := new(models.ManageUser)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(user); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if user.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	// query
	res, err := controllers.FindAllUser()
	if err != nil {
		resJSON = models.ResponseMessage{Status: 404, Message: err.Error()}
		return c.Status(404).JSON(resJSON)
	}
	var userList []models.User
	for _, v := range res {
		var user models.User
		user.Username = v.Username
		user.Uuid = v.Uuid
		userList = append(userList, user)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Users: userList}}
	return c.Status(200).JSON(resJSON)
}

func FindUser(c *fiber.Ctx) error {
	user := new(models.ManageUser)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(user); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if user.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	// query
	res, err := controllers.FindUser(user.Username)
	if err != nil {
		resJSON = models.ResponseMessage{Status: 404, Message: err.Error()}
		return c.Status(404).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Uuid: res.Uuid}}
	return c.Status(200).JSON(resJSON)
}

func AddUser(c *fiber.Ctx) error {
	newuser := new(models.ManageUser)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(newuser); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if newuser.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	password, err := controllers.HashPassword(newuser.Password)
	if err != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: err.Error()}
		return c.Status(500).JSON(resJSON)
	}
	user := models.User{Username: newuser.Username, Password: string(password), Uuid: controllers.GenUUID()}
	res := controllers.Insert("users", &user)
	if res.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: res.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if res.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", res.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 201, Message: "success"}
	return c.Status(201).JSON(resJSON)
}

func UpdateUser(c *fiber.Ctx) error {
	updateUser := new(models.ManageUser)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(updateUser); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if updateUser.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	newPassword, err := controllers.HashPassword(updateUser.Password)
	if err != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: err.Error()}
		return c.Status(500).JSON(resJSON)
	}

	user := models.User{Password: string(newPassword)}
	res := controllers.Update("users", "username", updateUser.Username, user)
	if res.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: res.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if res.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", res.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}
	resJSON = models.ResponseMessage{Status: 200, Message: "success"}
	return c.Status(200).JSON(resJSON)
}

func DeleteUser(c *fiber.Ctx) error {
	deleteUser := new(models.ManageUser)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(deleteUser); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if deleteUser.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	var user models.User
	res := controllers.Delete("users", "username", deleteUser.Username, &user)
	if res.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: res.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if res.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", res.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success"}
	return c.Status(200).JSON(resJSON)
}

func FindAllTree(c *fiber.Ctx) error {
	user := new(models.ManageTree)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(user); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if user.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	// query
	res, err := controllers.FindAllTree()
	if err != nil {
		resJSON = models.ResponseMessage{
			Status:  404,
			Message: err.Error(),
		}
		return c.Status(404).JSON(resJSON)
	}

	var treeList []models.Tree
	for _, v := range res {
		var tree models.Tree
		tree.TreeName = v.TreeName
		tree.Owner = v.Owner
		tree.Level = v.Level
		tree.State = v.State
		tree.StartTime = v.StartTime
		tree.StopTime = v.StopTime
		treeList = append(treeList, tree)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Trees: treeList}}
	return c.Status(200).JSON(resJSON)
}

func FindTree(c *fiber.Ctx) error {
	tree := new(models.ManageTree)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(tree); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if tree.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	// query
	res, err := controllers.FindTree("owner", tree.Owner)
	if err != nil {
		resJSON = models.ResponseMessage{
			Status:  404,
			Message: err.Error(),
		}
		return c.Status(404).JSON(resJSON)
	}
	var treeList []models.Tree
	for _, v := range res {
		var tree models.Tree
		tree.TreeName = v.TreeName
		tree.Owner = v.Owner
		tree.Level = v.Level
		tree.State = v.State
		tree.StartTime = v.StartTime
		tree.StopTime = v.StopTime
		treeList = append(treeList, tree)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Trees: treeList}}
	return c.Status(200).JSON(resJSON)
}

func AddTree(c *fiber.Ctx) error {
	newTree := new(models.ManageTree)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(newTree); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if newTree.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	tree := models.Tree{TreeName: newTree.TreeName, Owner: newTree.Owner}
	res := controllers.Insert("trees", &tree)
	if res.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: res.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if res.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", res.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 201, Message: "success"}
	return c.Status(201).JSON(resJSON)
}

func UpdateTree(c *fiber.Ctx) error {
	updateTree := new(models.ManageTree)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(updateTree); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if updateTree.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	tree := models.Tree{Owner: updateTree.Owner, Level: updateTree.Level, State: updateTree.State, StartTime: updateTree.StartTime, StopTime: updateTree.StopTime}
	res := controllers.Update("trees", "tree_name", updateTree.TreeName, tree)
	if res.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: res.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if res.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", res.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success"}
	return c.Status(200).JSON(resJSON)
}

func DeleteTree(c *fiber.Ctx) error {
	deleteTree := new(models.ManageTree)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(deleteTree); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if deleteTree.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	var tree models.Tree
	res := controllers.Delete("trees", "tree_name", deleteTree.TreeName, &tree)
	if res.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: res.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if res.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", res.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success"}
	return c.Status(200).JSON(resJSON)
}

func Transfer(c *fiber.Ctx) error {
	newOwner := new(models.ManageTree)
	var resJSON models.ResponseMessage
	if err := c.BodyParser(newOwner); err != nil {
		resJSON = models.ResponseMessage{Status: 400, Message: err.Error()}
		return c.Status(400).JSON(resJSON)
	}

	if newOwner.Uuid != "uuid-9how,hlug-up;" {
		resJSON = models.ResponseMessage{Status: 403, Message: "Forbidden"}
		return c.Status(403).JSON(resJSON)
	}

	// query
	resQuery, err := controllers.FindUser(newOwner.Username)
	if err != nil { // err = record not found
		resJSON = models.ResponseMessage{Status: 404, Message: err.Error()}
		return c.Status(404).JSON(resJSON)
	}

	tree := models.Tree{Owner: resQuery.Uuid}
	resUpdate := controllers.Update("trees", "tree_name", newOwner.TreeName, tree)
	if resUpdate.Error != nil {
		resJSON = models.ResponseMessage{Status: 500, Message: resUpdate.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if resUpdate.RowsAffected != 1 {
		resJSON = models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", resUpdate.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}

	resJSON = models.ResponseMessage{Status: 200, Message: "success"}
	return c.Status(200).JSON(resJSON)
}
