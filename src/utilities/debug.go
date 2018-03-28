package utilities

import (
	"time"
	"fmt"
	"encoding/json"
)


type Debug struct {
	PrintInfo bool
}


func (d *Debug) Info(s string) {
	if d.PrintInfo {
		fmt.Println(s)
	}
}


func (d *Debug) Start(s string) *time.Time {
	t := time.Now()
	if d.PrintInfo {
		fmt.Println(s)
	}
	return &t
}



func (d *Debug) Finish(t *time.Time) {
	if d.PrintInfo {
		fmt.Println(" duration:", time.Since(*t))
	}
}



func (d *Debug) PrintPretty(v interface{}) {
	if d.PrintInfo {
		a, e := json.Marshal(v)
		if e != nil {
			fmt.Println(e)
			return
		}
		fmt.Println(string(a))
	}
}

