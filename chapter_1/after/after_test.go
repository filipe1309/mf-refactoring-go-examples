package main

import (
    "testing"
    "os"
    "strings"
)

func TestStatement(t *testing.T) {
    plays := map[string]Play{
        "hamlet":   {Name: "Hamlet", Type: "tragedy"},
        "as-like":  {Name: "As You Like It", Type: "comedy"},
        "othello":  {Name: "Othello", Type: "tragedy"},
    }

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

    result, err := statement(invoice, plays)
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

    // Run the tests
    code := m.Run()

    // Clean up the environment if necessary
    // For example, remove temporary files or reset environment variables

    os.Exit(code)
}

func TestUnknownPlayType(t *testing.T) {
    plays := map[string]Play{
        "unknown": {Name: "Unknown Play", Type: "unknown"},
    }

    invoice := Invoice{
        Customer: "BigCo",
        Performances: []Performance{
            {PlayID: "unknown", Audience: 10},
        },
    }

    _, err := statement(invoice, plays)
    if err == nil {
        t.Fatalf("expected an error but got nil")
    }

    if !strings.Contains(err.Error(), "unknown performance type") {
        t.Errorf("expected error message to contain 'unknown performance type' but got %v", err.Error())
    }
}
