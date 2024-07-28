package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var asia *Province

func setup() {
	asia = sampleProvinceData() // Create a new Province object before each test to avoid side effects
}

func setupTest() func() {
	// Setup code here
	setup()
	// tear down later
	return func() {
		// tear-down code here
	}
}

// This function is called before the test starts and after the test ends
func TestMain(m *testing.M) {
	// setup()
	fmt.Println("Before All tests")
	m.Run()
	// shutdown()
	fmt.Println("After All tests")
}

func TestShortfall(t *testing.T) {
	defer setupTest()()
	assert.Equal(t, 5, asia.getShortfall())
}

func TestProfit(t *testing.T) {
	defer setupTest()()
	assert.Equal(t, 230, asia.getProfit())
}

func TestChangeProduction(t *testing.T) {
	defer setupTest()()
	asia.producers[0].setProduction("20")
	assert.Equal(t, -6, asia.getShortfall())
	assert.Equal(t, 292, asia.getProfit())
}
