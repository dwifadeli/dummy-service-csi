package main
import (
    "log"
    "dummy_service/database"
    "dummy_service/routes"
    "github.com/joho/godotenv"
    "net/http"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    database.ConnectPostgres()
    database.ConnectRedis()

    router := routes.SetupRoutes()

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
