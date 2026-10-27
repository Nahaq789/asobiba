package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range n {
		var aaa int
		fmt.Scan(&aaa)
		a[i] = aaa
	}

	var k int
	fmt.Scan(&k)

	cnt := 0
	for _, v := range a {
		if v >= k {
			cnt++
		}
	}

	fmt.Println(cnt)
}
