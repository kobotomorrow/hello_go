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

	var n int
	fmt.Fscan(in, &n)

	total := 0
	count := 1
	for {
		total += count
		if total >= n {
			break
		}
		count++
	}
	fmt.Fprintln(out, count)
}
