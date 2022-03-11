package main

import (
	"github.com/gofiber/fiber/v2"
	"test-go-project/database"
	"test-go-project/testcontroller"
)

func apiRoutes(app *fiber.App) {
	app.Get("/api/test", testcontroller.GetAll)
	app.Get("/api/test/:id", testcontroller.GetOne)
	app.Post("/api/test", testcontroller.Create)
	app.Put("/api/test/:id", testcontroller.UpdateOne)
	app.Delete("/api/test/:id", testcontroller.RemoveOne)
}

func main() {
	app := fiber.New()

	apiRoutes(app)
	database.ConnectDB()
	app.Listen(":3001")
}
