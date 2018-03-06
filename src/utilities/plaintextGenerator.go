package utilities

import "strings"

func GeneratePlaintext(ft *FrequencyTable) string {
	result := ""
	for _, v := range *ft {
		result += strings.Repeat(string(v.Key), int(v.Rate * 1000))
	}
	return result
}