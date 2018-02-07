package enigma


type enigma struct {
	rotors 	[]*Rotor
	key 	[]rune
}


func NewEnigma(rs []*Rotor, k []rune) *enigma {
	if nil == rs {
		r1 := Rotor{Alphabet: []rune("abcdefghijklmnopqrstuvwxyz")}
		r2 := Rotor{Alphabet: []rune("abcdefghijklmnopqrstuvwxyz")}
		r3 := Rotor{Alphabet: []rune("abcdefghijklmnopqrstuvwxyz")}
		rs = []*Rotor{&r1, &r2, &r3}
	}
	m := enigma{rotors: rs, key: k}
	for i, r := range m.rotors {
		r.SetOffset(m.key[i])
	}
	return &m
}


func (m *enigma) Encode(ptext string) string{
	result := make([]rune, len(ptext))
	rotorsEncode(ptext, m.rotors, result)
	return string(result)
}

func rotorsEncode(ptext string, rs []*Rotor, result []rune) {
	for i, c := range ptext {
		for _, r := range rs {
			c = r.EncodeLetter(c)
		}
		result[i] = c
	}

}


func (m *enigma) Decode(ctext string) string{
	result := make([]rune, len(ctext))
	rotorsDecode(stringReverse(ctext), rotorReverse(m.rotors), result)
	return stringReverse(string(result))
}


func rotorsDecode(ctext string, rs []*Rotor, result []rune) {
	for i, c := range ctext {
		for _, r := range rs {
			c = r.DecodeLetter(c)
		}
		result[i] = c
	}
}


func rotorReverse(a []*Rotor) []*Rotor {
	result := make([]*Rotor, len(a))
	for i, v := range a {
		result[len(a) - i - 1] = v
	}
	return result
}

func stringReverse(a string) string {
	result := make([]rune, len(a))
	for i, v := range a {
		result[len(a) - i - 1] = v
	}
	return string(result)
}
