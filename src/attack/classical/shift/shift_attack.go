package shift

import (
	method "../../../classical/shift"
	. "../../../utilities"
	"sort"
	"time"
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



func LetterFrequency(c string, ft *FrequencyTable, debug bool) [] string {

	startAll := time.Now()
	d := Debug{debug}

	t := d.Start("- count the occurrences... ")
	countTab := NewFrequencyTable(c)
	d.Finish(t)

	sort.Sort(countTab)

	t = d.Start("- calculate rating... ")
	countTab.CalcPerc(len(c))
	d.Finish(t)

	t = d.Start("- create matching tables... ")
	pairs := MatchesMap{}
	for _, l := range *countTab {
		deltasTab := ft.CalcDeltas(l.Rate)
		sort.Sort(deltasTab)
		//pairs[l.Key] = deltasTab.GetSmallerThan(0.0002)
		pairs[l.Key] = deltasTab.GetBestNmatch(1)
	}
	d.Finish(t)

	d.PrintPretty(pairs)

	t = d.Start("- building association tables... ")
	tables := BuildTabs(&pairs)
	//tables := BuildTabs(&pairs)
	d.Finish(t)

	d.PrintPretty(tables)

	t = d.Start("- building strings... ")
	results := BuildStrings(tables, c, 100)
	d.Finish(t)

	d.Info(" TOTAL TIME: " + string(time.Since(startAll)))
	return results
}


func BuildTabs(matches *MatchesMap) *[]ChangeTable {
	results := &[]ChangeTable{}
	l := matches.GetChar()
	if l != 0 {
		v := (*matches)[l]
		if len(v) != 0 {
			k := v[0]
			clone := Clone(matches)
			delete(*clone, l)
			clone.ClearTab(k)
			subTabs := BuildTabs(clone)
			AppendToAll(subTabs, l, k)
			results = Append(results, subTabs)
		}
	}
	return results
}




