package main

import "fmt"

func main() {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	res := map[int32]bool{}
	count := int32(0)

	for _, v := range ar {
		if res[v] {
			res[v] = false
			count++
		} else {
			res[v] = true
		}
	}

	fmt.Println(count)
}
