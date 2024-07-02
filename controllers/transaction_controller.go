package controllers

import (
    "encoding/json"
    "dummy_service/database"
    "dummy_service/models"
    "github.com/gorilla/mux"
    "net/http"
    "log"
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

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(transaction)
}

func GetTransactionsByAccountID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    accountID := vars["account_id"]

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

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(transactions)
}
