package main

import (
	"fmt"
	"strings"
)

func main() {
	var c string
	fmt.Scan(&c)

	order := "BYR"
	i := strings.Index(order, c)
	fmt.Println(string(order[(i+1)%len(order)]))
}
