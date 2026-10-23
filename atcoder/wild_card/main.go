package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var s, t string
	fmt.Scan(&s, &t)

	ok := true
	for i := range n {
		if t[i] == '*' {
			continue
		}

		if s[i] != t[i] {
			ok = false
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
