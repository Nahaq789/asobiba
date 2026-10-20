package main

import (
	"fmt"
)

func main() {
	var s [2]string
	fmt.Scan(&s[0], &s[1])

	var ar, ac, dr, dc int
	for i, str := range s {
		for j, v := range str {
			if v == 'A' {
				ar = i
				ac = j
			}
			if v == 'D' {
				dr = i
				dc = j
			}
		}
	}

	if ar == dr || ac == dc {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
