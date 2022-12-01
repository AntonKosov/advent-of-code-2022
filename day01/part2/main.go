package main

import (
	"fmt"
	"sort"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() [][]int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	var data [][]int
	for i := 0; i < len(lines); i++ {
		var list []int
		for ; i < len(lines) && len(lines[i]) > 0; i++ {
			list = append(list, aoc.StrToInt(lines[i]))
		}

		data = append(data, list)
	}

	return data
}

func process(data [][]int) int {
	sums := make([]int, 0, len(data))
	for _, list := range data {
		sum := 0
		for _, calories := range list {
			sum += calories
		}
		sums = append(sums, sum)
	}

	sort.Ints(sums)
	n := len(sums)

	return sums[n-1] + sums[n-2] + sums[n-3]
}
