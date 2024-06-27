package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type StatementData struct {
	Customer           string
	Performances       []Performance
	TotalAmount        int
	TotalVolumeCredits int
}

type Calculator interface {
	amount() (int, error)
	volumeCredits() int
	play() Play
}

type TragedyCalculator struct {
	p *PerformanceCalculator
}

func (t *TragedyCalculator) amount() (int, error) {
	result := 40000
	if t.p.aPerformance.Audience > 30 {
		result += 1000 * (t.p.aPerformance.Audience - 30)
	}
	return result, nil
}

func (t *TragedyCalculator) volumeCredits() int {
	return int(math.Max(float64(t.p.aPerformance.Audience)-30, 0))
}

func (t *TragedyCalculator) play() Play {
	return t.p.aPlay
}

type ComedyCalculator struct {
	p *PerformanceCalculator
}

func (c *ComedyCalculator) amount() (int, error) {
	result := 30000
	if c.p.aPerformance.Audience > 20 {
		result += 10000 + 500*(c.p.aPerformance.Audience-20)
	}
	result += 300 * c.p.aPerformance.Audience
	return result, nil
}

func (c *ComedyCalculator) volumeCredits() int {
	result := int(math.Max(float64(c.p.aPerformance.Audience)-30, 0))
	result += int(math.Floor(float64(c.p.aPerformance.Audience) / 5))
	return result
}

func (c *ComedyCalculator) play() Play {
	return c.p.aPlay
}

type PerformanceCalculator struct {
	aPerformance Performance
	aPlay        Play
}

func createPerformanceCalculator(aPerformance Performance) Calculator {
	aPlay := playFor(aPerformance)
	switch aPlay.Type {
	case "tragedy":
		return &TragedyCalculator{&PerformanceCalculator{aPerformance, aPlay}}
	case "comedy":
		return &ComedyCalculator{&PerformanceCalculator{aPerformance, aPlay}}
	default:
		panic(fmt.Sprintf("unknown type: %s", aPlay.Type))
	}
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
	var calculator = createPerformanceCalculator(aPerformance)
	aPerformance.Play = calculator.play()
	amount, err := calculator.amount()
	if err != nil {
		return nil, err
	}
	aPerformance.Amount = amount
	aPerformance.VolumeCredits = calculator.volumeCredits()
	return &aPerformance, nil
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
