package utilities


type MatchesMap map[rune][]rune


func (mm *MatchesMap)ClearTab(l rune) {
	for i, a := range *mm {
		j := contains(a, l)
		if j > -1 {
			a = append(a[:j], a[j+1:]...)
		}
		(*mm)[i] = a
	}
}


func Clone(matches *MatchesMap) *MatchesMap {
	result := MatchesMap{}
	for i, v := range *matches {
		result[i] = v
	}
	return &result
}


func contains(s []rune, e rune) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
