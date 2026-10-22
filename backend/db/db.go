package db

import (
	"backend/internal/entities"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// DB CONNECTION
	dbURL := os.Getenv("DBURL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Connected to database!")
	db.AutoMigrate(&entities.User{}, &entities.Post{})
	DB = db
}
