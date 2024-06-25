package main

import (
	"os"
	"strings"
	"testing"
)

func TestAmountFor(t *testing.T) {
	performance := Performance{PlayID: "hamlet", Audience: 55}
	expected := 65000
	result, _ := amountFor(performance)
	if result != expected {
		t.Errorf("expected %v but got %v", expected, result)
	}
}

func TestStatement(t *testing.T) {
	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		},
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

func TestUnknownPlayType(t *testing.T) {
	invoice := Invoice{
		Customer: "BigCo",
		Performances: []Performance{
			{PlayID: "unknown", Audience: 10},
		},
	}

	_, err := statement(invoice)
	if err == nil {
		t.Fatalf("expected an error but got nil")
	}

	if !strings.Contains(err.Error(), "unknown performance type") {
		t.Errorf("expected error message to contain 'unknown performance type' but got %v", err.Error())
	}
}
