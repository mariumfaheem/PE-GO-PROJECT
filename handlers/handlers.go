package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/mariumfaheem/PE-GO-PROJECT/database"
	"github.com/mariumfaheem/PE-GO-PROJECT/models"
)

func UserList(ctx *fiber.Ctx) error {
	user := database.Get()

	return ctx.JSON(fiber.Map{
		"sucess": true,
		"users":  user,
	})

}

func UserCreate(ctx *fiber.Ctx) error {
	user := &models.User{
		Name: utils.CopyString(ctx.FormValue("user")),
	}

	database.Insert(user)

	return ctx.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})

}

func NotFound(ctx *fiber.Ctx) error {
	return ctx.Status(404).SendFile("./static/private/404.html")
}
