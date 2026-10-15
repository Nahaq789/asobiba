package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	digitSum(n, k)
}

func digitSum(n, k int) {
	cnt := 0
	for i := range n + 1 {
		var r rune
		for _, c := range strconv.Itoa(i) {
			r += (c - '0')
		}
		if int(r) == k {
			cnt += 1
		}
	}

	fmt.Println(cnt)
}
