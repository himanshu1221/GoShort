package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	model "github.com/himanshu1221/GoShort/Model"
)

func GetAllRedirects(ctx *fiber.Ctx) error {
	shorts, err := model.GetAllShorts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all short links" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(shorts)
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	router.Get("/shorts", GetAllRedirects)
	router.Listen(":3000")
}
