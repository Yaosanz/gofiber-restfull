package main

import (
	"log"
	"golang-resfull/database"
	"golang-resfull/database/migrations"
	"golang-resfull/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Environment")
	}

	app := fiber.New()

	database.DatabaseInit()

	migrations.Migration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	routes.RouteInit(app)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080" 
	}
	app.Listen(serverPort)
}
