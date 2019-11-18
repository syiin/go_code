package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	trans := &models.Transaction{}
	params := mux.Vars(r)
	var id = params["id"]
	db.Table("transactions").First(&trans, id)
	json.NewEncoder(w).Encode(&trans)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var trans models.Transaction
	db.Table("transactions").First(&trans, id)
	db.Table("transactions").Delete(&trans)
	json.NewEncoder(w).Encode("Transaction deleted")
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	trans := &models.Transaction{}
	params := mux.Vars(r)
	var id = params["id"]
	db.Table("transactions").First(&trans, id)
	json.NewDecoder(r.Body).Decode(trans)
	db.Table("transactions").Save(&trans)
	json.NewEncoder(w).Encode(&trans)
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
