package shift

import (
	method "../../../classical/shift"
	. "../../../utilities"
	"sort"
	"time"
	"fmt"
)



func Exhaustive(c string) [] string {
	result := make([]string, 26)
	for i := 0; i < 26; i++ {
		result[i] = method.Decode(c, i)
	}
	return result
}




var ENG_FREQTABLE = FrequencyTable{  {'a', 0.082}, {'b', 0.015}, {'c', 0.028}, {'d', 0.043},
									 {'e', 0.127}, {'f', 0.022}, {'g', 0.020}, {'h', 0.061},
									 {'i', 0.070}, {'j', 0.002}, {'k', 0.008}, {'l', 0.040},
									 {'m', 0.024}, {'n', 0.067}, {'o', 0.075}, {'p', 0.019},
									 {'q', 0.001}, {'r', 0.060}, {'s', 0.063}, {'t', 0.091},
									 {'u', 0.028}, {'v', 0.010}, {'w', 0.023}, {'x', 0.001},
									 {'y', 0.020}, {'z', 0.001} }



func LetterFrequency(c string, ft *FrequencyTable) [] string {

	startAll := time.Now()

	t := Start("- count the occurrences... ")
	countTab := NewFrequencyTable(c)
	Finish(t)

	sort.Sort(countTab)
	PrintPretty(countTab)

	t = Start("- calculate rating... ")
	countTab.CalcPerc(len(c))
	PrintPretty(countTab)
	Finish(t)

	t = Start("- create matching tables... ")
	pairs := MatchesMap{}
	for _, l := range *countTab {
		deltasTab := ft.CalcDeltas(l.Rate)
		sort.Sort(deltasTab)
		fmt.Print(string(l.Key), " ", l.Key, l.Rate, " -> deltas: ")
		//pairs[l.Key] = deltasTab.GetSmallerThan(0.0002)
		pairs[l.Key] = deltasTab.GetBestNmatch(1)
	}
	PrintPretty(pairs)
	Finish(t)

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



func BuildTabs(matches *MatchesMap) *[]ChangeTable {
	results := &[]ChangeTable{}

	//PrintPretty(matches)

	// prendo un array di corrispondenze
	for l, v := range *matches {
		for _, k := range v {
			// clono la sottoMappa di corrispondenze
			clone := Clone(matches)

			// rimuovo la lettera corrente
			delete(*clone, l)

			// rimuovo la corrispondenza scelta
			clone.ClearTab(k)

			// calcolo i sottoalberi
			subTabs := BuildTabs(clone)

			// aggiungo ad ogni sottoalbero la radice comune
			AppendToAll(subTabs, l, k)

			results = Append(results, subTabs)
		}
	}
	fmt.Print("partial result: ")
	PrintPretty(results)
	return results
}




