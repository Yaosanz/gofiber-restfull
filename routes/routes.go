package routes

import (
	"golang-resfull/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

 r.Get("/users", controllers.GetAllUsers)
 r.Post("/users", controllers.CreateUsers)
 r.Get("/users/:id", controllers.GetUserById)
 r.Patch("/users/:id", controllers.UpdateUser)
 r.Delete("/users/:id", controllers.DeleteUser)

}