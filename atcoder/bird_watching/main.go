package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	birdMap := map[int][]int{}
	for i := 0; i < n; i++ {
		kind := a[i]
		value := b[i]
		birdMap[kind] = append(birdMap[kind], value)
	}

	for _, k := range slices.Sorted(maps.Keys(birdMap)) {
		sum := 0
		for _, v := range birdMap[k] {
			sum += v
		}
		fmt.Printf("%.20f\n", float64(sum)/float64(len(birdMap[k])))
	}
}
