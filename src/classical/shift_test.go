package classical_test

import (
	"testing"
	"../classical"
	"github.com/stretchr/testify/assert"
)



func TestEncode_1a(t *testing.T) {
	p := "aaaaa"
	k := 1
	assert.Equal(t, "bbbbb", classical.Encode (p, k))
}

func TestEncode_1abc(t *testing.T) {
	p := "abc"
	k := 1
	assert.Equal(t, "bcd", classical.Encode (p, k))
}

func TestEncode_1az(t *testing.T) {
	p := "az"
	k := 1
	assert.Equal(t, "ba", classical.Encode (p, k))
}

func TestEncode_2a(t *testing.T) {
	p := "aaaaa"
	k := 2
	assert.Equal(t, "ccccc", classical.Encode (p, k))
}

func TestEncode_2abc(t *testing.T) {
	p := "abc"
	k := 2
	assert.Equal(t, "cde", classical.Encode (p, k))
}

func TestEncode_2az(t *testing.T) {
	p := "az"
	k := 2
	assert.Equal(t, "cb", classical.Encode (p, k))
}

func TestEncode_26az(t *testing.T) {
	p := "az"
	k := 26
	assert.Equal(t, "az", classical.Encode (p, k))
}

func TestDecode_1a(t *testing.T) {
	c := "aaaaa"
	k := 1
	assert.Equal(t, "zzzzz", classical.Decode (c, k))
}

func TestDecode_1abc(t *testing.T) {
	c := "abc"
	k := 1
	assert.Equal(t, "zab", classical.Decode (c, k))
}

func TestDecode_1az(t *testing.T) {
	c := "az"
	k := 1
	assert.Equal(t, "zy", classical.Decode (c, k))
}

func TestDecode_2a(t *testing.T) {
	c := "aaaaa"
	k := 2
	assert.Equal(t, "yyyyy", classical.Decode (c, k))
}

func TestDecode_2abc(t *testing.T) {
	c := "abc"
	k := 2
	assert.Equal(t, "yza", classical.Decode (c, k))
}

func TestDecode_2az(t *testing.T) {
	c := "az"
	k := 2
	assert.Equal(t, "yx", classical.Decode (c, k))
}

func TestDecode_26az(t *testing.T) {
	c := "az"
	k := 26
	assert.Equal(t, "az", classical.Decode (c, k))
}



func TestComplete(t *testing.T) {
	p := "abcdefghijklmnopqrstuvwxyz"
	k := 21
	c := classical.Encode (p, k)
	assert.Equal(t, p, classical.Decode (c, k))
}