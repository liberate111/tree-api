package handlers

import (
	"tree-web-server/controllers"
	"tree-web-server/models"

	"github.com/gofiber/fiber/v2"
)

func GetTreeList(c *fiber.Ctx) error {
	// get tree list form db
	uuid := c.Params("id")

	// query
	res, err := controllers.FindTree(uuid)
	if err != nil {
		resJSON := models.ResponseMessage{
			Status:  404,
			Message: err.Error(),
		}
		return c.Status(404).JSON(resJSON)
	}
	var treeList []string
	for _, v := range res {
		treeList = append(treeList, v.TreeName)
	}

	resJSON := models.ResponseMessage{
		Status:  200,
		Message: "success",
		Data: &models.Data{
			Tree: treeList,
		},
	}
	return c.Status(200).JSON(resJSON)
}
