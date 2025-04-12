package db

import (
	"fmt"
	"log"
	"os"
	"to-do-planner/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL") // Get from environment variable
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment variables.")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to the database!")

	// Auto Migrate the schemas
	err = DB.AutoMigrate(&model.Provider{}, &model.Task{}, &model.Developer{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
