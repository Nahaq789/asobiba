package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	split := strings.Split(s, "S")
	cnt := 0
	for _, v := range split {
		if len(v) > cnt {
			cnt = len(v)
		}
	}

	fmt.Println(cnt)
}
