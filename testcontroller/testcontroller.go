package testcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"test-go-project/database"
	test "test-go-project/model"
)

func GetAll(c *fiber.Ctx) error {
	db := database.DB
	var test []test.Test

	db.Find(&test)

	if len(test) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No test present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Test Found", "data": test})
}
func GetOne(c *fiber.Ctx) error {
	db := database.DB
	var object test.Test

	err := db.Find(&object, test.Test{ID: uuid.MustParse(c.Params("id"))}).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Test not found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Test Found", "data": object})

}
func Create(c *fiber.Ctx) error {
	db := database.DB
	test := new(test.Test)

	err := c.BodyParser(test)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	test.ID = uuid.New()
	err = db.Create(&test).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": test})
}
func UpdateOne(c *fiber.Ctx) error {
	db := database.DB
	object := new(test.Test)

	err := db.Find(&object, test.Test{ID: uuid.MustParse(c.Params("id"))}).Error

	err = c.BodyParser(object)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Save(&object).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update test", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Updated Test", "data": object})
}
func RemoveOne(c *fiber.Ctx) error {
	db := database.DB
	var object test.Test

	err := db.Find(&object, test.Test{ID: uuid.MustParse(c.Params("id"))}).Error

	if object.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Test not found", "data": nil})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete test", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Test", "data": object})

}
