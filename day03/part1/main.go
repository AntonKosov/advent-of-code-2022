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
	priorities := buildPriorities()
	sum := 0
	for _, line := range data {
		bytes := []byte(line)
		half := len(bytes) / 2
		firstCompartment := map[byte]bool{}
		for i := 0; i < half; i++ {
			firstCompartment[bytes[i]] = true
		}
		for i := half; i < len(bytes); i++ {
			b := bytes[i]
			if firstCompartment[b] {
				sum += priorities[b]
				break
			}
		}
	}

	return sum
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
