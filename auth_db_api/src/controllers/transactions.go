package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
)

func FetchTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transactions
	db.Table("transactions").Limit(10).Find(&transactions)
	json.NewEncoder(w).Encode(transactions)

	//The below is a superfluous block to verify how contexts work
	transCtx := r.Context().Value("transactions")
	structFromCtx, _ := json.Marshal(transCtx)
	fmt.Println("\nFrom FetchUsers(): ")
	fmt.Println(string(structFromCtx))
}
