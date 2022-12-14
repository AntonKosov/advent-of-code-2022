package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() [][]aoc.Vector2 {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	rocks := make([][]aoc.Vector2, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " -> ")
		rl := make([]aoc.Vector2, len(parts))
		for j, part := range parts {
			c := strings.Split(part, ",")
			rl[j] = aoc.NewVector2(aoc.StrToInt(c[0]), aoc.StrToInt(c[1]))
		}

		rocks[i] = rl
	}

	return rocks
}

func process(data [][]aoc.Vector2) int {
	path := []aoc.Vector2{aoc.NewVector2(500, 0)}
	rockMap := buildRockMap(data)
	width, height := len(rockMap[0]), len(rockMap)
	count := 0
	dirs := []aoc.Vector2{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(-1, 1),
		aoc.NewVector2(1, 1),
	}

	for {
		pos := aoc.LastItem(path)
		for {
			startPos := pos
			for _, dir := range dirs {
				c := pos.Add(dir)
				if c.Y < 0 || c.Y >= height || c.X < 0 || c.X >= width {
					return count
				}
				if !rockMap[c.Y][c.X] {
					pos = c
					path = append(path, pos)
					break
				}
			}
			if startPos == pos {
				aoc.RemoveLastItem(&path)
				rockMap[pos.Y][pos.X] = true
				break
			}
		}
		count++
	}
}

func buildRockMap(data [][]aoc.Vector2) [][]bool {
	w, h := 0, 0
	for _, row := range data {
		for _, c := range row {
			w = aoc.Max(w, c.X)
			h = aoc.Max(h, c.Y)
		}
	}
	w++
	h++

	m := make([][]bool, h)
	for i := range m {
		m[i] = make([]bool, w)
	}

	for _, line := range data {
		for i := 0; i < len(line)-1; i++ {
			from, to := line[i], line[i+1]
			dir := to.Sub(from).Norm()
			for {
				m[from.Y][from.X] = true
				if from == to {
					break
				}
				from = from.Add(dir)
			}
		}
	}

	return m
}
