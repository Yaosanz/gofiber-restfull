package migrations

import (
	"fmt"
	"golang-resfull/database"

	"log"
)

func Migration() {
 err := database.DB.AutoMigrate()
 
 if err != nil {
  log.Fatal("Failed to migrate...")
 }

 fmt.Println("Migrated successfully")
}