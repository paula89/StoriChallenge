package main

import (
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"sendEmails/calculations"
	csvDataFile "sendEmails/csv"
	"sendEmails/db"
)

const layout = "2006/01/02"

var dataTransactions csvDataFile.DataTransactions

func main() {
	conn := db.OpenConn()
	rows := 0
	totalBalance := 0.0
	numTransactionsByMonth := make(map[string]int)
	avgCreditByMonth := make(map[string]float64)
	avgDebitByMonth := make(map[string]float64)

	lines := strings.Split(csvDataFile.CsvData, "\r\n")
	for _, line := range lines {
		rows = rows + 1
		if line == "" || rows == 1 {
			continue
		}
		columns := strings.Split(line, ";")
		userId, err := uuid.Parse(columns[0])
		if err != nil {
			log.Fatalf("Error, invalid user uuid : %v %v", columns[0], err)
		}
		dataTransactions.UserId = userId
		txId, err := uuid.Parse(columns[1])
		if err != nil {
			log.Fatalf("Error, invalid tx uuid : %v %v", columns[1], err)
		}
		dataTransactions.Id = txId

		date, err := time.Parse(layout, columns[2])
		if err != nil {
			log.Fatalf("Error al parsear la fecha: %v", err)
		}
		dataTransactions.CreationDate = date
		numTransactionsByMonth[date.Month().String()]++
		if strings.HasPrefix(columns[3], "+") {
			// uint
			monto, err := calculations.ObtenerMonto(columns[3])
			if err != nil {
				log.Print(err)
				continue
			}
			dataTransactions.Transaction.Credit = monto
			dataTransactions.Transaction.Debit = 0
			totalBalance += monto
			avgCreditByMonth[date.Month().String()] += monto
			//fmt.Printf("el saldo es : %v %v %v %v %v \n", userId, txId, date, 0, credit)

			db.SaveTransaction(conn, dataTransactions.UserId, dataTransactions.Id, dataTransactions.CreationDate, dataTransactions.Transaction.Debit, dataTransactions.Transaction.Credit)

		} else if strings.HasPrefix(columns[3], "-") {
			monto, err := calculations.ObtenerMonto(columns[3])
			if err != nil {
				log.Print(err)
				continue
			}
			dataTransactions.Transaction.Credit = 0
			dataTransactions.Transaction.Debit = monto
			totalBalance -= monto
			avgDebitByMonth[date.Month().String()] -= monto
			//fmt.Printf("el saldo es : %v %v %v %v %v\n", userId, txId, date, debit, 0)

			//db.SaveTransaction(conn, dataTransactions.UserId, dataTransactions.Id, dataTransactions.CreationDate, dataTransactions.Transaction.Debit, dataTransactions.Transaction.Credit)
		}

		// Calcular promedios de débitos y créditos por mes
		for month := range numTransactionsByMonth {
			avgCreditByMonth[month] /= float64(numTransactionsByMonth[month])
			avgDebitByMonth[month] /= float64(numTransactionsByMonth[month])
		}

	}

	/*
		// Enviar resumen por correo electrónico
		err = enviarResumenPorCorreo(creditos, debitos, saldo)
		if err != nil {
			log.Fatalf("Error al enviar resumen por correo electrónico: %v", err)
		}

	*/
}
