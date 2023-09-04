package server

import (
	"strconv"

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

func GetRedirect(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id" + err.Error(),
		})
	}
	short, err := model.SingleShorts(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not short from DB" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(short)
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	router.Get("/shorts", GetAllRedirects)
	router.Get("/shorts/:id", GetRedirect)
	router.Listen(":3000")
}
