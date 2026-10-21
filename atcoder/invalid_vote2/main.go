package main

import (
	"bufio"
	"fmt"
	"os"
)

func diffOne(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	s := make([]string, n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	votes := make([]int, n)

	for j := 0; j < m; j++ {
		var t string
		fmt.Fscan(in, &t)

		exact := -1
		cand := []int{}
		for i := 0; i < n; i++ {
			if s[i] == t {
				exact = i
				break
			}
			if diffOne(s[i], t) {
				cand = append(cand, i)
			}
		}

		if exact >= 0 {
			votes[exact]++
		} else if len(cand) == 1 {
			votes[cand[0]]++
		}
	}

	for _, v := range votes {
		fmt.Println(v)
	}
}
