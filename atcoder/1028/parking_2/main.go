package main

import "fmt"

func main() {
	var x, y, l, r, a, b int
	fmt.Scan(&x, &y, &l, &r, &a, &b)

	result := 0

	for i := a; i < (b); i++ {
		if i >= l && r > i {
			result += x
		} else {
			result += y
		}
	}

	fmt.Println(result)
}
