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

	var n, v int
	fmt.Fscan(in, &n, &v)
	fmt.Fprintln(out, n+v)

	var s string
	fmt.Fscan(in, &s)
	fmt.Fprintln(out, s)
}
