package utilities_test

import (
	"testing"
	"../utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetLetter_0(t *testing.T) {
	c := 'a'
	shift := 0
	assert.Equal(t, 'a', utilities.GetLetter(c, shift))
}

func TestGetLetter_1(t *testing.T) {
	c := 'a'
	shift := 1
	assert.Equal(t, 'b', utilities.GetLetter(c, shift))
}

func TestGetLetter_2(t *testing.T) {
	c := 'a'
	shift := 2
	assert.Equal(t, 'c', utilities.GetLetter(c, shift))
}

func TestGetLetter_26(t *testing.T) {
	c := 'a'
	shift := 26
	assert.Equal(t, 'a', utilities.GetLetter(c, shift))
}

func TestGetLetter_m1(t *testing.T) {
	c := 'a'
	shift := -1
	assert.Equal(t, 'z', utilities.GetLetter(c, shift))
}

func TestGetLetter_m2(t *testing.T) {
	c := 'a'
	shift := -2
	assert.Equal(t, 'y', utilities.GetLetter(c, shift))
}
