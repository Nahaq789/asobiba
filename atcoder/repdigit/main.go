package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%111 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
