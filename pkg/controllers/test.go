package controllers

import (
	"github.com/gofiber/fiber/v2"
)


func Test(c *fiber.Ctx) error {
	return c.SendString("Working!!!")
}

func AuthHandler(c *fiber.Ctx) error {
	return c.SendString("Auth Endpoint")
}
