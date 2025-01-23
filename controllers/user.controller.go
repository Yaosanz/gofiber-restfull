package controllers

import (
	"golang-resfull/database"
	"golang-resfull/models"
	"strconv"
	"golang-resfull/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
 var users []*models.User

 database.DB.Debug().Find(&users)

 return c.Status(fiber.StatusOK).JSON(fiber.Map{
  "message": "Success Get All Users",
  "users":   users,
 })
}

func CreateUsers(c *fiber.Ctx) error {
 user := new(models.User)

 if err := c.BodyParser(user); err != nil {
  return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
 }

 validate := validator.New()
 errValidate := validate.Struct(user)
 if errValidate != nil {
  return c.Status(400).JSON(fiber.Map{
   "message": "failed to validate",
   "error":   errValidate.Error(),
  })
 }

 newUser := models.User{
  Name:  user.Name,
  Email: user.Email,
  Phone: user.Phone,
 }

 hashPassword, err := utils.HashPassword(user.Password)
 if err != nil {
  return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
   "message": "status internal server error",
  })
 }

 newUser.Password = hashPassword

 database.DB.Debug().Create(&newUser)

 return c.Status(fiber.StatusOK).JSON(fiber.Map{
  "message": "Success Created new User",
 })
}

func GetUserById(c *fiber.Ctx) error {
 var user []*models.User

 result := database.DB.Debug().First(&user, c.Params("id"))

 if result.Error != nil {
  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
   "message": "User not found",
  })
 }

 return c.Status(fiber.StatusOK).JSON(fiber.Map{
  "user": user,
 })

}

func UpdateUser(c *fiber.Ctx) error {
 user := new(models.User)

 if err := c.BodyParser(user); err != nil {
  return c.Status(fiber.StatusBadRequest).JSON(err.Error())
 }

 id, _ := strconv.Atoi(c.Params("id"))

 database.DB.Debug().Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
  "name":  user.Name,
  "email": user.Email,
  "phone": user.Phone,
 })

 return c.Status(fiber.StatusOK).JSON(fiber.Map{
  "message": "succes update user",
 })
}

func DeleteUser(c *fiber.Ctx) error {
 user := new(models.User)

 id, _ := strconv.Atoi(c.Params("id"))

 database.DB.Debug().Where("id = ?", id).Delete(&user)

 return c.Status(fiber.StatusOK).JSON(fiber.Map{
  "message": "delete user successfully",
 })
}