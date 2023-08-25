package cmd 

import (
	"makerspace-api/api"
	"fmt"
	"net/http"
)

func Run() {

 	// Set up the API routes and handlers
	router := api.SetupRouter()
	const serverPort int = 5000
	fmt.Printf("Server listening on port %d...\n", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router)

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
