package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func save(body []byte) error {
	filename := "email/email" + ".html"
	return os.WriteFile(filename, body, 0600)
}

func SendResumeByEmail(totalBalance float64, avgTxCreditByMonth, avgTxDebitByMonth map[string]int, avgCreditByMonth, avgDebitByMonth map[string]float64) error {
	t, err := template.ParseFiles("email/templateEmail.html")
	if err != nil {
		log.Fatalf("error parsing file %v", err)
		return err
	}
	fmt.Printf("t %v", t)
	buf := new(bytes.Buffer)
	err = t.Execute(buf, map[string]interface{}{"data": totalBalance}) //strconv.FormatFloat(totalBalance, 'f', -1, 64))
	if err != nil {
		log.Fatalf("error adding variables %v", err)
		return err
	}
	err = save(buf.Bytes())
	if err != nil {
		log.Fatalf("error writing email %v", err)
		return err
	}
	fmt.Printf("Email Sended")
	return nil
}
