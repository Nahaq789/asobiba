package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	r := change(a)
	fmt.Println(strings.Trim(fmt.Sprint(r), "[]"))
}

type Hoge struct {
	list []int
}

func change(a []int) []int {
	final := make([]int, 3)
	for _, v := range a {
		if v%1000 == 0 {
			continue
		}
		oturi := 1000 - (v % 1000)

		c100 := oturi / 100
		oturi -= c100 * 100

		c10 := oturi / 10
		oturi -= c10 * 10

		c1 := oturi / 1
		oturi -= c1

		final[0] += c1
		final[1] += c10
		final[2] += c100
	}
	return final
}

//
// func change(a []int) []int {
// 	var nokori int
// 	hoges := make([]Hoge, len(a))
// 	for i, v := range a {
// 		senyenCount := 0
// 		c100 := 0
// 		c10 := 0
// 		c1 := 0
// 		for {
// 			senyenCount += 1
// 			if (v - 1000) > 0 {
// 				v -= 1000
// 			} else {
// 				break
// 			}
// 		}
// 		nokori = (1000 * senyenCount) - a[i]
//
// 		for {
// 			if nokori/100 > 0 {
// 				c100 += 1
// 				nokori -= 100
// 			}
// 			if nokori/100 == 0 {
// 				break
// 			}
// 		}
// 		for {
// 			if nokori/10 > 0 {
// 				c10 += 1
// 				nokori -= 10
// 			}
// 			if nokori/10 == 0 {
// 				break
// 			}
// 		}
// 		for {
// 			if nokori/1 > 0 {
// 				c1 += 1
// 				nokori -= 1
// 			}
// 			if nokori/1 == 0 {
// 				break
// 			}
// 		}
// 		hoges[i] = Hoge{
// 			list: []int{c1, c10, c100},
// 		}
// 		nokori = 0
// 	}
//
// 	result := make([]int, 3)
// 	for _, v := range hoges {
// 		result[0] += v.list[0]
// 		result[1] += v.list[1]
// 		result[2] += v.list[2]
// 	}
//
// 	return result
// }
