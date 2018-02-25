package utilities

import (
	"math"
)

type Value struct {
	Key rune
	Value float64
}


type FrequencyTable []Value


func (p FrequencyTable) Len() int           { return len(p) }
func (p FrequencyTable) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p FrequencyTable) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func NewFrequencyTable(s string) FrequencyTable {
	count := map[rune]float64{}
	result := FrequencyTable{}

	for _, v := range s {
		count[v]++
	}

	for i, v := range count {
		result = append(result, Value{i, v})
	}

	return result
}

func (ft *FrequencyTable) CalcPerc(l int) {
	k := float64(l / 100)
	for _, v := range *ft {
		v.Value = v.Value * k
	}
}

func (ft *FrequencyTable) GetSmallerThan(x float64) []rune {
	result := make([]rune, 0, len(*ft))
	for _, v := range *ft {
		if v.Value <= x {
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
		deltas[i] = Value{ v.Key, math.Abs(v.Value - count)}
	}
	return deltas
}



