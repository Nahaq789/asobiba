package main

import "fmt"

func main() {
	fmt.Println(NotX(3))
}

func NotX(x int) int {
	list := []int{1, 2, 3}
	var result int
	for _, v := range list {
		if v == x {
			continue
		}
		result = v
	}
	return result
}
