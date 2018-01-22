package shift_test

import (
	"testing"
	"../../classical/shift"
	"github.com/stretchr/testify/assert"
)



func TestEncode_1a(t *testing.T) {
	p := "aaaaa"
	k := 1
	assert.Equal(t, "bbbbb", shift.Encode (p, k))
}

func TestEncode_1abc(t *testing.T) {
	p := "abc"
	k := 1
	assert.Equal(t, "bcd", shift.Encode (p, k))
}

func TestEncode_1az(t *testing.T) {
	p := "az"
	k := 1
	assert.Equal(t, "ba", shift.Encode (p, k))
}

func TestEncode_2a(t *testing.T) {
	p := "aaaaa"
	k := 2
	assert.Equal(t, "ccccc", shift.Encode (p, k))
}

func TestEncode_2abc(t *testing.T) {
	p := "abc"
	k := 2
	assert.Equal(t, "cde", shift.Encode (p, k))
}

func TestEncode_2az(t *testing.T) {
	p := "az"
	k := 2
	assert.Equal(t, "cb", shift.Encode (p, k))
}

func TestEncode_26az(t *testing.T) {
	p := "az"
	k := 26
	assert.Equal(t, "az", shift.Encode (p, k))
}

func TestDecode_1a(t *testing.T) {
	c := "aaaaa"
	k := 1
	assert.Equal(t, "zzzzz", shift.Decode (c, k))
}

func TestDecode_1abc(t *testing.T) {
	c := "abc"
	k := 1
	assert.Equal(t, "zab", shift.Decode (c, k))
}

func TestDecode_1az(t *testing.T) {
	c := "az"
	k := 1
	assert.Equal(t, "zy", shift.Decode (c, k))
}

func TestDecode_2a(t *testing.T) {
	c := "aaaaa"
	k := 2
	assert.Equal(t, "yyyyy", shift.Decode (c, k))
}

func TestDecode_2abc(t *testing.T) {
	c := "abc"
	k := 2
	assert.Equal(t, "yza", shift.Decode (c, k))
}

func TestDecode_2az(t *testing.T) {
	c := "az"
	k := 2
	assert.Equal(t, "yx", shift.Decode (c, k))
}

func TestDecode_26az(t *testing.T) {
	c := "az"
	k := 26
	assert.Equal(t, "az", shift.Decode (c, k))
}



func TestComplete(t *testing.T) {
	p := "abcdefghijklmnopqrstuvwxyz"
	k := 21
	c := shift.Encode (p, k)
	assert.Equal(t, p, shift.Decode (c, k))
}