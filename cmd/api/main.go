package main

import (
	"log"

	"github.com/LuisCabantac/go-portfolyo-api/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	server.New().Run()
}
