package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Account struct {
	ID             int     `json:"id"`
	Limit          float64 `json:"avaliable_credit_limit"`
	DocumentNumber string  `json:"document_number"`
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
	dbURL := os.Getenv("DATABASE_URL")

	var err error
	db, err = sql.Open("mysql", dbURL)

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
	fmt.Println(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO accounts (document_number, avaliable_credit_limit) VALUES (?, ?)", account.DocumentNumber, account.Limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	err = db.QueryRow("SELECT * FROM accounts WHERE account_id = ?", accountID).Scan(&account.ID, &account.DocumentNumber, &account.Limit)
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
	var account Account

	err := json.NewDecoder(r.Body).Decode(&transaction)

	if transaction.OperationTypeID != 4 {
		transaction.Amount = -1 * transaction.Amount
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.QueryRow("SELECT * FROM accounts WHERE account_id = ?", transaction.AccountID).Scan(&account.ID, &account.DocumentNumber, &account.Limit)
	if account.Limit+transaction.Amount >= 0 {
		_, err = db.Exec("INSERT INTO transactions (account_id, operationType_ID, amount) VALUES (?, ?, ?)", transaction.AccountID, transaction.OperationTypeID, transaction.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = db.Exec("UPDATE accounts SET avaliable_credit_limit = ? WHERE account_id = ?", account.Limit+transaction.Amount, transaction.AccountID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Limite não disponivel", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Transaction created successfully")
}
