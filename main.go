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
	"sendEmails/email"
)

const layout = "2006/01/02"

var dataTransactions csvDataFile.DataTransactions

func init() {
	log.Printf("Starting ...")
}
func main() {
	conn, err := db.OpenConn()
	defer conn.Close()

	if err != nil {
		log.Fatalf("error connecting to the db %v", err)
		return
	}

	lines := strings.Split(csvDataFile.CsvData, "\r\n")

	rows := 0
	totalBalance := 0.0
	numTransactionsByMonth := make(map[string]int)
	avgCreditByMonth := make(map[string]float64)
	avgDebitByMonth := make(map[string]float64)
	avgTxCreditByMonth := make(map[string]int)
	avgTxDebitByMonth := make(map[string]int)

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
			log.Fatalf("Error parsing date: %v", err)
		}
		dataTransactions.CreationDate = date

		numTransactionsByMonth[date.Month().String()]++

		amount, err := calculations.GetAmount(columns[3])
		if err != nil {
			log.Print(err)
			continue
		}
		if strings.HasPrefix(columns[3], "+") {
			// option: use uint to the amounts.
			dataTransactions.Transaction.Credit = amount
			dataTransactions.Transaction.Debit = 0
			totalBalance += amount
			avgCreditByMonth[date.Month().String()] += amount
			avgTxCreditByMonth[date.Month().String()] += 1
			//fmt.Printf("el saldo es : %v %v %v %v %v \n", userId, txId, date, 0, credit)

			db.SaveTransaction(conn, dataTransactions)

		} else if strings.HasPrefix(columns[3], "-") {
			dataTransactions.Transaction.Credit = 0
			dataTransactions.Transaction.Debit = amount
			totalBalance -= amount
			avgDebitByMonth[date.Month().String()] -= amount
			avgTxDebitByMonth[date.Month().String()] += 1
			//fmt.Printf("el saldo es : %v %v %v %v %v\n", userId, txId, date, debit, 0)

			db.SaveTransaction(conn, dataTransactions)
		}

		// Calculate debit and credit avg by month
		for month := range numTransactionsByMonth {
			avgCreditByMonth[month] /= float64(numTransactionsByMonth[month])
			avgDebitByMonth[month] /= float64(numTransactionsByMonth[month])
		}

	}

	err = email.SendResumeByEmail(totalBalance, avgTxCreditByMonth, avgTxDebitByMonth, avgCreditByMonth, avgDebitByMonth)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	/*
				// Enviar resumen por correo electrónico
				err = enviarResumenPorCorreo(creditos, debitos, saldo)
				if err != nil {
					log.Fatalf("Error al enviar resumen por correo electrónico: %v", err)
				}
			echo "hola" >> /opt/emails/test.txt

		// Recorrer el mapa numTransactionsByMonth
			for month, numTransactions := range numTransactionsByMonth {
				log.Printf("Mes: %s, Número de transacciones: %d\n", month, numTransactions)
			}
	*/
}
