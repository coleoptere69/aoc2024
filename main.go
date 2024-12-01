package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		leftNum, rightNum := 0, 0
		_, err = fmt.Sscanf(scanner.Text(), "%d   %d", &leftNum, &rightNum)
		if err != nil {
			panic(err)
		}
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println("part 1:", sum)
}
