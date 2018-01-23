package vigenere_test

import (
	"testing"
	"../../classical/vigenere"
	"github.com/stretchr/testify/assert"
)

func TestEncode_1a(t *testing.T) {
	p := "aaaaa"
	k := []int{ 1, 2 }
	assert.Equal(t, "bcbcb", vigenere.Encode (p, k))
}

func TestEncode_1abc(t *testing.T) {
	p := "abc"
	k := []int{ 1, 2 }
	assert.Equal(t, "bdd", vigenere.Encode (p, k))
}

func TestEncode_1az(t *testing.T) {
	p := "az"
	k := []int{ 1, 2 }
	assert.Equal(t, "bb", vigenere.Encode (p, k))
}

func TestEncode_2a(t *testing.T) {
	p := "aaaaa"
	k := []int{ 1, 2 }
	assert.Equal(t, "bcbcb", vigenere.Encode (p, k))
}

func TestEncode_2abc(t *testing.T) {
	p := "abc"
	k := []int{ 1, 2 }
	assert.Equal(t, "bdd", vigenere.Encode (p, k))
}

func TestEncode_2az(t *testing.T) {
	p := "az"
	k := []int{ 1, 2 }
	assert.Equal(t, "bb", vigenere.Encode (p, k))
}

func TestEncode_26az(t *testing.T) {
	p := "az"
	k := []int{ 1, 2 }
	assert.Equal(t, "bb", vigenere.Encode (p, k))
}

func TestDecode_1a(t *testing.T) {
	c := "aaaaa"
	k := []int{ 1, 2 }
	assert.Equal(t, "zyzyz", vigenere.Decode (c, k))
}

func TestDecode_1abc(t *testing.T) {
	c := "abc"
	k := []int{ 1, 2 }
	assert.Equal(t, "zzb", vigenere.Decode (c, k))
}

func TestDecode_1az(t *testing.T) {
	c := "az"
	k := []int{ 1, 2 }
	assert.Equal(t, "zx", vigenere.Decode (c, k))
}

func TestDecode_2a(t *testing.T) {
	c := "aaaaa"
	k := []int{ 1, 2 }
	assert.Equal(t, "zyzyz", vigenere.Decode (c, k))
}

func TestDecode_2abc(t *testing.T) {
	c := "abc"
	k := []int{ 1, 2 }
	assert.Equal(t, "zzb", vigenere.Decode (c, k))
}

func TestDecode_2az(t *testing.T) {
	c := "az"
	k := []int{ 1, 2 }
	assert.Equal(t, "zx", vigenere.Decode (c, k))
}

func TestDecode_26az(t *testing.T) {
	c := "az"
	k := []int{ 1, 2 }
	assert.Equal(t, "zx", vigenere.Decode (c, k))
}



func TestComplete(t *testing.T) {
	p := "abcdefghijklmnopqrstuvwxyz"
	k := []int{ 1, 2 }
	c := vigenere.Encode(p, k)
	assert.Equal(t, p, vigenere.Decode(c, k))
}






