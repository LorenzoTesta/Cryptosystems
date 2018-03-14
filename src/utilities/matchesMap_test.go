package utilities_test

import (
	"testing"
	. "../utilities"
	"github.com/stretchr/testify/assert"
)


func TestMatchesMap_ClearTab_0(t *testing.T) {
	mm := MatchesMap{'a':[]rune{'b'}, 'b':[]rune{'b','c'}, 'c':[]rune{'b','a'}}
	expected := MatchesMap{'a':[]rune{}, 'b':[]rune{'c'}, 'c':[]rune{'a'}}
	mm.ClearTab('b')
	assert.Equal(t, expected, mm)
}


func TestMatchesMap_Clone_0(t *testing.T) {
	mm 		 := MatchesMap{'a':[]rune{'b'}, 'b':[]rune{'b','c'}, 'c':[]rune{'b','a'}}
	expected := MatchesMap{'a':[]rune{'b'}, 'b':[]rune{'b','c'}, 'c':[]rune{'b','a'}}
	assert.Equal(t, expected, mm)
}
