package csv

import (
	_ "embed"
	"time"

	"github.com/google/uuid"
)

//go:embed txns.csv
var CsvData string

type DataTransactions struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	CreationDate time.Time
	Transaction  Transaction
}

type Transaction struct {
	Credit float64
	Debit  float64
}
