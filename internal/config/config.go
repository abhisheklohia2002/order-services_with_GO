package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type LoadEnv struct {
	HOST     string
	USER     string
	PASSWORD string
	DATABASE string
	PORT     string
	DBPORT   string
}

func ConfigLoadEnv() LoadEnv {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// JWT_ISSUER := os.Getenv("JWT_ISSUER")

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")
	DBPort := os.Getenv("DB_PORT")
	return LoadEnv{
		PORT:     port,
		USER:     user,
		PASSWORD: password,
		DATABASE: database,
		HOST:     host,
		DBPORT:   DBPort,
	}

}
