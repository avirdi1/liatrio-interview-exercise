package main

import (
	"time" 
	"github.com/gofiber/fiber/v2"
)

type MT struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(MT{"My name is Anmol Virdi", time.Now().UnixMilli()})
	})

	app.Listen(":80")
}
