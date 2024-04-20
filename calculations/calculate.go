package calculations

import (
	"fmt"
	"strconv"
)

func ObtenerMonto(transaction string) (float64, error) {
	tx := transaction[1:len(transaction)]
	monto, err := strconv.ParseFloat(tx, 64)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir el monto %s: %w", transaction, err)
	}
	//math.Round()
	return monto, nil
}