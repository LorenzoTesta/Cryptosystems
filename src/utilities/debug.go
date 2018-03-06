package utilities

import (
	"time"
	"fmt"
	"encoding/json"
)



func Start(s string) *time.Time {
	t := time.Now()
	fmt.Println(s)
	return &t
}



func Finish(t *time.Time) {
	fmt.Println(" duration:", time.Since(*t))
}



func PrintPretty(v interface{}){
	a, e := json.Marshal(v)
	if e != nil {
		return
	}
	fmt.Println(string(a))
}

