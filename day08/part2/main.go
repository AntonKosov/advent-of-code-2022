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

func read() [][]byte {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	patch := make([][]byte, len(lines))
	for i, line := range lines {
		row := make([]byte, len(line))
		for j, r := range line {
			row[j] = byte(r - '0')
		}
		patch[i] = row
	}

	return patch
}

func process(patch [][]byte) int {
	width, height := len(patch[0]), len(patch)

	score := func(x, y int) int {
		h := patch[y][x]
		countScore := func(off aoc.Vector2[int], count int) int {
			s := 0
			pos := aoc.NewVector2(x, y)
			for i := 0; i < count; i++ {
				pos = pos.Add(off)
				v := patch[pos.Y][pos.X]
				s++
				if v >= h {
					break
				}
			}
			return s
		}
		up := countScore(aoc.NewVector2(0, -1), y)
		down := countScore(aoc.NewVector2(0, 1), height-y-1)
		left := countScore(aoc.NewVector2(-1, 0), x)
		right := countScore(aoc.NewVector2(1, 0), width-x-1)

		return up * right * down * left
	}

	// monotonic stack could reduce the runtime complexity significantly
	maxScore := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maxScore = aoc.Max(maxScore, score(x, y))
		}
	}

	return maxScore
}
