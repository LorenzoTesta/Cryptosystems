package utilities_test

import (
	"testing"
	. "../utilities"
	"github.com/stretchr/testify/assert"
)



func TestBuildStrings_0(t *testing.T) {
	ct1 := []ChangeTable{	{'a': 'a', 'b': 'a', 'c': 'a'},
							{'a': 'b', 'b': 'b', 'c': 'b'},
							{'a': 'c', 'b': 'c', 'c': 'c'},
							{'a': 'z', 'b': 'k', 'c': 'y'}}
	expected := []string{"aaaaaaaaa", "bbbbbbbbb", "ccccccccc", "zzzkkkyyy"}
	result :=  BuildStrings(&ct1, "aaabbbccc", 9)
	assert.Equal(t, expected, result)
}



func TestAppend_0(t *testing.T) {
	ct1 := []ChangeTable{{'a': 'b'}}
	ct2 := []ChangeTable{{'b': 'c', 'c': 'd'}}
	expected := []ChangeTable{{'a': 'b'}, {'b': 'c', 'c': 'd'}}
	result :=  Append(&ct1, &ct2)
	assert.Equal(t, expected, *result)
}


func TestAppendToAll_0(t *testing.T) {
	ct1 := []ChangeTable{{'a': 'b'}, {'b': 'c', 'c': 'd'}}
	AppendToAll(&ct1, 'z', 'k')
	expected := []ChangeTable{{'a': 'b', 'z': 'k'}, {'b': 'c', 'c': 'd', 'z': 'k'}}
	assert.Equal(t, expected, ct1)
}


