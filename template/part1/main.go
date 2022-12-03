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

func read() []string {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	return lines
}

func process(data []string) int {
	return 0
}
