package controllers

import (
	"golang-resfull/database"
	"golang-resfull/models"
	"golang-resfull/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	if user.Password != user.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password and confirm password do not match",
		})
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
			"error":   err.Error(),
		})
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating token",
			"error":   err.Error(),
		})
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashPassword,
		Token:    token,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":    newUser.ID,
			"name":  newUser.Name,
			"email": newUser.Email,
			"phone": newUser.Phone,
		},
		"token": token,
	})
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
			"error":   err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	token, err := utils.GenerateJWT(existingUser.ID, existingUser.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating token",
			"error":   err.Error(),
		})
	}

	existingUser.Token = token
	if err := database.DB.Save(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user token",
			"error":   err.Error(),
		})
	}

	existingUser.Password = ""
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"user": fiber.Map{
			"id":    existingUser.ID,
			"name":  existingUser.Name,
			"email": existingUser.Email,
			"phone": existingUser.Phone,
		},
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
			"error":   err.Error(),
		})
	}

	email := claims["email"].(string)

	var users []models.User
	if err := database.DB.Debug().Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving users",
			"error":   err.Error(),
		})
	}

	var userList []fiber.Map
	for _, user := range users {
		if user.Email == email { 
			userList = append(userList, fiber.Map{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
				"phone": user.Phone,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get all users",
		"users":   userList,
	})
}

func GetUserById(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
			"error":   err.Error(),
		})
	}

	userID := claims["sub"].(string)
	id := c.Params("id")
	if userID != id {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You do not have access to this resource",
		})
	}

	var user models.User
	if err := database.DB.Debug().First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"phone": user.Phone,
		},
	})
}

func DeleteUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
	}

	var user models.User
	if err := database.DB.Debug().First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := database.DB.Debug().Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}