package utilities_test

import (
	"testing"
	. "../utilities"
	"github.com/stretchr/testify/assert"
)


func TestGetSmallerThan_0(t *testing.T) {
	ol := FrequencyTable{{'b', 0.015}, {'c', 0.028},
						{'d', 0.043}, {'a', 0.082}}
	expected := []rune{'b','c','d'}
	assert.Equal(t, expected, ol.GetSmallerThan(0.050))
}

func TestCalcDeltas_0(t *testing.T) {
	ft := FrequencyTable{{'a', 0.082}, {'b', 0.015}, {'c', 0.028}, {'d', 0.043}}
	expected :=  FrequencyTable{{'a', 0.918}, {'b', 0.985}, {'c', 0.972}, {'d', 0.957}}
	assert.Equal(t, expected, ft.CalcDeltas(1))
}