package utilities

import (
	"math"
	"sort"
)

type Value struct {
	Key  rune
	Rate float64
}


type FrequencyTable []*Value


func (p FrequencyTable) Len() int           { return len(p) }
func (p FrequencyTable) Less(i, j int) bool { return p[i].Key < p[j].Key || p[i].Rate < p[j].Rate}
func (p FrequencyTable) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func NewFrequencyTable(s string) FrequencyTable {
	count := map[rune]float64{}
	result := FrequencyTable{}

	for _, v := range s {
		count[v]++
	}

	for i, v := range count {
		result = append(result, &Value{i, v})
	}

	return result
}

func (ft *FrequencyTable) CalcPerc(l int) {
	k := 1 / float64(l)
	for _, v := range *ft {
		v.Rate = v.Rate * k
	}
}

func (ft *FrequencyTable) GetSmallerThan(x float64) []rune {
	result := make([]rune, 0, len(*ft))
	for _, v := range *ft {
		if v.Rate <= x {
			result = append(result, v.Key)
		} else {
			break
		}
	}
	return result
}


func (ft *FrequencyTable) CalcDeltas(count float64) FrequencyTable {
	deltas := make(FrequencyTable, len(*ft))
	for i, v := range *ft {
		deltas[i] = &Value{ v.Key, math.Abs(v.Rate - count)}
	}
	return deltas
}


func (a FrequencyTable) Equal(b FrequencyTable) bool {
	if len(a) != len(b) { return false }
	sort.Sort(a)
	sort.Sort(b)

	for i := 0; i < len(a); i++ {

		if !a[i].Equal(b[i]) { return false }
	}
	return true
}


func (a *Value) Equal(b *Value) bool {
	return a.Key == b.Key && a.Rate == b.Rate
}


































