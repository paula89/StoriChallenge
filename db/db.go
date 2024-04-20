package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

func SaveTransaction(db *sql.DB, UserId, TxId uuid.UUID, Date time.Time, Debit, Credit float64) {
	// Guardar información en la base de datos MySQL
	err := saveTxInDb(db, UserId, TxId, Date, Debit, Credit)
	if err != nil {
		log.Fatalf("Error saving transactions in the db: %v", err)
	}
}

func saveTxInDb(db *sql.DB, UserId, TxId uuid.UUID, Date time.Time, Debit, Credit float64) error {
	_, err := db.Exec("INSERT INTO Transactions (Id, UserId, CreationDate, Debit, Credit) VALUES (?, ?, ?)", TxId, UserId, Date, Debit, Credit)
	return err
}

func OpenConn() *sql.DB {
	// Abrir la conexión a la base de datos MySQL

	db, err := sql.Open("mysql", "root:Stori2024!!@tcp(db:3306)/StoriChallenge")
	if err != nil {
		log.Fatalf("Error connecting to the db: %v", err)
	}
	//defer db.Close()
	return db
}
