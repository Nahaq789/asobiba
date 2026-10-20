package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	s := make([]string, n)
	for i := range n {
		fmt.Fscan(in, &s[i])
	}

	t := make([]string, m)
	for i := range m {
		fmt.Fscan(in, &t[i])
	}

	invalidVote(s, t)
}

func invalidVote(s, t []string) {
	list := map[string]int{}
	for _, name := range s {
		list[name] += 0
	}
	for _, tname := range t {
		for _, sname := range s {
			if tname == sname {
				list[sname] += 1
			}
			if diff(tname, sname) {
				list[sname] += 1
			}
		}
	}

	for _, k := range s {
		fmt.Println(list[k])
	}
}

func diff(tname, sname string) bool {
	if len(tname) != len(sname) {
		return false
	}
	cnt := 0
	for i := 0; i < len(tname); i++ {
		if tname[i] != sname[i] {
			cnt += 1
		}
	}
	return cnt == 1
}
