package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)

	for i := 0; int64(i) < n; i++ {
		if int64(2<<i) <= n {
			continue
		} else {
			fmt.Println(i)
			break
		}
	}
}
