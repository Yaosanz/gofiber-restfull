package routes

import (
	"golang-resfull/controllers"
	"golang-resfull/middleware"
	"github.com/gofiber/fiber/v2"
)

// RouteInit initializes all the routes
func RouteInit(app *fiber.App) {
	// Register Endpoint
	// @Summary Register a new user
	// @Description Register a new user with their credentials
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param user body models.User true "User data"
	// @Success 200 {object} models.User
	// @Failure 400 {object} fiber.Map
	// @Router /register [post]
	app.Post("/register", controllers.Register)

	// Login Endpoint
	// @Summary User login
	// @Description Login for an existing user to get access token
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param credentials body models.LoginCredentials true "Login credentials"
	// @Success 200 {string} string "JWT Token"
	// @Failure 401 {object} fiber.Map
	// @Router /login [post]
	app.Post("/login", controllers.Login)

	// Get All Users Endpoint
	// @Summary Get all users
	// @Description Get a list of all registered users
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.User
	// @Failure 400 {object} fiber.Map
	// @Router /users [get]
	app.Get("/users", middleware.AuthMiddleware(), controllers.GetAllUsers)

	// Get User By ID Endpoint
	// @Summary Get a user by ID
	// @Description Get a single user by their unique ID
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param id path string true "User ID"
	// @Success 200 {object} models.User
	// @Failure 404 {object} fiber.Map
	// @Router /users/{id} [get]
	app.Get("/users/:id", middleware.AuthMiddleware(), controllers.GetUserById)

	// Delete User Endpoint
	// @Summary Delete a user by ID
	// @Description Remove a user by their unique ID
	// @Tags Users
	// @Accept json
	// @Produce json
	// @Param id path string true "User ID"
	// @Success 200 {object} fiber.Map
	// @Failure 404 {object} fiber.Map
	// @Router /users/{id} [delete]
	app.Delete("/users/:id", middleware.AuthMiddleware(), controllers.DeleteUserById)
}
