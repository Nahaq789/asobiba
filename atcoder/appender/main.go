package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	suffix := "er"
	if s[len(s)-1] == 'e' {
		suffix = "r"
	}
	fmt.Println(s + suffix)
}
