package shift_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	algo "../../../classical/shift"
	"../../../attack/classical/shift"
	. "../../../utilities"
	"fmt"
)

var FILEPATH = "../../../test_files/dummy.txt"


// EXHAUSTIVE

func TestExhaustive(t *testing.T) {
	p := "abcdefg"
	c := algo.Encode(p, 7)
	results := shift.Exhaustive(c)
	assert.Contains(t, results, p)
}

// LETTER FREQUENCY

func TestLetterFrequency(t *testing.T) {
	ft := FrequencyTable{{'b', 0.1}, {'c', 0.2},
		{'d', 0.3}, {'a', 0.4}}
	p := GeneratePlaintext(&ft)
	fmt.Println("plaintext:", p)
	results := shift.LetterFrequency(p, &ft)
	PrintPretty(results)
	assert.Contains(t, results, p)
}
