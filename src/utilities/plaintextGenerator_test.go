package utilities_test

import (
	"testing"
. "../utilities"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePlaintext_0(t *testing.T) {
	ft := FrequencyTable{{'a', 1}, {'b', 2}, {'c', 2}, {'d', 3}}
	expected := "abbccddd"
	assert.Equal(t, expected, GeneratePlaintext(&ft))
}
