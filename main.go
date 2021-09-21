package main

import (
	"tree-web-server/controllers"
	"tree-web-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// sqlite
	controllers.ConnectDB()

	// fiber
	app := fiber.New()

	app.Get("/", handlers.HealthCheck)
	app.Post("/login", handlers.Login)
	app.Post("/tree", handlers.GetTreeList)
	app.Post("/adduser", handlers.AddUser)
	app.Post("/updateuser", handlers.UpdateUser)
	app.Post("/deleteuser", handlers.DeleteUser)
	app.Listen(":3000")
}
