package main

import (
	"github.com/gofiber/fiber/v2"
)

type JSONTextResponse struct{
	Message string
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString("Hello fiber!")
		return c.JSON(JSONTextResponse{Message:"Hello fiber!"})
	})
	app.Listen(":8080")
}
