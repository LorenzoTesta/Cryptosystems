package shift_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	algo "../../../classical/shift"
	"../../../attack/classical/shift"
	. "../../../utilities"
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

func TestLetterFrequency_0(t *testing.T) {
	ft := FrequencyTable{{'b', 0.1}, {'c', 0.2},
		{'d', 0.3}, {'a', 0.4}}
	p := GeneratePlaintext(&ft)
	results := shift.LetterFrequency(p, &ft, false)
	assert.Contains(t, results, p)
}

func TestLetterFrequency_1(t *testing.T) {
	ft := FrequencyTable{{'b', 0.1}, {'c', 0.2},
		{'d', 0.3}, {'a', 0.4}}
	p := GeneratePlaintext(&ft)
	c := algo.Encode(p, 1)
	results := shift.LetterFrequency(c, &ft, false)
	assert.Contains(t, results, p)
}

func TestLetterFrequency_2(t *testing.T) {
	ft := FrequencyTable{{'b', 0.1}, {'c', 0.2},
		{'d', 0.3}, {'a', 0.4}}
	p := GeneratePlaintext(&ft)
	c := algo.Encode(p, 22)
	results := shift.LetterFrequency(c, &ft, false)
	assert.Contains(t, results, p)
}

func TestLetterFrequency_8(t *testing.T) {
	test8_fq := FrequencyTable{  {'a', 0.082}, {'b', 0.015}, {'c', 0.028}, {'d', 0.043},
									{'e', 0.127}, {'f', 0.022}, {'g', 0.020}, {'h', 0.061},
									{'i', 0.602}}

	p := GeneratePlaintext(&test8_fq)
	//c := algo.Encode(p, 22)
	results := shift.LetterFrequency(p, &test8_fq, false)
	assert.Contains(t, results, p)
}

func TestLetterFrequency_16(t *testing.T) {
	test16_fq := FrequencyTable{  {'a', 0.082}, {'b', 0.015}, {'c', 0.028}, {'d', 0.043},
								{'e', 0.127}, {'f', 0.022}, {'g', 0.020}, {'h', 0.061},
								{'i', 0.070}, {'j', 0.002}, {'z', 0.530}}

	p := GeneratePlaintext(&test16_fq)
	//c := algo.Encode(p, 22)
	results := shift.LetterFrequency(p, &test16_fq, false)
	assert.Contains(t, results, p)
}
