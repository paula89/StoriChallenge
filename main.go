package main

import (
	"database/sql"
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	csvDataFile "sendEmails/csv"
)

var layout = "2006/01/02"
var rows = 0

func main() {
	//db := openConn()

	lines := strings.Split(csvDataFile.CsvData, ":")
	for _, line := range lines {
		rows = rows + 1
		if line == "" || rows == 1 {
			continue
		}
		columns := strings.Split(line, ";")
		//fmt.Printf("el columns es : %v %v", uuid.New(), uuid.New())
		userId, err := uuid.Parse(columns[0])

		if err != nil {
			log.Fatalf("Error, invalid uuid : %v %v", columns[0], err)
		}
		txId, err := uuid.Parse(columns[1])
		if err != nil {
			log.Fatalf("Error, invalid uuid : %v %v", columns[1], err)
		}
		date, err := time.Parse(layout, columns[2])
		if err != nil {
			log.Fatalf("Error al parsear la fecha: %v", err)
		}

		if strings.HasPrefix(columns[3], "+") {
			monto, err := obtenerMonto(columns[3])
			if err != nil {
				log.Printf("Error al procesar el monto: %v", err)
				continue
			}
			credit := monto
			fmt.Printf("el saldo es : %v %v %v %v %v", userId, txId, date, 0, credit)

			//saveTransaction(db, uuid.MustParse(userId), uuid.MustParse(txId), date, 0, credit)

		} else if strings.HasPrefix(columns[3], "-") {
			monto, err := obtenerMonto(columns[3])
			if err != nil {
				log.Printf("Error al procesar el monto: %v", err)
				continue
			}
			debit := monto
			fmt.Printf("el saldo es : %v %v %v %v %v", userId, txId, date, debit, 0)

			//saveTransaction(db, uuid.MustParse(userId), uuid.MustParse(txId), date, debit, 0)
		}

	}

	// Calcular el saldo final
	//saldo := creditos - debitos
	//fmt.Printf("el saldo es : %v", saldo)

	/*
		// Enviar resumen por correo electrónico
		err = enviarResumenPorCorreo(creditos, debitos, saldo)
		if err != nil {
			log.Fatalf("Error al enviar resumen por correo electrónico: %v", err)
		}

	*/
}

func obtenerMonto(transaction string) (float64, error) {
	tx := transaction[1:len(transaction)]
	monto, err := strconv.ParseFloat(tx, 64)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir el monto %s: %w", transaction, err)
	}
	//math.Round()
	return monto, nil
}

func enviarResumenPorCorreo(creditos, debitos, saldo float64) error {
	// Aquí iría la lógica para enviar el correo electrónico con el resumen
	fmt.Printf("Resumen de transacciones procesadas. Créditos: %.2f, Débitos: %.2f, Saldo final: %.2f\n", creditos, debitos, saldo)
	return nil
}

func saveTransaction(db *sql.DB, UserId, TxId uuid.UUID, Date time.Time, Debit, Credit float64) {
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

func openConn() *sql.DB {
	// Abrir la conexión a la base de datos MySQL

	db, err := sql.Open("mysql", "admin:Stori2024!!@tcp(rds-mysql.database-stori.cpy4qmq0clr0.us-east-2.rds.amazonaws.com:3306)/transactions")
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()
	return db
}

func loadCsv() {
	r := csv.NewReader(strings.NewReader(csvDataFile.CsvData))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error al leer el archivo CSV: %v", err)
	}

	for _, row := range records {
		fmt.Println(row)
	}

}
