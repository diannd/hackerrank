package main

import "fmt"

func main() {
	a := []int32{-1, 0, 4, 2}
	k := 2

	var count int

	for _, v := range a {
		if v <= 0 {
			count++
		}
	}

	if count <= k {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
