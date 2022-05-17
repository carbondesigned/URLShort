package server

import (
	"fmt"
	"strconv"
	"url-shortener/model"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	shortUrl := c.Params("redirect")
	short, err := model.FindByShortUrl(shortUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find goly in DB " + err.Error(),
		})
	}
	// grab any stats you want...
	short.Clicked += 1
	err = model.UpdateShort(short)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(short.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllShorts(c *fiber.Ctx) error {
	shorts, err := model.GetAllShorts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error getting all shorts",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    shorts,
	})
}

func getShort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't get id" + err.Error(),
		})
	}
	short, err := model.GetShort(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't get id" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    short,
	})
}

func createShort(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var short model.URLShort
	if err := c.BodyParser(&short); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't parse body",
		})
	}
	if short.Random {
		short.URLShort = utils.RandomURL(8)
	}
	err := model.CreateShort(short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't create short",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    short,
	})
}

func updateShort(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var short model.URLShort
	if err := c.BodyParser(&short); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't parse body",
		})
	}
	err := model.UpdateShort(short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't update short",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    short,
	})
}

func DeleteShort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't get id" + err.Error(),
		})
	}
	err = model.DeleteShort(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error: couldn't delete short",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
	})
}

func SetupAndListen() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "welcome",
		})
	})
	app.Get("/r/:redirect", redirect)
	app.Get("/shorts", getAllShorts)
	app.Get("/short/:id", getShort)
	app.Delete("/short/:id", DeleteShort)
	app.Post("/short", createShort)
	app.Patch("/short", updateShort)

	app.Listen(":3000")
}
