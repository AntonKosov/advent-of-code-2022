package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() [][]string {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	var res [][]string
	for i := 0; i < len(lines)/3; i++ {
		res = append(res, []string{lines[i*3], lines[i*3+1], lines[i*3+2]})
	}

	return res
}

func process(groups [][]string) int {
	priorities := buildPriorities()
	sum := 0
	for _, group := range groups {
		ci := commonItem(group)
		sum += priorities[ci]
	}

	return sum
}

func commonItem(group []string) byte {
	count := map[rune]int{}
	countRunes := func(str string) {
		counted := map[rune]bool{}
		for _, r := range str {
			if !counted[r] {
				counted[r] = true
				count[r]++
			}
		}
	}

	for _, line := range group {
		countRunes(line)
	}

	for r, c := range count {
		if c == 3 {
			return byte(r)
		}
	}

	panic("common item not found")
}

func buildPriorities() map[byte]int {
	priorities := make(map[byte]int, 2*('z'-'a'+1))

	for r := 'a'; r <= 'z'; r++ {
		priorities[byte(r)] = int(r - 'a' + 1)
	}

	for r := 'A'; r <= 'Z'; r++ {
		priorities[byte(r)] = int(r - 'A' + 27)
	}

	return priorities
}
