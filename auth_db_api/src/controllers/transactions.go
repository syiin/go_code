package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transaction := &models.Transaction{}
	json.NewDecoder(r.Body).Decode(transaction)
	fmt.Println(transaction)
	createdTrans := db.Table("transactions").Create(transaction)

	var err = createdTrans.Error
	if createdTrans.Error != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(createdTrans)

}

func FetchTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	db.Table("transactions").Limit(10).Find(&transactions)
	json.NewEncoder(w).Encode(transactions)

	//The below is a superfluous block to verify how contexts work
	transCtx := r.Context().Value("transactions")
	structFromCtx, _ := json.Marshal(transCtx)
	fmt.Println("\nFrom FetchUsers(): ")
	fmt.Println(string(structFromCtx))
}
