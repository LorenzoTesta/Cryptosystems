package utilities


func GetLetter (r rune, shift int) rune {

	s := int(r) + (shift % 26)

	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}


func GetAlphabetOffset(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - 'a'
	} else {
		return 0
	}
}