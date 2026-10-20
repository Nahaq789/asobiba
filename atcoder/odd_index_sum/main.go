package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range n {
		fmt.Scan(&a[i])
	}

	cnt := 0
	for i, v := range a {
		if (i+1)%2 != 0 {
			cnt += v
		}
	}

	fmt.Println(cnt)
}
