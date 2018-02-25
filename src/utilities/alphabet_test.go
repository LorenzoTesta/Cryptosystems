package utilities_test

import (
	"testing"
	. "../utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetLetter_0(t *testing.T) {
	c := 'a'
	shift := 0
	assert.Equal(t, 'a', GetLetter(c, shift))
}

func TestGetLetter_1(t *testing.T) {
	c := 'a'
	shift := 1
	assert.Equal(t, 'b', GetLetter(c, shift))
}

func TestGetLetter_2(t *testing.T) {
	c := 'a'
	shift := 2
	assert.Equal(t, 'c', GetLetter(c, shift))
}

func TestGetLetter_26(t *testing.T) {
	c := 'a'
	shift := 26
	assert.Equal(t, 'a', GetLetter(c, shift))
}

func TestGetLetter_m1(t *testing.T) {
	c := 'a'
	shift := -1
	assert.Equal(t, 'z', GetLetter(c, shift))
}

func TestGetLetter_m2(t *testing.T) {
	c := 'a'
	shift := -2
	assert.Equal(t, 'y', GetLetter(c, shift))
}

func TestGetAlphabetOffset_a(t *testing.T) {
	c := 'a'
	assert.Equal(t, 0, GetAlphabetOffset(c))
}

func TestGetAlphabetOffset_b(t *testing.T) {
	c := 'b'
	assert.Equal(t, 1, GetAlphabetOffset(c))
}

func TestGetAlphabetOffset_c(t *testing.T) {
	c := 'c'
	assert.Equal(t, 2, GetAlphabetOffset(c))
}

func TestGetAlphabetOffset_z(t *testing.T) {
	c := 'z'
	assert.Equal(t, 25, GetAlphabetOffset(c))
}

