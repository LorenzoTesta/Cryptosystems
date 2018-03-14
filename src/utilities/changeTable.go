package utilities


type ChangeTable map[rune]rune


func BuildStrings(tables *[]ChangeTable, ciphertext string, nchar int) []string {
	result := []string{}
	for _, tab := range *tables {			// scorro le tabelle
		plaintext := []rune{}
		//c := []rune(ciphertext)
		for _, c := range ciphertext {
			plaintext = append(plaintext, tab[c])
		}
		//for j:=0; j<nchar; j++ {
		//	plaintext[j] = tab[ c[j] ]
		//}
		result = append(result, string(plaintext))
	}
	return result
}


func Append(a *[]ChangeTable, b *[]ChangeTable) *[]ChangeTable {
	result := []ChangeTable{}
	for _, v := range *a {
		result = append(*a, v)
	}
	for _, v := range *b {
		result = append(*a, v)
	}
	return &result
}


func AppendToAll(tables *[]ChangeTable, l rune, s rune) {
	if len(*tables) == 0 {
		*tables = append(*tables, ChangeTable{l: s})
	}
	for _, v := range *tables {
		v[l] = s
	}
}
