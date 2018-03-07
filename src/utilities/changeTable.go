package utilities


type ChangeTable map[rune]rune


func BuildStrings(tables []*ChangeTable, ciphertext string, nchar int) []string {
	result := []string{}
	for i, tab := range tables {
		plaintext := []rune{}
		c := []rune(ciphertext)
		for j:=0; j<nchar; j++ {
			plaintext[j] = (*tab)[ c[j] ]
		}
		result[i] = string(plaintext)
	}
	return result
}


func AppendToAll(tables []*ChangeTable, l rune, s rune) {
	for _, v := range tables {
		(*v)[l] = s
	}
}
