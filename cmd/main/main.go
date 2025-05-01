package main

import (
	"github.com/abdullahelwalid/aws-cognito-authentication-go/pkg/routes"
	"github.com/abdullahelwalid/aws-cognito-authentication-go/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)


func main() {
	app := fiber.New()
	routes.MainRouter(app)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//init DB
	utils.InitDB()
	log.Info("Server running on port 8000")
	app.Listen(":8000")
}
