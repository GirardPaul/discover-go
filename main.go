package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"test-go-project/database"
	_ "test-go-project/docs/testproject"
	"test-go-project/testcontroller"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /
// @schemes http
func main() {
	app := fiber.New()
	app.Get("/api/test", testcontroller.GetAll)
	app.Get("/api/test/:id", testcontroller.GetOne)
	app.Post("/api/test", testcontroller.Create)
	app.Put("/api/test/:id", testcontroller.UpdateOne)
	app.Delete("/api/test/:id", testcontroller.RemoveOne)
	app.Get("/swagger/*", swagger.HandlerDefault)
	database.ConnectDB()
	app.Listen(":3001")
}
