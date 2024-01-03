package main

import (
	"makerspace-api/cmd"
	"log"
	"github.com/joho/godotenv"

)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Run()
}




