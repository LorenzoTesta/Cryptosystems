package utilities



func (mm *MatchesMap)ClearTab(l rune) {
	for _, a := range *mm {
		i := contains(a, l)
		if i > 0 {
			a = append(a[:i], a[i+1:]...)
		}
	}
}

type Match struct {
	Values rune
}

type MatchesMap map[rune][]rune

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
