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
	setup()
	tearDown := func() {
		// teardown code here
	}
	return tearDown
}

// This function is called before the test starts and after the test ends
func TestMain(m *testing.M) {
	// setup()
	fmt.Println("Before All tests")
	m.Run()
	// shutdown()
	fmt.Println("After All tests")
}

func TestProvince(t *testing.T) {
	t.Run("shortfall", func(t *testing.T) {
		defer setupTest()()
		assert.Equal(t, 5, asia.getShortfall())
	})

	t.Run("profit", func(t *testing.T) {
		defer setupTest()()
		assert.Equal(t, 230, asia.getProfit())
	})

	t.Run("change production", func(t *testing.T) {
		defer setupTest()()
		asia.producers[0].setProduction("20")
		assert.Equal(t, -6, asia.getShortfall())
		assert.Equal(t, 292, asia.getProfit())
	})

	t.Run("zero demand", func(t *testing.T) {
		defer setupTest()()
		asia.setDemand("0")
		assert.Equal(t, -25, asia.getShortfall())
		assert.Equal(t, 0, asia.getProfit())
	})

	t.Run("negative demand", func(t *testing.T) {
		defer setupTest()()
		asia.setDemand("-1")
		assert.Equal(t, -26, asia.getShortfall())
		assert.Equal(t, -10, asia.getProfit())
	})
}

func TestNoProducers(t *testing.T) {
	asia := &Province{name: "No Producers", demand: 30, price: 20}
	t.Run("shortfall", func(t *testing.T) {
		assert.Equal(t, 30, asia.getShortfall())
	})
	t.Run("profit", func(t *testing.T) {
		assert.Equal(t, 0, asia.getProfit())
	})
}
