// Chapter 4 of Refactoring book
package main

import (
	"fmt"
	"math"
	"strconv"
)

type Province struct {
	name            string
	producers       []Producer
	totalProduction int
	demand          int
	price           int
}

func (p *Province) getName() string {
	return p.name
}

func (p *Province) getProducers() []Producer {
	return p.producers
}

func (p *Province) getTotalProduction() int {
	return p.totalProduction
}

func (p *Province) getDemand() int {
	return p.demand
}

func (p *Province) setDemand(demand string) error {
	i, err := strconv.Atoi(demand)
	if err != nil {
		return err
	}
	p.demand = i
	return nil
}

func (p *Province) getPrice() int {
	return p.price
}
func (p *Province) setPrice(price string) {
	i, err := strconv.Atoi(price)
	if err != nil {
		panic(err)
	}
	p.price = i
}

func (p *Province) getShortfall() int {
	return p.demand - p.totalProduction
}

func (p *Province) getProfit() int {
	return p.demandValue() - p.demandCost()
}

func (p *Province) demandCost() int {
	var remainingDemand = p.demand
	var result int

	// sort producers by cost
	sortedProducers := make([]Producer, len(p.producers))
	copy(sortedProducers, p.producers)
	for i := 0; i < len(sortedProducers); i++ {
		for j := i + 1; j < len(sortedProducers); j++ {
			if sortedProducers[i].cost > sortedProducers[j].cost {
				sortedProducers[i], sortedProducers[j] = sortedProducers[j], sortedProducers[i]
			}
		}
	}

	// calculate cost
	for _, producer := range sortedProducers {
		contribution := int(math.Min(float64(remainingDemand), float64(producer.production)))
		remainingDemand -= contribution
		result += contribution * producer.cost
	}

	return result
}

func (p *Province) demandValue() int {
	return p.satisfiedDemand() * p.price
}

func (p *Province) satisfiedDemand() int {
	return int(math.Min(float64(p.demand), float64(p.totalProduction)))
}

func (p *Province) addProducer(producer Producer) {
	p.producers = append(p.producers, producer)
	p.totalProduction += producer.production
}

type Producer struct {
	name       string
	cost       int
	production int
	province   *Province
}

func (p *Producer) getName() string {
	return p.name
}

func (p *Producer) getCost() int {
	return p.cost
}

func (p *Producer) setCost(cost string) {
	i, err := strconv.Atoi(cost)
	if err != nil {
		panic(err)
	}
	p.cost = i
}

func (p *Producer) getProduction() int {
	return p.production
}

func (p *Producer) setProduction(amountStr string) {
	amount, err := strconv.Atoi(amountStr)
	var newProduction int
	if err == nil {
		newProduction = amount
	}
	p.province.totalProduction += newProduction - p.production
	p.production = newProduction
}

func (p *Producer) getProvince() Province {
	return *p.province
}

func sampleProvinceData() *Province {
	province := &Province{name: "Asia"}
	province.addProducer(Producer{name: "Byzantium", cost: 10, production: 9, province: province})
	province.addProducer(Producer{name: "Attalia", cost: 12, production: 10, province: province})
	province.addProducer(Producer{name: "Sinope", cost: 10, production: 6, province: province})
	province.setDemand("30")
	province.setPrice("20")
	return province
}

func main() {
	fmt.Println("Production Plan")
	asia := sampleProvinceData()
	fmt.Println("\nProvince:", asia.getName())
	fmt.Println("\tShortfall:\t", asia.getShortfall())
	fmt.Println("\tProfit:\t\t", asia.getProfit())
}
