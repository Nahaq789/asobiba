package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	var xs []int
	for range n {
		var x int
		fmt.Scan(&x)
		xs = append(xs, x)
	}

	result := []int{}
	for i, v := range xs {
		ok := true
		for j, vv := range xs {
			if i != j && abs(v-vv) < d {
				ok = false
				break
			}
		}
		if ok {
			result = append(result, i+1)
		}
	}
	fmt.Println(len(result))
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
