package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var q int
	fmt.Fscan(in, &q)
	var s, t string
	fmt.Fscan(in, &s, &t)

	query := [][]int{}
	for range q {
		var l, r int
		fmt.Fscan(in, &l, &r)
		query = append(query, []int{l, r})
	}

	for _, v := range query {
		l, r := v[0], v[1]
		aaa := s[l-1 : r]

		if strings.Contains(aaa, t) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
