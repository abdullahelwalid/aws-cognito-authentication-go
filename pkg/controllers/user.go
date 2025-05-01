package controllers

import (
	"fmt"

	"github.com/abdullahelwalid/aws-cognito-authentication-go/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	type FormData struct {
			Email    string `form:"email"`
			Password string `form:"password"`
		}
	var data FormData
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse form data",
			})
		}
	fmt.Println(data)
	auth, err := utils.InitAWSConfig()
	if err != nil {
		fmt.Println(err)
		c.SendStatus(500)
		return c.SendString("Something went wrong")
	}
	err = auth.Signup(data.Email, data.Password) 
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			})	}
	return c.SendString("sign up successful")
}

func ConfirmSignUp(c *fiber.Ctx) error {
	type FormData struct {
		Email string `form:"email"`
		Code string `form:"code"`
	}
	var data FormData
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse form data",
			})
		}
	auth, err := utils.InitAWSConfig()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})	}
	err = auth.ConfirmSignUp(data.Email, data.Code)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			})	}
	return c.SendString("sign up successful")
}

func Login(c *fiber.Ctx) error {
	type FormData struct {
		Email string `form:"email"`
		Password string `form:"password"`
	}
	var data FormData
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse form data",
			})
		}
	auth, err := utils.InitAWSConfig()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})	}
	resp, err := auth.Login(data.Email, data.Password)
	if err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{"error": "Invalid credentials"})
	}
	return c.JSON(resp)
}
