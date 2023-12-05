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

	var a, b float64
	fmt.Fscan(in, &a, &b)

	var x = b / 100

	fmt.Fprintln(out, a*x)
}
