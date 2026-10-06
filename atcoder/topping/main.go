package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topping(10, 16, 102228, 448944, 131224, 326172, 500169, 670309, 976672, 579051, 974511, 773940))
	fmt.Println(topping(5, 9, 31, 41, 59, 26, 53))
}

type Result struct {
	list []int
	sum  int
	sum2 int
}

func topping(n, v int, w ...int) int {
	results := make([]Result, 0)
	for i := range n {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				a := []int{i + 1, j + 1, k + 1}
				sum := sumf(a)
				sum2 := sumf([]int{w[i], w[j], w[k]})
				if sum <= v {
					results = append(results, Result{
						list: a,
						sum:  sum,
						sum2: sum2,
					})
				}
			}
		}
	}

	a := make([]Result, 0)
	for _, vv := range results {
		if vv.sum <= v {
			a = append(a, vv)
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].sum2 > a[j].sum2
	})

	maxV := a[0].sum2
	return maxV
}

func sumf(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
