package main

import "fmt"

func main() {
	var sum, k, b int32
	bill := []int32{3, 10, 2, 9}
	k = bill[1]
	b = 7

	for _, v := range bill {
		sum += v
	}

	tmp := (sum - k) / 2
	// fmt.Println(tmp)
	total := b - tmp

	if b == tmp {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(total)
	}
}
