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

func read() []byte {
	lines := aoc.ReadAllInput()

	return []byte(lines[0])
}

func process(data []byte) int {
	visited := make([]bool, 'z'-'a'+1)
	left, right := 0, -1
	for {
		right++
		b := data[right] - 'a'
		if visited[b] {
			for {
				rb := data[left] - 'a'
				visited[rb] = false
				left++
				if b == rb {
					break
				}
			}
		}
		if right-left == 3 {
			return right + 1
		}
		visited[b] = true
	}
}
