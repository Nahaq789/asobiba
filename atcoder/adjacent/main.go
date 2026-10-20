package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]string, 2)
	for i := range 2 {
		fmt.Scan(&s[i])
	}
	res := []string{}
	for _, str := range s {
		res = append(res, strings.Split(str, "")...)
	}
	adjacent(res)
}

func adjacent(s []string) {
	hoge := make(map[int][]string, 2)
	for i, v := range s {
		if i == 0 || i == 1 {
			hoge[0] = append(hoge[0], v)
		}
		if i == 2 || i == 3 {
			hoge[1] = append(hoge[1], v)
		}
	}

	if hoge[0][0] == "A" {
		if hoge[0][1] == "D" {
			fmt.Println("Yes")
			return
		}
	}
	if hoge[0][0] == "D" {
		if hoge[0][1] == "A" {
			fmt.Println("Yes")
			return
		}
	}

	if hoge[1][0] == "A" {
		if hoge[1][1] == "D" {
			fmt.Println("Yes")
			return
		}
	}
	if hoge[1][0] == "D" {
		if hoge[1][1] == "A" {
			fmt.Println("Yes")
			return
		}
	}

	index := []int{}
	for _, v := range hoge {
		for j, str := range v {
			if str == "A" {
				index = append(index, j)
			}
			if str == "D" {
				index = append(index, j)
			}
		}

	}
	if index[0] == index[1] {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
