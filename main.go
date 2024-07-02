// Test koneksi DB
// package main

// import (
//     "log"
//     "dummy_service/database"
//     "github.com/joho/godotenv"
// )

// func main() {
//     err := godotenv.Load()
//     if err != nil {
//         log.Fatalf("Error loading .env file")
//     }

//     database.Connect()
// }

package main

import (
	"log"
	// "dummy_service/config"
	"dummy_service/database"
	"dummy_service/routes"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Connect()
	router := routes.SetupRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
