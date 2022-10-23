package user

import (
	"api/db"
	"encoding/json"
	"fmt"

	"api/utils"

	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Routes(app *fiber.App) {
  g := app.Group("/v1/user")

  g.Post("/login", func (c *fiber.Ctx) error {
    var body map[string]string
    json.Unmarshal(c.Body(), &body)

    user, err := db.GetUser(
      base.Query {
        {"email": body["email"]},
      },
    )
    if err != nil {
      return utils.MessageError(c, "there is no user with that email :(") 
    }

    compareErr := bcrypt.CompareHashAndPassword(
      []byte(user.Password),
      []byte(body["password"]),
    )
    if compareErr != nil {
      return utils.MessageError(c, "that password is not correct.")
    }

    token := db.GenUserToken(user)

    return c.JSON(map[string]interface {} {
      "user": user,
      "token": token,
    })
  })
  
  g.Post("/signup", func (c *fiber.Ctx) error {
    var user db.User
    json.Unmarshal(c.Body(), &user)

    user.Put()

    return c.JSON(user)
  })

  g.Get("/profile", AuthMiddleware, func (c *fiber.Ctx) error {
    c.Context().Response.Header.Add("Content-Type", "application/json")
    return c.Send([]byte(fmt.Sprintf("%v", c.Locals("user"))))
  })
}