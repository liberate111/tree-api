package handlers

import (
	"log"
	"tree-web-server/controllers"
	"tree-web-server/models"

	"github.com/gofiber/fiber/v2"
)

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
