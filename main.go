package main

import (
	"api/db"
	"api/notes"
	"api/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  app := fiber.New()

  app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
  }))

  db.InitDB()

  app.Get("/test", func (c *fiber.Ctx) error {
    return c.SendString("Hello world")
  })

  user.Routes(app)
  notes.Routes(app)

  app.Listen(":4200")
}