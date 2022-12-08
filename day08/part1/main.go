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
	visible := map[aoc.Vector2[int]]bool{}

	addVisible := func(pos, off aoc.Vector2[int], count int) {
		visible[pos] = true
		tallest := patch[pos.Y][pos.X]
		for i := 1; i < count && tallest < 9; i++ {
			pos = pos.Add(off)
			h := patch[pos.Y][pos.X]
			if h > tallest {
				tallest = h
				visible[pos] = true
			}
		}
	}

	w, h := len(patch[0]), len(patch)

	for i := 0; i < h; i++ {
		addVisible(aoc.NewVector2(0, i), aoc.NewVector2(1, 0), w)
		addVisible(aoc.NewVector2(w-1, i), aoc.NewVector2(-1, 0), w)
	}

	for i := 0; i < w; i++ {
		addVisible(aoc.NewVector2(i, 0), aoc.NewVector2(0, 1), h)
		addVisible(aoc.NewVector2(i, h-1), aoc.NewVector2(0, -1), h)
	}

	return len(visible)
}
