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

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)

	if a == b {
		fmt.Fprintln(out, c)
		return
	}
	if a == c {
		fmt.Fprintln(out, b)
		return
	}
	if b == c {
		fmt.Fprintln(out, a)
		return
	}
	fmt.Fprintln(out, 0)
}
