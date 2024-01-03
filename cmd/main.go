package cmd

import (
	"fmt"
	"log"
	"makerspace-api/api"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Run() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portStr := os.Getenv("PORT")

	serverPort, err := strconv.Atoi(portStr) 
	if err != nil {
		fmt.Println("Error converting port to string", err)
		return
	}

 	// Set up the API routes and handlers
	router := api.SetupRouter()
	fmt.Printf("Server listening on port %d...\n", serverPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router)
	if err != nil {
		log.Fatal("Server error:", err)
	}

}

// import (
// 	"makerspace-api/api"
// 	"makerspace-api/pkg/database"
// 	"fmt"
// 	"log"
// 	"net/http"
// )
//
//
// func Run() {
// 	// Initialize the database connection
// 	db, err := database.Connect()
// 	if err != nil {
// 		log.Fatal("Failed to connect to the database:", err)
// 	}
//
// 	// Set up the API routes and handlers
// 	router := api.SetupRouter()
//
// 	// Start the server
// 	serverPort := 8080
// 	fmt.Printf("Server listening on port %d...\n", serverPort)
// 	err = http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router)
// 	if err != nil {
// 		log.Fatal("Server error:", err)
// 	}
// }
