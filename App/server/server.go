package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	model "github.com/himanshu1221/GoShort/Model"
	"github.com/himanshu1221/GoShort/utils"
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

func CreateShort(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var short model.Goshort
	err := c.BodyParser(&short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON" + err.Error(),
		})
	}
	if short.Random {
		short.Goshort = utils.RandomUrl(8)
	}
	err = model.CreateShort(short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not short in DB" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(short)
}

func UpdateShorts(c *fiber.Ctx) error {
	c.Accepts("applications/json")
	var goshort model.Goshort
	err := c.BodyParser(&goshort)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Message could not parse" + err.Error(),
		})
	}
	err = model.UpdateShort(goshort)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update short link in DB" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goshort)
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	router.Get("/shorts", GetAllRedirects)
	router.Get("/shorts/:id", GetRedirect)
	router.Post("/shorts", CreateShort)
	router.Patch("/updateshorts", UpdateShorts)
	router.Listen(":3000")
}
