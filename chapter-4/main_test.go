package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortfall(t *testing.T) {
	asia := sampleProvinceData()
	assert.Equal(t, 5, asia.getShortfall())
}

func TestProfit(t *testing.T) {
	asia := sampleProvinceData()
	assert.Equal(t, 230, asia.getProfit())
}
