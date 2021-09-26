package handlers

import (
	"fmt"
	"strconv"
	"time"
	"tree-web-server/controllers"
	"tree-web-server/models"

	"github.com/gofiber/fiber/v2"
)

func GetTreeList(c *fiber.Ctx) error {
	// get tree list form db
	uuid := c.Params("id")

	// query
	res, err := controllers.FindTree("owner", uuid)
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

func GetTreeState(c *fiber.Ctx) error {
	// /api/v1/users/:id/trees/:treeId/state
	uuid := c.Params("id")
	treeId := c.Params("treeId")

	// query
	res, err := controllers.FindTree("tree_name", treeId)
	if err != nil {
		resJSON := models.ResponseMessage{
			Status:  404,
			Message: err.Error(),
		}
		return c.Status(404).JSON(resJSON)
	}
	if len(res) != 1 {
		resJSON := models.ResponseMessage{Status: 500, Message: "RowsAffected not equal to 1"}
		return c.Status(500).JSON(resJSON)
	}

	tree := res[0]
	if tree.Owner != uuid {
		resJSON := models.ResponseMessage{Status: 400, Message: "Wrong uuid"}
		return c.Status(400).JSON(resJSON)
	}
	// phase spawn tree
	if tree.StartTime == 0 && tree.StopTime == 0 {
		resJSON := models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Level: tree.Level, State: tree.State, StartTime: tree.StartTime, StopTime: tree.StopTime}}
		return c.Status(200).JSON(resJSON)
	}
	if int(time.Now().UTC().Unix()) > tree.StopTime {
		// update new state: grow
		tree.State = "grow"
		res := controllers.Update("trees", "tree_name", tree.TreeName, tree)
		if res.Error != nil {
			resJSON := models.ResponseMessage{Status: 500, Message: res.Error.Error()}
			return c.Status(500).JSON(resJSON)
		} else if res.RowsAffected != 1 {
			resJSON := models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", res.RowsAffected)}
			return c.Status(400).JSON(resJSON)
		}
		resJSON := models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Level: tree.Level, State: tree.State, StartTime: tree.StartTime, StopTime: tree.StopTime}}
		return c.Status(200).JSON(resJSON)
	}
	resJSON := models.ResponseMessage{Status: 200, Message: "reject"}
	return c.Status(200).JSON(resJSON)
}

func UpdateTreeState(c *fiber.Ctx) error {
	// /api/v1/users/:id/trees/:treeId/state
	uuid := c.Params("id")
	treeId := c.Params("treeId")

	// query
	res, err := controllers.FindTree("tree_name", treeId)
	if err != nil {
		resJSON := models.ResponseMessage{
			Status:  404,
			Message: err.Error(),
		}
		return c.Status(404).JSON(resJSON)
	}
	if len(res) != 1 {
		resJSON := models.ResponseMessage{Status: 500, Message: "RowsAffected not equal to 1"}
		return c.Status(500).JSON(resJSON)
	}

	tree := res[0]
	if tree.Owner != uuid {
		resJSON := models.ResponseMessage{Status: 400, Message: "Wrong uuid"}
		return c.Status(400).JSON(resJSON)
	}

	// phase watering
	tree.State = "wet"
	tree.StartTime = int(time.Now().UTC().Unix())
	// tree.StopTime = tree.StartTime + (24 * 60 * 60) // add 1 day
	tree.StopTime = tree.StartTime + (60) // add 1 min

	resUpdate := controllers.Update("trees", "tree_name", tree.TreeName, tree)
	if resUpdate.Error != nil {
		resJSON := models.ResponseMessage{Status: 500, Message: resUpdate.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if resUpdate.RowsAffected != 1 {
		resJSON := models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", resUpdate.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}
	resJSON := models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Level: tree.Level, State: tree.State, StartTime: tree.StartTime, StopTime: tree.StopTime}}
	return c.Status(200).JSON(resJSON)
}

func UpdateTreeLevel(c *fiber.Ctx) error {
	// /api/v1/users/:id/trees/:treeId/levelup?level=1
	uuid := c.Params("id")
	treeId := c.Params("treeId")
	level := c.Query("level")
	if level == "" {
		resJSON := models.ResponseMessage{Status: 400, Message: "missing query params"}
		return c.Status(400).JSON(resJSON)
	}

	// query
	res, err := controllers.FindTree("tree_name", treeId)
	if err != nil {
		resJSON := models.ResponseMessage{
			Status:  404,
			Message: err.Error(),
		}
		return c.Status(404).JSON(resJSON)
	}
	if len(res) != 1 {
		resJSON := models.ResponseMessage{Status: 500, Message: "RowsAffected not equal to 1"}
		return c.Status(500).JSON(resJSON)
	}

	tree := res[0]
	if tree.Owner != uuid {
		resJSON := models.ResponseMessage{Status: 400, Message: "Wrong uuid"}
		return c.Status(400).JSON(resJSON)
	}

	levelInt, err := strconv.Atoi(level)
	if err != nil {
		resJSON := models.ResponseMessage{Status: 400, Message: "Wrong level"}
		return c.Status(400).JSON(resJSON)
	}
	if tree.Level != levelInt {
		resJSON := models.ResponseMessage{Status: 200, Message: "reject"}
		return c.Status(200).JSON(resJSON)
	}

	// phase levelup
	tree.Level = tree.Level + 1
	tree.State = "dry"

	resUpdate := controllers.Update("trees", "tree_name", tree.TreeName, tree)
	if resUpdate.Error != nil {
		resJSON := models.ResponseMessage{Status: 500, Message: resUpdate.Error.Error()}
		return c.Status(500).JSON(resJSON)
	} else if resUpdate.RowsAffected != 1 {
		resJSON := models.ResponseMessage{Status: 400, Message: fmt.Sprintf("RowsAffected: %d", resUpdate.RowsAffected)}
		return c.Status(400).JSON(resJSON)
	}
	resJSON := models.ResponseMessage{Status: 200, Message: "success", Data: &models.Data{Level: tree.Level, State: tree.State, StartTime: tree.StartTime, StopTime: tree.StopTime}}
	return c.Status(200).JSON(resJSON)
}
