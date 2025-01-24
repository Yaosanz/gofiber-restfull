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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Fiber app
	app := fiber.New()

	// Initialize the database
	database.DatabaseInit()

	// Run the migrations (fresh migration)
	migrations.MigrateFresh()

	// Setup Swagger documentation info
	docs.SwaggerInfo.Title = "Golang RESTful API"
	docs.SwaggerInfo.Description = "API Documentation for Golang RESTful API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Setup the root route to provide a simple message
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Restful API with Fiber Golang, go to /swagger/index.html for documentation",
		})
	})

	// Setup version route
	app.Get("/version*", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "API Version 1.0.0",
		})
	})

	// Setup Swagger UI route to show the API documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Initialize all routes (from routes package)
	routes.RouteInit(app)

	// Get server port from environment variable or use default
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080" // Default port
	}

	// Start the Fiber server on the defined port
	app.Listen(serverPort)
}
