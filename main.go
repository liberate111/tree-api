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

	// user
	app.Get("/", handlers.HealthCheck)
	app.Post("/login", handlers.Login)
	app.Post("/change_password", handlers.ChangePassword)

	// tree
	app.Post("/tree", handlers.GetTreeList)

	// admin manage users
	app.Post("/getuser", handlers.FindUser)
	app.Post("/adduser", handlers.AddUser)
	app.Post("/updateuser", handlers.UpdateUser)
	app.Post("/deleteuser", handlers.DeleteUser)

	// admin manage tree
	app.Post("/gettree", handlers.FindTree)
	app.Post("/addtree", handlers.AddTree)
	app.Post("/updatetree", handlers.UpdateTree)
	app.Post("/deletetree", handlers.DeleteTree)

	// admin transfer tree
	app.Post("/transfer", handlers.Transfer)

	app.Listen(":3000")
}
