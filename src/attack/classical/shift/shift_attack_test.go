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

/*

// GET_BEST

func TestGetBest(t *testing.T) {
	var countTable = map[rune]float64 {'a': 1.0, 'b': 7.0, 'c': 2.0, 'd': 2.5}
	result := shift.GetBest(shift.Rank(countTable), 2.4)
	assert.Equal(t, []rune{'a', 'c'}, result)
}

// CALC_DELTAS

func TestCalcDeltas(t *testing.T) {
	var countTable = map[rune]float64 {'a': 1.0, 'b': 1.5, 'c': 2.0, 'd': 2.5}
	result := shift.CalcDeltas(countTable, 1.75)
	prev := map[rune]float64 {'a': 0.75, 'b': 0.25, 'c': 0.25, 'd': 0.75}
	assert.Equal(t, prev, result)
}



// RANK

func TestRank(t *testing.T) {
	var countTable = map[rune]float64 {'a': 1.0, 'b': 7.0, 'c': 2.0, 'd': 2.5}
	result := shift.Rank(countTable)
	prev := make(shift.PairList, 4)
	prev[0] = shift.Pair{'a', 1.0}
	prev[1] = shift.Pair{'c', 2.0}
	prev[2] = shift.Pair{'d', 2.5}
	prev[3] = shift.Pair{'b', 7.0}
	assert.Equal(t, prev, result)
}
*/