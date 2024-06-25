package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set up the environment if necessary
	// For example, create temporary files or set environment variables

	// Change the current working directory to the root of the project
	// so that the program can find the JSON file
	if err := os.Chdir("../.."); err != nil {
		panic(err)
	}

	// Run the tests
	code := m.Run()

	// Clean up the environment if necessary
	// For example, remove temporary files or reset environment variables

	os.Exit(code)
}

var performances = []Performance{
	{PlayID: "hamlet", Audience: 55},
	{PlayID: "as-like", Audience: 35},
	{PlayID: "othello", Audience: 40},
}

var aPerformance = Performance{PlayID: "hamlet", Audience: 55}

func TestTotalAmount(t *testing.T) {
	expected := 173000
	result := totalAmount(performances)
	if result != expected {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestTotalVolumeCredits(t *testing.T) {
	expected := 47
	result := totalVolumeCredits(performances)
	if result != expected {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestVolumeCreditsFor(t *testing.T) {
	expected := 25
	result := volumeCreditsFor(aPerformance)
	if result != expected {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestPlayFor(t *testing.T) {
	expected := "Hamlet"
	result := playFor(aPerformance)
	if result.Name != expected {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestAmountFor(t *testing.T) {
	expected := 65000
	result, _ := amountFor(aPerformance)
	if result != expected {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestStatementHtml(t *testing.T) {
	invoice := Invoice{
		Customer:     "BigCo",
		Performances: performances,
	}

	expectedOutput := `<h1>Statement for BigCo</h1>
<table>
<tr><th>play</th><th>seats</th><th>cost</th></tr>
<tr><td>Hamlet</td><td>USD 650.00</td><td>55</td></tr>
</table>
<tr><td>As You Like It</td><td>USD 580.00</td><td>35</td></tr>
</table>
<tr><td>Othello</td><td>USD 500.00</td><td>40</td></tr>
</table>
<p>Amount owed is <em>USD 1,730.00</em></p><p>You earned <em>47</em> credits</p>
`

	result, err := statementHtml(invoice)
	if err != nil {
		t.Fatalf("statementHtml() returned an error: %v", err)
	}

	if result != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, result)
	}
}

func TestStatement(t *testing.T) {
	invoice := Invoice{
		Customer:     "BigCo",
		Performances: performances,
	}

	expectedOutput := `Statement for BigCo
Hamlet: USD 650.00 (55 seats) 
As You Like It: USD 580.00 (35 seats) 
Othello: USD 500.00 (40 seats) 
Amount owed is USD 1,730.00
You earned 47 credits
`

	result, err := statement(invoice)
	if err != nil {
		t.Fatalf("statement() returned an error: %v", err)
	}

	if result != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, result)
	}
}

func TestUnknownPlayType(t *testing.T) {
	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "unknown", Audience: 10},
		},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// panic expected
	statement(invoice)
}
