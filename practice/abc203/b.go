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

	var n, k int
	fmt.Fscan(in, &n, &k)

	total := 0
	for i := 1; i <= n; i++ {
		a := 100 * i
		for j := 1; j <= k; j++ {
			total += a + j
		}
	}

	fmt.Fprintln(out, total)
}
