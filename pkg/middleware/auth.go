package middleware

import (
	"fmt"
	"strings"

	"github.com/abdullahelwalid/aws-cognito-authentication-go/pkg/utils"
	"github.com/gofiber/fiber/v2"
)


func Authentication(c *fiber.Ctx) error{
	authHeader := c.Get("Authorization")	
	fmt.Println("AUTH ROUTE")
	if (authHeader == "" || len(strings.Split(authHeader, " ")) < 2){
		fmt.Println("NO AUTH HEADER")
		c.SendStatus(401)
		return c.SendString("Auth Header can't be empty")
	}
	auth, err := utils.InitAWSConfig()
	if err != nil {
		fmt.Println(err)
		c.SendStatus(500)
		return c.SendString("Something went wrong")
	}
	_, err = auth.ValidateToken(strings.Split(authHeader, " ")[1])
	if err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{"error": "Something went wrong when validating token"})
	}
	return c.Next()
}
