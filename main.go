package main

import (
	"strconv"
	//"encoding/json"
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
		/*return c.JSON(fiber.Map{
			"message":"My name is Anmol Virdi",
			"timestamp": time.Now().UnixMilli()})
		
		data := fiber.Map{
			"message":"My name is Anmol Virdi",
			"timestamp": time.Now().UnixMilli(),
		}
		b, _ := json.Marshal(data)
		c.Set("Content-Type", "application/json")
		return c.Send(b)
		*/
		ts := strconv.FormatInt(time.Now().UnixMilli(), 10)
		// exact string, no spaces or newline outside of JSON syntax
		body := `{"message":"My name is Anmol Virdi","timestamp":` + ts + `}`
		c.Set("Content-Type", "application/json")
		return c.SendString(body)
	})

	app.Listen(":80")
}
