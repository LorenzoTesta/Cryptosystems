package classical_test

import (
	"testing"
	"../classical"
	"github.com/stretchr/testify/assert"
)



func TestEncode_1a(t *testing.T) {
	p := "aaaaa"
	k := 1
	assert.Equal(t, classical.Encode (p, k), "bbbbb")
}

func TestEncode_1abc(t *testing.T) {
	p := "abc"
	k := 1
	assert.Equal(t, classical.Encode (p, k), "bcd")
}

func TestEncode_1az(t *testing.T) {
	p := "az"
	k := 1
	assert.Equal(t, classical.Encode (p, k), "ba")
}

func TestEncode_2a(t *testing.T) {
	p := "aaaaa"
	k := 2
	assert.Equal(t, classical.Encode (p, k), "ccccc")
}

func TestEncode_2abc(t *testing.T) {
	p := "abc"
	k := 2
	assert.Equal(t, classical.Encode (p, k), "cde")
}

func TestEncode_2az(t *testing.T) {
	p := "az"
	k := 2
	assert.Equal(t, classical.Encode (p, k), "cb")
}

func TestEncode_26az(t *testing.T) {
	p := "az"
	k := 26
	assert.Equal(t, classical.Encode (p, k), "az")
}



func TestDecode_1a(t *testing.T) {
	p := "aaaaa"
	k := 1
	assert.Equal(t, classical.Decode (p, k), "zzzzz")
}

func TestDecode_1abc(t *testing.T) {
	p := "abc"
	k := 1
	assert.Equal(t, classical.Decode (p, k), "zab")
}

func TestDecode_1az(t *testing.T) {
	p := "az"
	k := 1
	assert.Equal(t, classical.Decode (p, k), "zy")
}

func TestDecode_2a(t *testing.T) {
	p := "aaaaa"
	k := 2
	assert.Equal(t, classical.Decode (p, k), "yyyyy")
}

func TestDecode_2abc(t *testing.T) {
	p := "abc"
	k := 2
	assert.Equal(t, classical.Decode (p, k), "yza")
}

func TestDecode_2az(t *testing.T) {
	p := "az"
	k := 2
	assert.Equal(t, classical.Decode (p, k), "yx")
}

func TestDecode_26az(t *testing.T) {
	p := "az"
	k := 26
	assert.Equal(t, classical.Decode (p, k), "az")
}



func TestDecode(t *testing.T) {

}
