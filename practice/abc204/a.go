package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x, y int
	fmt.Fscan(in, &x, &y)

	if x == y {
		fmt.Fprintln(out, x)
		return
	}

	a := [3]bool{false, false, false}
	a[x] = true
	a[y] = true

	for i := 0; i < 3; i++ {
		if !a[i] {
			fmt.Fprintln(out, i)
			return
		}
	}
}
