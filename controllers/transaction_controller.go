package controllers

import (
    "encoding/json"
    "dummy_service/database"
    "dummy_service/models"
    "github.com/gorilla/mux"
    "github.com/go-redis/redis/v8"
    "context"
    "net/http"
    "log"
    "fmt"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
    var transaction models.Transaction
    err := json.NewDecoder(r.Body).Decode(&transaction)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    transaction.Status = "Sukses"

    sqlStatement := `INSERT INTO transactions (account_id, amount, status) VALUES ($1, $2, $3) RETURNING id`
    id := 0
    err = database.DB.QueryRow(sqlStatement, transaction.AccountID, transaction.Amount, transaction.Status).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }
    transaction.ID = id

    // Save to Redis
    ctx := context.Background()
    key := fmt.Sprintf("transaction:%d", transaction.ID)
    transactionJson, err := json.Marshal(transaction)
    if err != nil {
        log.Fatal(err)
    }
    err = database.RedisClient.Set(ctx, key, transactionJson, 0).Err()
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(transaction)
}
func GetTransactionsByAccountID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    accountID := vars["account_id"]

    ctx := context.Background()
    redisKey := fmt.Sprintf("transactions:%s", accountID)
    
    // Try to get data from Redis
    transactionsJson, err := database.RedisClient.Get(ctx, redisKey).Result()
    if err == redis.Nil {
        // Data not found in Redis, get from PostgreSQL
        rows, err := database.DB.Query("SELECT id, account_id, amount, status FROM transactions WHERE account_id = $1", accountID)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var transactions []models.Transaction
        for rows.Next() {
            var transaction models.Transaction
            err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Amount, &transaction.Status)
            if err != nil {
                log.Fatal(err)
            }
            transactions = append(transactions, transaction)
        }

        if err = rows.Err(); err != nil {
            log.Fatal(err)
        }

        // Save to Redis
        transactionsJsonBytes, err := json.Marshal(transactions)
        if err != nil {
            log.Fatal(err)
        }
        err = database.RedisClient.Set(ctx, redisKey, string(transactionsJsonBytes), 0).Err()
        if err != nil {
            log.Fatal(err)
        }

        // Send response
        w.Header().Set("Content-Type", "application/json")
        w.Write(transactionsJsonBytes)
    } else if err != nil {
        // Some other error occurred
        log.Fatal(err)
    } else {
        // Data found in Redis
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(transactionsJson))
    }
}

