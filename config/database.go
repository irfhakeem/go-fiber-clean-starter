package config

import (
	"fmt"
	"os"

	"github.com/irfhakeem/go-fiber-clean-starter/helpers/constants"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	if os.Getenv("APP_ENV") != constants.ENV_RUN_PRODUCTION {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file: " + err.Error())
		}
	}

	requiredEnvVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_NAME", "DB_PORT"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			panic(fmt.Sprintf("Environment variable %s is not set", envVar))
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	return db
}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
