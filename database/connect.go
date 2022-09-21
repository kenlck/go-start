package database

import (
	"crm_backend/config"
	"crm_backend/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	DB, err = gorm.Open(postgres.Open(config.DATABASE_URL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
