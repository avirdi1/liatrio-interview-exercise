package main

import (
	"time" 
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() 

	// if someone sends a get request to the path '/', it runs the handler function which gives us access to HTTP request and response
	app.Get("/", func(c *fiber.Ctx) error {
		// more convenient to use fiber.Map than a struct
		return c.JSON(fiber.Map{
			"message":"My name is Anmol Virdi",
			"timestamp":time.Now().UnixMilli(),
			"update":"Test #2"
			})
	})

	app.Listen(":80") 
}
