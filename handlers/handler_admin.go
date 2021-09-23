package handlers

import (
	"log"
	"tree-web-server/controllers"
	"tree-web-server/models"

	"github.com/gofiber/fiber/v2"
)

func FindUser(c *fiber.Ctx) error {
	user := new(models.ManageUser)

	if err := c.BodyParser(user); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(400).SendString(err.Error())
	}

	log.Println("uuid: ", user.Uuid)

	if user.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	// query
	res, err := controllers.FindUser(user.Username)
	if err != nil { // err = record not found
		return c.Status(404).SendString(err.Error())
	}

	return c.Status(200).SendString(res.Uuid)
}

func AddUser(c *fiber.Ctx) error {
	newuser := new(models.ManageUser)

	if err := c.BodyParser(newuser); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(400).SendString(err.Error())
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

func FindTree(c *fiber.Ctx) error {
	tree := new(models.ManageTree)

	if err := c.BodyParser(tree); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("uuid: ", tree.Uuid)

	if tree.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	// query
	res, err := controllers.FindTree(tree.Owner)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	var resJSON models.Item
	for _, v := range res {
		resJSON.Tree = append(resJSON.Tree, v.TreeName)
	}

	return c.Status(200).JSON(resJSON)
}

func AddTree(c *fiber.Ctx) error {
	newTree := new(models.ManageTree)

	if err := c.BodyParser(newTree); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(400).SendString(err.Error())
	}

	log.Println("uuid: ", newTree.Uuid)

	if newTree.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	tree := models.Tree{TreeName: newTree.TreeName, Owner: newTree.Owner}
	res := controllers.Insert("trees", &tree)
	if res.Error != nil {
		log.Println("insert err:", res.Error.Error())
		return c.Status(500).SendString(res.Error.Error())
	} else if res.RowsAffected != 1 {
		log.Println("RowsAffected:", res.RowsAffected)
		return c.Status(500).SendString(res.Error.Error())
	}

	return c.Status(200).SendString("tree: " + newTree.TreeName + " is created")
}

func UpdateTree(c *fiber.Ctx) error {
	updateTree := new(models.ManageTree)

	if err := c.BodyParser(updateTree); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("uuid: ", updateTree.Uuid)

	if updateTree.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	tree := models.Tree{Owner: updateTree.Owner}
	res := controllers.Update("trees", "tree_name", updateTree.TreeName, tree)
	if res.Error != nil {
		log.Println("updateUser err:", res.Error.Error())
		return c.SendStatus(500)
	} else if res.RowsAffected != 1 {
		log.Println("RowsAffected:", res.RowsAffected)
		return c.SendStatus(500)
	}

	return c.Status(200).SendString("tree: " + updateTree.TreeName + " is updated")
}

func DeleteTree(c *fiber.Ctx) error {
	deleteTree := new(models.ManageTree)

	if err := c.BodyParser(deleteTree); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	log.Println("uuid: ", deleteTree.Uuid)

	if deleteTree.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	var tree models.Tree
	res := controllers.Delete("trees", "tree_name", deleteTree.TreeName, &tree)
	if res.Error != nil {
		log.Println("deleteUser err:", res.Error.Error())
		return c.Status(500).SendString(res.Error.Error())
	} else if res.RowsAffected != 1 {
		log.Println("RowsAffected:", res.RowsAffected)
		return c.Status(500).SendString(res.Error.Error())
	}

	return c.Status(200).SendString("tree: " + deleteTree.TreeName + " is deleted")
}

func Transfer(c *fiber.Ctx) error {
	newOwner := new(models.ManageTree)

	if err := c.BodyParser(newOwner); err != nil {
		log.Println("BodyParser err:", err.Error())
		return c.Status(400).SendString(err.Error())
	}

	log.Println("uuid: ", newOwner.Uuid)

	if newOwner.Uuid != "uuid-9how,hlug-up;" {
		return c.SendStatus(403)
	}

	// query
	resQuery, err := controllers.FindUser(newOwner.Username)
	if err != nil { // err = record not found
		return c.Status(404).SendString(err.Error())
	}

	tree := models.Tree{Owner: resQuery.Uuid}
	resUpdate := controllers.Update("trees", "tree_name", newOwner.TreeName, tree)
	if resUpdate.Error != nil {
		log.Println("updateUser err:", resUpdate.Error.Error())
		return c.SendStatus(500)
	} else if resUpdate.RowsAffected != 1 {
		log.Println("RowsAffected:", resUpdate.RowsAffected)
		return c.SendStatus(500)
	}

	return c.Status(200).SendString("tree: " + newOwner.TreeName + " is transfered")
}
