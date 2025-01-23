package models

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name" gorm:"type:varchar(255)"`
	Email          string `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	ConfirmPassword string `json:"confirm_password" gorm:"-"`
}

var validate = validator.New()

func (u *User) Validate() error {
	return validate.Struct(u)
}
