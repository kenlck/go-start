package database

import (
	"crm_backend/config"
	"crm_backend/model"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.DB_PORT
	port, _ := strconv.ParseUint(p, 10, 32)
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", config.DB_HOST, port, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, port, config.DB_NAME)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
