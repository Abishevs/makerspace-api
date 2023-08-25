package main

import (
	"makerspace-api/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	cmd.Run()
}




