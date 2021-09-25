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

	// health check
	app.Get("/", handlers.HealthCheck)

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1

	// user
	users := v1.Group("/users")                             // /api/v1/users
	users.Post("/login", handlers.Login)                    // /api/v1/users/login
	users.Post("/change-password", handlers.ChangePassword) // /api/v1/users/change-password
	users.Get("/:id/trees", handlers.GetTreeList)           // /api/v1/users/:id/trees

	// admin
	admin := v1.Group("/admin") // /api/v1/admin

	// admin manage user
	admin.Post("/users", handlers.FindUser)          // /api/v1/admin/users
	admin.Post("/users/insert", handlers.AddUser)    // /api/v1/admin/users/insert
	admin.Post("/users/update", handlers.UpdateUser) // /api/v1/admin/users/update
	admin.Post("/users/delete", handlers.DeleteUser) // /api/v1/admin/users/delete

	// admin manage tree
	admin.Post("/trees", handlers.FindTree)          // /api/v1/admin/trees
	admin.Post("/trees/insert", handlers.AddTree)    // /api/v1/admin/trees/insert
	admin.Post("/trees/update", handlers.UpdateTree) // /api/v1/admin/trees/update
	admin.Post("/trees/delete", handlers.DeleteTree) // /api/v1/admin/trees/delete

	// admin transfer tree
	admin.Post("/trees/transfer", handlers.Transfer) // /api/v1/admin/trees/transfer

	app.Listen(":3000")
}
