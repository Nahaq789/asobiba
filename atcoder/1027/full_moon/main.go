package main

import "fmt"

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)

	if m > n {
		fmt.Println(0)
		return
	}

	cnt := 1
	i := 1 // today
	for range n {
		now := m + (i * p)
		if now <= n {
			cnt++
		}
		i++
	}

	fmt.Println(cnt)
}
