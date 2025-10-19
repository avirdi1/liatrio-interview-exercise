package main

import (
	// To be able to use unix time and be able to use the Fiber framework
	"time" 
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":"My name is Anmol Virdi",
			"timestamp":time.Now().UnixMilli(),
			"quote":"To be or not to be"})
	})

	app.Listen(":80")
}
