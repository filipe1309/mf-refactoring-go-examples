package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var asia = sampleProvinceData()

func TestShortfall(t *testing.T) {
	assert.Equal(t, 5, asia.getShortfall())
}

func TestProfit(t *testing.T) {
	assert.Equal(t, 230, asia.getProfit())
}
