package notes

import (
	"api/db"
	"api/user"
	"api/utils"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
  g := app.Group("/v1/notes")

  g.Post("/new", user.AuthMiddleware, func (c *fiber.Ctx) error {
    var note db.Note
    json.Unmarshal(c.Body(), &note)

    key := fmt.Sprintf("%v", c.Locals("key"))

    note.Put(key)

    return c.JSON(note)
  })

  g.Get("", user.AuthMiddleware, func (c *fiber.Ctx) error {
    key := fmt.Sprintf("%v", c.Locals("key"))

    dateMonth, _ := strconv.Atoi(c.Query("dateMonth"))
    dateDay, _ := strconv.Atoi(c.Query("dateDay"))

    notes, err := db.GetNotes(
      base.Query {
        {
          "dateMonth": dateMonth,
          "dateDay": dateDay,
          "userKey": key,
        },
      },
    )
    if err != nil {
      return utils.Error(c, err)
    }

    return c.JSON(notes)
  })

  g.Put("/toggle", user.AuthMiddleware, func (c *fiber.Ctx) error {
    note, err := db.ToggleComplete(c.Query("key"))

    if err != nil {
      return utils.Error(c, err)
    }

    return c.JSON(note)
  })
}