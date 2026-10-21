package main

import "fmt"

func main() {
	var w, b int
	fmt.Scan(&w, &b)

	weight := w * 1000

	result := weight / b

	fmt.Println(result + 1)
}
