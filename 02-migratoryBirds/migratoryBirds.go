package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr []int32
	maps := map[int32]int32{}
	double := make([]int32, 0, len(maps))

	for _, v := range arr {
		maps[v]++
	}

	for key, v := range maps {
		if v > 1 {
			double = append(double, key)

		}
	}

	sort.SliceStable(double, func(i, j int) bool {
		return maps[double[i]] > maps[double[j]]
	})

	fmt.Println(double[0])
}
