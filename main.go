package main

import (
	"time" 
	"github.com/gofiber/fiber/v2"
)
/*
type MT struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}
*/

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":"My name is Anmol Virdi",
			"timestamp": time.Now().UnixMilli()})
	})

	app.Listen(":80")
}
