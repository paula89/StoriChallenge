package email

import (
	"fmt"
)

func SendResumeByEmail(creditos, debitos, saldo float64) error {

	fmt.Printf("Resumen de transacciones procesadas. Créditos: %.2f, Débitos: %.2f, Saldo final: %.2f\n", creditos, debitos, saldo)
	return nil
}
