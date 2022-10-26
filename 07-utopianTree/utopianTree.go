package main

import "fmt"

func main() {
	var H int32 = 1
	var n int32 = 5

	for i := 1; int32(i) < n+1; i++ {
		if i%2 != 0 {
			H = 2 * H
		} else {
			H += 1
		}
	}
	fmt.Println(H)
}
