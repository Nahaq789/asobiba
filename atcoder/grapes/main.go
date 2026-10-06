package main

import "fmt"

func main() {
	p := grapes(12, 3456)
	fmt.Println(p)
}

func grapes(n, m int) []int {
	p := make([]int, n)
	index := 0
	g := m
	for {
		now := p[index]
		p[index] = now + 1
		index += 1
		if index == (n) {
			index = 0
		}

		g -= 1
		if g == 0 {
			break
		}
	}
	return p
}
