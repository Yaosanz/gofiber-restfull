package migrations

import (
	"fmt"
	"golang-resfull/database"
	"golang-resfull/models"
	"log"
	"gorm.io/gorm"
)

func MigrateFresh() {
	db := database.DB

	err := dropAllTables(db)
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		// &models.AnotherModel{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migration fresh completed successfully")
}

func dropAllTables(db *gorm.DB) error {
	models := []interface{}{
		&models.User{},
		// &models.AnotherModel{},
	}

	for _, model := range models {
		err := db.Migrator().DropTable(model)
		if err != nil {
			return fmt.Errorf("failed to drop table for model %v: %w", model, err)
		}
	}

	return nil
}
