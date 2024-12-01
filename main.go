package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println("part 1:", sum)
}
