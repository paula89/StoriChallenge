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
var results calculations.Results

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

	results.NumTransactionsByMonth = make(map[string]int)
	results.AvgCreditByMonth = make(map[string]float64)
	results.AvgDebitByMonth = make(map[string]float64)
	results.AvgTxCreditByMonth = make(map[string]int)
	results.AvgTxDebitByMonth = make(map[string]int)

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

		results.NumTransactionsByMonth[date.Month().String()]++

		amount, err := calculations.GetAmount(columns[3])
		if err != nil {
			log.Fatalf("error getting the amount %v", err)
		}
		if strings.HasPrefix(columns[3], "+") {
			// option: use uint to the amounts.
			dataTransactions.Transaction.Credit = amount
			dataTransactions.Transaction.Debit = 0
			results.TotalBalance += amount
			results.AvgCreditByMonth[date.Month().String()] += amount
			results.AvgTxCreditByMonth[date.Month().String()] += 1

			db.SaveTransaction(conn, dataTransactions)

		} else if strings.HasPrefix(columns[3], "-") {
			dataTransactions.Transaction.Credit = 0
			dataTransactions.Transaction.Debit = amount
			results.TotalBalance -= amount
			results.AvgDebitByMonth[date.Month().String()] -= amount
			results.AvgTxDebitByMonth[date.Month().String()] += 1

			db.SaveTransaction(conn, dataTransactions)
		}

		// Calculate debit and credit avg by month
		for month := range results.NumTransactionsByMonth {
			results.AvgCreditByMonth[month] /= float64(results.AvgTxCreditByMonth[month])
			results.AvgDebitByMonth[month] /= float64(results.AvgTxDebitByMonth[month])
		}
	}

	email.SendResumeByEmail(results)

}
