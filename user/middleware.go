package user

import (
	"api/db"
	"api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
  var token string

  authHeader := c.Get("Authorization")

  if string(authHeader) != "" && strings.HasPrefix(string(authHeader), "Bearer") {
    token = strings.Fields(string(authHeader))[1]

    user, err := db.ParseUserToken(token) 

    if err != nil {
      return utils.Error(c, err)
    }
    
    c.Locals("key", user.Key)
    utils.SetLocals(c, "user", user)
  }

  if (token == "") {
    return utils.MessageError(c, "no token")
  }

  return c.Next()
}