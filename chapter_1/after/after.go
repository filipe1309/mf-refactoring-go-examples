// Imagine a company of theatrical players who go out to various
// events performing plays. Typically a customer will request a play and the
// company charges them based on the size of the audience and the kind
// of play they perform. There are currently two kinds of plays that
// the players perform: tragedies and comedies. As well as providing a
// bill for the performance, the company gives is it's customers
// "volume credits", a loyalty mechanism they can use for discounts on
// future performances.

// The performers store data about their plays in a JSON file called
// plays.json. They store data for their bills in a file called invoices.json
// The code that prints the bill is a function called statement.

// Running this code on the test data files results in the following output:
// Hamlet: USD 650.00 (55 seats)
// As You Like It: USD 580.00 (35 seats)
// Othello: USD 500.00 (40 seats)
// Amount owed is USD 1730.00
// You earned 47 credits

package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/currency"
	"math"
	"os"
	"strings"
)

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Performance struct {
	PlayID        string `json:"playID"`
	Audience      int    `json:"audience"`
	Play          Play
	Amount        int
	VolumeCredits int
}

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type StatementData struct {
	Customer           string
	Performances       []Performance
	TotalAmount        int
	TotalVolumeCredits int
}

type PerformanceCalculator struct {
	aPerformance Performance
	aPlay        Play
}

func (p *PerformanceCalculator) amount() (int, error) {
	var result int
	switch p.aPlay.Type {
	case "tragedy":
		result = 40000
		if p.aPerformance.Audience > 30 {
			result += 1000 * (p.aPerformance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if p.aPerformance.Audience > 20 {
			result += 10000 + 500*(p.aPerformance.Audience-20)
		}
		result += 300 * p.aPerformance.Audience

	default:
		return result, fmt.Errorf("error unknown performance type %s", p.aPerformance.Play.Type)
	}

	return result, nil
}

func main() {
	invoiceFile, err := os.ReadFile("chapter_1/after/invoices.json")
	if err != nil {
		fmt.Println(err)
	}

	var invoice Invoice
	if err := json.Unmarshal(invoiceFile, &invoice); err != nil {
		fmt.Println(err)
	}

	result, err := statement(invoice)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	resultHtml, err := statementHtml(invoice)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultHtml)
}

func usd(amount int) string {
	return fmt.Sprintf("%+v", currency.USD.Amount(amount/100))
}

func statement(invoice Invoice) (string, error) {
	return renderPlainText(createStatementData(invoice))
}

func createStatementData(invoice Invoice) StatementData {
	var enrichedPerformances []Performance
	for _, perf := range invoice.Performances {
		enrichedPerformance, err := enrichPerformance(perf)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		enrichedPerformances = append(enrichedPerformances, *enrichedPerformance)
	}

	statementData := StatementData{
		Customer:           invoice.Customer,
		Performances:       enrichedPerformances,
		TotalAmount:        totalAmount(enrichedPerformances),
		TotalVolumeCredits: totalVolumeCredits(enrichedPerformances),
	}
	return statementData
}

func enrichPerformance(aPerformance Performance) (*Performance, error) {
	var calculator = PerformanceCalculator{aPerformance: aPerformance, aPlay: playFor(aPerformance)}
	aPerformance.Play = calculator.aPlay
	amount, err := amountFor(aPerformance)
	if err != nil {
		return nil, err
	}
	aPerformance.Amount = amount
	aPerformance.VolumeCredits = volumeCreditsFor(aPerformance)
	return &aPerformance, nil
}

func statementHtml(invoice Invoice) (string, error) {
	return renderHtml(createStatementData(invoice))
}

func renderHtml(statementData StatementData) (string, error) {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("<h1>Statement for %s</h1>\n", statementData.Customer))
	result.WriteString("<table>\n")
	result.WriteString("<tr><th>play</th><th>seats</th><th>cost</th></tr>\n")

	for _, perf := range statementData.Performances {
		result.WriteString(fmt.Sprintf("<tr><td>%s</td>", perf.Play.Name))
		result.WriteString(fmt.Sprintf("<td>%s</td>", usd(perf.Amount)))
		result.WriteString(fmt.Sprintf("<td>%d</td></tr>\n", perf.Audience))
		result.WriteString("</table>\n")
	}
	result.WriteString(fmt.Sprintf("<p>Amount owed is <em>%s</em></p>", usd(statementData.TotalAmount)))
	result.WriteString(fmt.Sprintf("<p>You earned <em>%d</em> credits</p>\n", statementData.TotalVolumeCredits))
	return result.String(), nil
}

func renderPlainText(statementData StatementData) (string, error) {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("Statement for %s\n", statementData.Customer))

	for _, perf := range statementData.Performances {
		result.WriteString(fmt.Sprintf("%s: %s (%d seats) \n", perf.Play.Name, usd(perf.Amount), perf.Audience))
	}

	result.WriteString(fmt.Sprintf("Amount owed is %s\n", usd(statementData.TotalAmount)))
	result.WriteString(fmt.Sprintf("You earned %d credits\n", statementData.TotalVolumeCredits))
	return result.String(), nil
}

func amountFor(aPerformance Performance) (int, error) {
  return (&PerformanceCalculator{aPerformance: aPerformance, aPlay: playFor(aPerformance)}).amount()
}

func playFor(aPerformance Performance) Play {
	playsFile, err := os.ReadFile("chapter_1/after/plays.json")
	if err != nil {
		fmt.Println(err)
	}

	var plays map[string]Play
	if err := json.Unmarshal(playsFile, &plays); err != nil {
		fmt.Println(err)
	}
	return plays[aPerformance.PlayID]
}

func volumeCreditsFor(aPerformance Performance) int {
	result := int(math.Max(float64(aPerformance.Audience)-30, 0))
	if aPerformance.Play.Type == "comedy" {
		result += int(math.Floor(float64(aPerformance.Audience) / 5))
	}
	return result
}

func totalVolumeCredits(performances []Performance) int {
	result := 0
	for _, perf := range performances {
		result += perf.VolumeCredits
	}
	return result
}

func totalAmount(performances []Performance) int {
	result := 0
	for _, perf := range performances {
		thisAmount := perf.Amount
		result += thisAmount
	}
	return result
}
