package db

import (
	"database/sql"
	"log"

	csvDataFile "sendEmails/csv"
)

// SaveTransaction Save the data in a MySQL db
func SaveTransaction(db *sql.DB, dataTransaction csvDataFile.DataTransactions) {
	err := saveTxInDb(db, dataTransaction)
	if err != nil {
		log.Fatalf("Error saving transactions in the db: %v", err)
	}
}

func saveTxInDb(db *sql.DB, dataTransaction csvDataFile.DataTransactions) error {
	log.Printf("el dataTransaction es : %v", dataTransaction)
	_, err := db.Exec("INSERT INTO StoriChallenge.Transactions (Id, UserId, CreationDate, Debit, Credit) VALUES (?, ?, ?, ?, ?)",
		dataTransaction.Id, dataTransaction.UserId, dataTransaction.CreationDate, dataTransaction.Transaction.Debit, dataTransaction.Transaction.Credit)

	return err
}

// OpenConn Open the db connection
func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Stori2024!!@tcp(db:3306)/StoriChallenge")
	if err != nil {
		log.Fatalf("Error connecting to the db: %v", err)
		return nil, err
	}
	/*
		It's not possible do db.Ping() by conflict with mac port (53)
			err = db.Ping()
			if err != nil {
				log.Printf("Error conectando: %v", err)
				return nil
			}
	*/
	log.Printf("Connection sucessfully")
	return db, nil
}
