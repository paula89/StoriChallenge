package email

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"sendEmails/calculations"
)

// getHtml represents the content of templateEmail.html
func getHtml() string {
	return "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Test email</title>\n</head>\n<body>\n    <h1>Resume bank account</h1>\n\n    <div class=\"jumbotron\">\n        <h1 class=\"display-4\">Bank account summary</h1>\n        <hr class=\"my-4\">\n        <div>\n            <ul>\n                <li>Total balance {{.totalBalance}}</li>\n                <li>{{.avg}}</li>\n            </ul>\n        </div>\n\n    </div>\n\n</body>\n</html>"
}

func SendResumeByEmail(results calculations.Results) {
	lines := getHtml()

	lines = strings.Replace(lines, "{{.totalBalance}}", strconv.FormatFloat(results.TotalBalance, 'f', -1, 64), -1)
	avg := ""

	log.Printf("avg ::: %v", results.AvgCreditByMonth)
	for month, numTransactions := range results.NumTransactionsByMonth {
		avg += "<li> Number of transactions in " + month + " " + strconv.Itoa(numTransactions) + "</li>"
	}
	for month, credit := range results.AvgCreditByMonth {
		avg += "<li> Total Average credit amount " + month + " " + strconv.FormatFloat(credit, 'f', -1, 64) + "</li>"
	}
	for month, debit := range results.AvgDebitByMonth {
		avg += "<li> Total Average debit amount " + month + " " + strconv.FormatFloat(debit, 'f', -1, 64) + "</li>"
	}

	lines = strings.Replace(lines, "<li>{{.avg}}</li>", avg, -1)

	log.Printf("new html email %v", lines)

	err := ioutil.WriteFile("outputEmail.html", []byte(lines), 0644)
	if err != nil {
		log.Fatalf("error writing the email %v", err)
	} else {
		log.Printf("Email sended %v", lines)
	}

}
