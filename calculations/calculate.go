package calculations

import (
	"fmt"
	"strconv"
)

type Results struct {
	TotalBalance           float64
	NumTransactionsByMonth map[string]int
	AvgCreditByMonth       map[string]float64
	AvgDebitByMonth        map[string]float64
	AvgTxCreditByMonth     map[string]int
	AvgTxDebitByMonth      map[string]int
}

//GetAmount converts string to float
func GetAmount(transaction string) (float64, error) {
	tx := transaction[1:len(transaction)]
	amount, err := strconv.ParseFloat(tx, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot convert amount %s: %w", transaction, err)
	}
	//math.Round()
	return amount, nil
}
