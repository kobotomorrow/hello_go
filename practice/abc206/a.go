package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n float64
	fmt.Fscan(in, &n)

	x := math.Floor(n * 1.08)

	if x < 206 {
		fmt.Fprintln(out, "Yay!")
	} else if x == 206 {
		fmt.Fprintln(out, "so-so")
	} else {
		fmt.Fprintln(out, ":(")
	}
}
