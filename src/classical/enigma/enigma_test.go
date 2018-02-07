package enigma_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"../../classical/enigma"
)


func TestEncDec_a(t *testing.T) {
	k := "aaa"
	ptext := "a"

	enigmaMachine := enigma.NewEnigma(nil, []rune(k))
	ctext := enigmaMachine.Encode(ptext)
	result := enigmaMachine.Decode(ctext)

	assert.Equal(t, ptext, result)
}

func TestEncDec_aa(t *testing.T) {
	k := "aaa"
	ptext := "aa"

	enigmaMachine := enigma.NewEnigma(nil, []rune(k))
	ctext := enigmaMachine.Encode(ptext)
	result := enigmaMachine.Decode(ctext)

	assert.Equal(t, ptext, result)
}

func TestEncDec_aaa(t *testing.T) {
	k := "aaa"
	ptext := "aa"

	enigmaMachine := enigma.NewEnigma(nil, []rune(k))
	ctext := enigmaMachine.Encode(ptext)
	result := enigmaMachine.Decode(ctext)

	assert.Equal(t, ptext, result)
}

func TestEncDec_a2(t *testing.T) {
	k := "aaa"
	ptext := "a"

	enigmaMachine := enigma.NewEnigma(nil, []rune(k))
	ctext := enigmaMachine.Encode(ptext)
	ctext = enigmaMachine.Encode(ctext)
	result := enigmaMachine.Decode(ctext)
	result = enigmaMachine.Decode(result)

	assert.Equal(t, ptext, result)
}

func TestEncDec_s(t *testing.T) {
	k := "aaa"
	ptext := "ciaociaociao"

	enigmaMachine := enigma.NewEnigma(nil, []rune(k))
	ctext := enigmaMachine.Encode(ptext)
	result := enigmaMachine.Decode(ctext)

	assert.Equal(t, ptext, result)
}

func TestEncDec_s3(t *testing.T) {
	k := "aaa"
	ptext := "ciaociaociao"

	enigmaMachine := enigma.NewEnigma(nil, []rune(k))
	ctext := enigmaMachine.Encode(ptext)
	ctext = enigmaMachine.Encode(ctext)
	ctext = enigmaMachine.Encode(ctext)

	result := enigmaMachine.Decode(ctext)
	result = enigmaMachine.Decode(result)
	result = enigmaMachine.Decode(result)

	assert.Equal(t, ptext, result)
}
