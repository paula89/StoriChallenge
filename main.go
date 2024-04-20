package main

import (
	"fmt"
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

func main() {
	conn := db.OpenConn()
	rows := 0
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
		txId, err := uuid.Parse(columns[1])
		if err != nil {
			log.Fatalf("Error, invalid tx uuid : %v %v", columns[1], err)
		}
		date, err := time.Parse(layout, columns[2])
		if err != nil {
			log.Fatalf("Error al parsear la fecha: %v", err)
		}

		if strings.HasPrefix(columns[3], "+") {
			monto, err := calculations.ObtenerMonto(columns[3])
			if err != nil {
				log.Print(err)
				continue
			}
			credit := monto
			fmt.Printf("el saldo es : %v %v %v %v %v \n", userId, txId, date, 0, credit)

			db.SaveTransaction(conn, userId, txId, date, 0, credit)

		} else if strings.HasPrefix(columns[3], "-") {
			monto, err := calculations.ObtenerMonto(columns[3])
			if err != nil {
				log.Print(err)
				continue
			}
			debit := monto
			fmt.Printf("el saldo es : %v %v %v %v %v\n", userId, txId, date, debit, 0)

			//saveTransaction(conn, uuid.MustParse(userId), uuid.MustParse(txId), date, debit, 0)
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
