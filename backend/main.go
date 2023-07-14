package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Account struct {
	ID             int    `json:"id"`
	DocumentNumber string `json:"document_number"`
}

type Transaction struct {
	ID              int     `json:"id"`
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	EventDate       string  `json:"event_date"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "example_user:example_password@tcp(db:3306)/example_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/accounts", createAccount).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", getAccountByID).Methods("GET")
	router.HandleFunc("/transactions", createTransaction).Methods("POST")

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Inserir a nova conta no banco de dados
	result, err := db.Exec("INSERT INTO accounts (document_number) VALUES (?)", account.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obter o ID da conta inserida
	accountID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	account.ID = int(accountID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func getAccountByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.Atoi(vars["accountId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var account Account
	err = db.QueryRow("SELECT * FROM accounts WHERE account_id = ?", accountID).Scan(&account.ID, &account.DocumentNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO transactions (account_id, operationType_ID, amount) VALUES (?, ?, ?)", transaction.AccountID, transaction.OperationTypeID, transaction.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Transaction created successfully")
}
