package main

import "fmt"

func main() {
	var x, y, z int32
	x = 1
	y = 4
	z = 3

	a := z - x
	b := z - y

	if b < 0 {
		b *= -1
	}

	if a < 0 {
		a *= -1
	}

	if a == b {
		fmt.Println("Mouse C")
	} else if a < b {
		fmt.Println("Cat A")
	} else {
		fmt.Println("Cat B")
	}
}
