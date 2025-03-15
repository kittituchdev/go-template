package example

import (
	"github.com/gofiber/fiber/v2"
)

func CreateExample(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    "create example",
	})
}

func GetAllExample(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    "get all example",
	})
}

func GetExample(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    "get example",
	})
}

func UpdateExample(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    "update example",
	})
}

func DeleteExample(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success": true,
		"data":    "delete example",
	})
}
