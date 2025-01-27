package main

import (
	"golang-resfull/database"
	"golang-resfull/database/migrations"
	"golang-resfull/routes"
	"log"
	"os"

	"golang-resfull/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	database.DatabaseInit()

	migrations.MigrateFresh()

	docs.SwaggerInfo.Title = "Golang RESTful API"
	docs.SwaggerInfo.Description = "API Documentation for Golang RESTful API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Restful API with Fiber Golang, go to /swagger/index.html for documentation",
		})
	})

	app.Get("/version*", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "API Version 1.0.0",
		})
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.RouteInit(app)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080" 
	}

	app.Listen(serverPort)
}
