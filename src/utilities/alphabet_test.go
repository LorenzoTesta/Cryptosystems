package utilities_test

import (
	"testing"
	"../utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetLetter_0(t *testing.T) {
	c := 'a'
	shift := 0
	assert.Equal(t, utilities.GetLetter(c, shift), 'a')
}

func TestGetLetter_1(t *testing.T) {
	c := 'a'
	shift := 1
	assert.Equal(t, utilities.GetLetter(c, shift), 'b')
}

func TestGetLetter_2(t *testing.T) {
	c := 'a'
	shift := 2
	assert.Equal(t, utilities.GetLetter(c, shift), 'c')
}

func TestGetLetter_26(t *testing.T) {
	c := 'a'
	shift := 26
	assert.Equal(t, utilities.GetLetter(c, shift), 'a')
}

func TestGetLetter_m1(t *testing.T) {
	c := 'a'
	shift := -1
	assert.Equal(t, utilities.GetLetter(c, shift), 'z')
}

func TestGetLetter_m2(t *testing.T) {
	c := 'a'
	shift := -2
	assert.Equal(t, utilities.GetLetter(c, shift), 'y')
}
