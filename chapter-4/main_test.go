package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShortfall(t *testing.T) {
	asia := sampleProvinceData()
	assert.Equal(t, 5, asia.getShortfall())
}
