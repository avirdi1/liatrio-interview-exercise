package main

import (
	
	"encoding/json"
	"time" 
	"github.com/gofiber/fiber/v2"
)
/*
type MT struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

type Resp struct {
    Message   string `json:"message"`
    Timestamp int64  `json:"timestamp"`
}
*/

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":"My name is Anmol Virdi",
			"timestamp":time.Now().UnixMilli()})
			/*
		data := fiber.Map{
			"message":"My name is Anmol Virdi",
			"timestamp": time.Now().UnixMilli(),
		}
		b, _ := json.Marshal(data)
		c.Set("Content-Type", "application/json")s
		return c.Send(b)
		
		    r := Resp{
        Message:   "My name is Anmol Virdi",
        Timestamp: time.Now().UnixMilli(),
		}
		b, _ := json.Marshal(r) // compact JSON, no spaces or newline
		c.Response().Header.SetContentType(fiber.MIMEApplicationJSON)
		c.Response().SetBodyRaw(b) // send bytes exactly
		return nil
		*/
	})

	app.Listen(":80")
}
