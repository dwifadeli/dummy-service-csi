package database

import (
    "context"
    "log"
    "dummy_service/config"
    "github.com/go-redis/redis/v8"
    "strconv"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis() {
    addr := config.GetEnv("REDIS_ADDR", "localhost:6379")
    password := config.GetEnv("REDIS_PASSWORD", "")
    dbStr := config.GetEnv("REDIS_DB", "0")

    db, err := strconv.Atoi(dbStr)
    if err != nil {
        log.Fatalf("Invalid Redis DB number: %v", err)
    }

    RedisClient = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    _, err = RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }

    log.Println("Successfully connected to Redis")
}
