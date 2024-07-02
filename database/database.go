package database

import (
    "database/sql"
    "fmt"
    "log"
    "dummy_service/config"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    host := config.GetEnv("DB_HOST", "localhost")
    port := config.GetEnv("DB_PORT", "5432")
    user := config.GetEnv("DB_USER", "user")
    password := config.GetEnv("DB_PASSWORD", "password")
    dbname := config.GetEnv("DB_NAME", "dbname")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected to the database")
}
