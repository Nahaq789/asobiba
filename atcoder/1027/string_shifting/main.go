package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	mn, mx := s, s
	for i := 0; i < len(s); i++ {
		t := s[i:] + s[:i]
		// fmt.Println("f", s[i:], "e", s[:i], "t", t)
		if t < mn {
			mn = t
		}

		if t > mx {
			mx = t
		}
	}

	fmt.Println(mn)
	fmt.Println(mx)
}
