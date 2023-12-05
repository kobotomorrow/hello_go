package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var a []int
	for i := 0; i < n; i++ {
		var aa int
		fmt.Fscan(in, &aa)
		a = append(a, aa)
	}

	sort.SliceStable(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for i, aa := range a {
		i = i + 1
		if i != aa {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
}
