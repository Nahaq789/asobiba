package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}

	unko := a[:3]
	k := 3
	for {
		slices.Sort(unko)
		slices.Reverse(unko)
		fmt.Fprintln(out, unko[2])

		if k == n {
			break
		}

		ak := a[k]
		unkoSuffix := unko[2]
		if unkoSuffix < ak {
			unko[2] = ak
		}

		k++
	}
}
