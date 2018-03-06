package utilities_test

import (
	"testing"
	. "../utilities"
	"github.com/stretchr/testify/assert"
)

/*
func TestFrequencyTable_NewFrequencyTable_0(t *testing.T) {
	ft := NewFrequencyTable("abbccddd")
	expected := FrequencyTable{{'a', 1}, {'b', 2}, {'c', 2}, {'d', 3}}
	assert.Equal(t, true, ft.Equal(expected))
}
*/

func TestFrequencyTable_GetSmallerThan_0(t *testing.T) {
	ft := FrequencyTable{{'b', 0.015}, {'c', 0.028},
						{'d', 0.043}, {'a', 0.082}}
	expected := []rune{'b','c','d'}
	assert.Equal(t, expected, ft.GetSmallerThan(0.050))
}


func TestFrequencyTable_CalcDeltas_0(t *testing.T) {
	ft := FrequencyTable{{'a', 0.082}, {'b', 0.015}, {'c', 0.028}, {'d', 0.043}}
	expected :=  FrequencyTable{{'a', 0.918}, {'b', 0.985}, {'c', 0.972}, {'d', 0.957}}
	assert.Equal(t, expected, ft.CalcDeltas(1))
}

func TestFrequencyTable_Equal(t *testing.T) {
	ft := FrequencyTable{{'a', 1}, {'b', 2}, {'c', 2}, {'d', 3}}
	ft2 := FrequencyTable{{'a', 1}, {'b', 2}, {'c', 2}, {'d', 3}}
	ft3 := FrequencyTable{{'a', 1}, {'b', 1}, {'c', 2}, {'d', 3}}
	ft4 := FrequencyTable{{'a', 1}, {'a', 2}, {'c', 2}, {'d', 3}}

	assert.Equal(t, true, ft.Equal(ft))
	assert.Equal(t, true, ft.Equal(ft2))
	assert.Equal(t, false, ft.Equal(ft3))
	assert.Equal(t, false, ft.Equal(ft4))
}


func TestValue_Equal(t *testing.T) {
	v := Value{'a', 1}
	v2 := Value{'a', 1}
	v3 := Value{'b', 1}
	v4 := Value{'a', 2}

	assert.Equal(t, true, v.Equal(&v))
	assert.Equal(t, true, v.Equal(&v2))
	assert.Equal(t,false, v.Equal(&v3))
	assert.Equal(t,false, v.Equal(&v4))
}