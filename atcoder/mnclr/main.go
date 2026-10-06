package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(mnclr(s))
}

func mnclr(s string) string {
	var result strings.Builder
	for _, v := range s {
		result.WriteString(string(v) + "o")
	}
	r := result.String()
	return r[:len(r)-1]
}
