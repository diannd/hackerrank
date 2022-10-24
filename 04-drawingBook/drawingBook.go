package main

import "fmt"

func main() {
	var n, p int32
	n = 6
	p = 2

	last := (n - p) / 2
	first := p / 2

	if n%2 == 0 {
		last = (n + 1 - p) / 2
	}

	if last < first {
		fmt.Println(last)
	} else {
		fmt.Println(first)
	}
}
