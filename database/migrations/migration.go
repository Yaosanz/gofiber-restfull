package migrations

import (
	"fmt"
	"golang-resfull/database"
	"golang-resfull/models"
	"log"
)

func Migration() {
 err := database.DB.AutoMigrate(
  &models.User{},

 )
 if err != nil {
  log.Fatal("Failed to migrate...")
 }

 fmt.Println("Migrated successfully")
}