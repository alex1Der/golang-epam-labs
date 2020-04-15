package main

import (
	"fmt"
	"sort"
)

func main() {
	var array = []int{4, 1, 3, 8}
	fmt.Printf("%v", median(array))
}

func median(i []int) float64 {
	arrayLength := len(i)
	sort.Ints(i)
	if arrayLength % 2 != 0 {
		return float64(i[arrayLength / 2])
	}

	return float64(i[(arrayLength - 1) / 2] +
		i[arrayLength / 2]) / 2
}