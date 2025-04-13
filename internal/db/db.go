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

	// Seed the database with initial data
	seedDevelopers(DB)
	seedProviders(DB)
	fmt.Println("Database migrated successfully!")
}

func GetDB() *gorm.DB {
	return DB
}

func seedDevelopers(db *gorm.DB) {
	developers := []model.Developer{
		{Name: "DEV1", Capacity: 1},
		{Name: "DEV2", Capacity: 2},
		{Name: "DEV3", Capacity: 3},
		{Name: "DEV4", Capacity: 4},
		{Name: "DEV5", Capacity: 5},
	}

	for _, dev := range developers {
		db.FirstOrCreate(&model.Developer{}, model.Developer{Name: dev.Name, Capacity: dev.Capacity})
	}
}

func seedProviders(db *gorm.DB) {
	providers := []model.Provider{
		{
			Name:         "Provider1",
			APIURL:       "https://raw.githubusercontent.com/WEG-Technology/mock/refs/heads/main/mock-one",
			ResponseKeys: "{\"hasKeyOfTaskList\": false,\"keyOfTaskList\": \"\",\"taskNameField\": \"id\",\"durationField\": \"estimated_duration\",\"difficultyField\": \"value\"}",
		},
		{
			Name:         "Provider2",
			APIURL:       "https://raw.githubusercontent.com/WEG-Technology/mock/refs/heads/main/mock-two",
			ResponseKeys: "{\"hasKeyOfTaskList\": false,\"keyOfTaskList\": \"\",\"taskNameField\": \"id\",\"durationField\": \"sure\",\"difficultyField\": \"zorluk\"}",
		},
	}

	for _, provider := range providers {
		db.FirstOrCreate(&model.Provider{}, model.Provider{Name: provider.Name, APIURL: provider.APIURL, ResponseKeys: provider.ResponseKeys})
	}
}
