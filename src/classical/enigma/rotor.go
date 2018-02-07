package enigma

import (
	"../../utilities"
)


type Rotor struct {
	Alphabet []rune
	Offset   int
}


func NewRotor(a []rune, k rune) *Rotor {
	return &Rotor{Alphabet: a, Offset: getOffset(k)}
}


func (r *Rotor) EncodeLetter(l rune) rune {
	index := getEncIndex(l, r)
	result := r.Alphabet[index]
	r.Shift()
	return result
}

func (r *Rotor) DecodeLetter(l rune) rune {
	r.UnShift()
	index := getDecIndex(l, r)
	result := r.Alphabet[index]
	return result
}



func (r *Rotor) SetOffset(k rune) {
	r.Offset = getOffset(k)
}

func (r *Rotor) Shift() {
	r.Offset = (r.Offset + 1) %  len(r.Alphabet)
}

func (r *Rotor) UnShift() {
	r.Offset = ((r.Offset - 1) + len(r.Alphabet)) %  len(r.Alphabet)
}

func getEncIndex(l rune, r *Rotor) int {
	return (getOffset(l) + r.Offset) % len(r.Alphabet)
}

func getDecIndex(l rune, r *Rotor) int {
	return ((getOffset(l) - r.Offset) + len(r.Alphabet)) % len(r.Alphabet)
}

func getOffset(l rune) int {
	return utilities.GetAlphabetOffset(l)
}

func (r *Rotor)GetKey() rune{
	return r.Alphabet[int(r.Offset)]
}