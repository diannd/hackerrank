package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int32

	x1 = 0
	v1 = 3
	x2 = 4
	v2 = 2

	if v1 < v2 {
		fmt.Println("NO")
	}

	if (x2-x1)%(v1-v2) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
