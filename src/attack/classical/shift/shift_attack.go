package shift

import (
	method "../../../classical/shift"
	. "../../../utilities"
	"sort"
	"time"
	"fmt"
)



// Ciphertext only
func Exhaustive(c string) [] string {
	result := make([]string, 26)
	for i := 0; i < 26; i++ {
		result[i] = method.Decode(c, i)
	}
	return result
}



	var FREQTABLE = FrequencyTable{  {'a', 0.082}, {'b', 0.015}, {'c', 0.028}, {'d', 0.043},
									 {'e', 0.127}, {'f', 0.022}, {'g', 0.020}, {'h', 0.061},
									 {'i', 0.070}, {'j', 0.002}, {'k', 0.008}, {'l', 0.040},
									 {'m', 0.024}, {'n', 0.067}, {'o', 0.075}, {'p', 0.019},
									 {'q', 0.001}, {'r', 0.060}, {'s', 0.063}, {'t', 0.091},
									 {'u', 0.028}, {'v', 0.010}, {'w', 0.023}, {'x', 0.001},
									 {'y', 0.020}, {'z', 0.001} }



func LetterFrequency(c string) [] string {

	startAll := time.Now()

	t := Start("- count the occurrences... ")
	countTab := NewFrequencyTable(c) // CONTO LE OCCORRENZE
	Finish(t)

	sort.Sort(countTab)

	PrintPretty(countTab)

	t = Start("- calculate rating... ")
	countTab.CalcPerc(len(c)) // CALCOLO DELLE PERCENTUALI
	Finish(t)

	PrintPretty(countTab)

	t = Start("- create matching tables... ")
	pairs := MatchesMap{}
	for _, l := range countTab { // ASSOCIAZIONE c -> possibili valori
		deltasTab := FREQTABLE.CalcDeltas(l.Rate)
		sort.Sort(deltasTab)
		fmt.Print(string(l.Key), " ", l.Key, l.Rate, " -> ")
		pairs[l.Key] = deltasTab.GetSmallerThan(0.0002)
		pairs[l.Key] = deltasTab.GetBestNmatch(1)
	}
	Finish(t)

	PrintPretty(pairs)

	t = Start("- building association tables... ")
	tables := BuildTabs(&pairs)
	Finish(t)

	PrintPretty(tables)

	t = Start("- building strings... ")
	results := BuildStrings(tables, c, 100)
	Finish(t)

	PrintPretty(results)

	fmt.Println(" TOTAL TIME:", time.Since(startAll))
	return results
}



func BuildTabs(matches *MatchesMap) []*ChangeTable {
	results := []*ChangeTable{}

	// prendo un array di corrispondenze
	for l, v := range *matches {
		for _, k := range v {
			// clono la sottoMappa di corrispondenze
			clone := Clone(matches)

			// rimuovo la lettera corrente
			delete(*clone, l)

			// rimuovo la corrispondenza scelta
			clone.clearTab(k)

			// calcolo i sottoalberi
			subTabs := BuildTabs(clone)

			// aggiungo ad ogni sottoalbero la radice comune
			appendToAll(subTabs, l, k)
		}
	}
	return results
}

func appendToAll(tables []*ChangeTable, l rune, s rune) {
	for _, v := range tables {
		(*v)[l] = s
	}
}


func (mm *MatchesMap)clearTab(l rune) {
	for _, a := range *mm {
		i := contains(a, l)
		if i > 0 {
			a = append(a[:i], a[i+1:]...)
		}
	}
}


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

type ChangeTable map[rune]rune

func contains(s []rune, e rune) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
