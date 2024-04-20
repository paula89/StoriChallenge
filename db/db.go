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
		log.Fatalf("Error al guardar transacciones en la base de datos: %v", err)
	}
}

func saveTxInDb(db *sql.DB, UserId, TxId uuid.UUID, Date time.Time, Debit, Credit float64) error {
	_, err := db.Exec("INSERT INTO transactions (Id, UserId, Date, Debit, Credit) VALUES (?, ?, ?)", TxId, UserId, Date, Debit, Credit)
	return err
}

func OpenConn() *sql.DB {
	// Abrir la conexión a la base de datos MySQL

	db, err := sql.Open("mysql", "admin:password!!@tcp(rds-mysql.database-stori.cpy4qmq0clr0.us-east-2.rds.amazonaws.com:3306)/transactions")
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	//defer db.Close()
	return db
}
