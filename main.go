package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	part1()
	part2()
}

func part1() {
	left, right := ParseInput()
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println("part 1:", sum)
}

func part2() {
	left, right := ParseInput()

	freq := map[int]int{}
	for _, num := range right {
		freq[num]++
	}
	sum := 0
	for _, num := range left {
		sum += num * freq[num]
	}
	fmt.Println("part 2:", sum)
}

func ParseInput() (left, right []int) {
	file, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftNum, rightNum := 0, 0
		_, err = fmt.Sscanf(scanner.Text(), "%d   %d", &leftNum, &rightNum)
		if err != nil {
			panic(err)
		}
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	return left, right
}
