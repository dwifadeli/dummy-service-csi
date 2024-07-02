package routes

import (
    "dummy_service/controllers"
    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
    router.HandleFunc("/transactions/{account_id}", controllers.GetTransactionsByAccountID).Methods("GET")
    return router
}
