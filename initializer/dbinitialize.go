package initializer

import (
	"log"
	"os"
	"url-shortner-backend/schema"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitDB() {
	schemaM := schema.Url_Data{}
	dsn := os.Getenv("DATABASE_URL")
	dbconnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to Database")
	}
	schema.DBConnect.DB = dbconnection
	dbconnection.AutoMigrate(&schemaM)
}
