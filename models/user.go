package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             string `gorm:"type:char(36);primaryKey" json:"id"`
	Name           string `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Email          string `json:"email" gorm:"type:varchar(255);uniqueIndex" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	Phone          string `json:"phone" validate:"required"`
	ConfirmPassword string `json:"confirm_password" gorm:"-" validate:"required"`
}

var validate = validator.New()

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.ID == "" {
        u.ID = uuid.New().String()
    }
    return
}