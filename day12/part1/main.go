package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	start, dest, elevationMap := read()
	r := process(start, dest, elevationMap)
	fmt.Printf("Answer: %v\n", r)
}

func read() (start, dest aoc.Vector2, elevationMap [][]int8) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	elevationMap = make([][]int8, len(lines))
	for i, line := range lines {
		byteRow := []byte(line)
		row := make([]int8, len(byteRow))
		for j, b := range byteRow {
			if b == 'S' {
				start = aoc.NewVector2(j, i)
				b = 'a'
			} else if b == 'E' {
				dest = aoc.NewVector2(j, i)
				b = 'z'
			}
			row[j] = int8(b)
		}
		elevationMap[i] = row
	}

	return start, dest, elevationMap
}

func process(start, dest aoc.Vector2, elevationMap [][]int8) int {
	w, h := len(elevationMap[0]), len(elevationMap)
	visited := make([][]bool, h)
	for i := range elevationMap {
		visited[i] = make([]bool, w)
	}

	dirs := []aoc.Vector2{
		aoc.NewVector2(-1, 0),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
	}
	steps := 0
	current := []aoc.Vector2{start}
	for !visited[dest.Y][dest.X] {
		steps++
		var next []aoc.Vector2
		for _, c := range current {
			currentElev := elevationMap[c.Y][c.X]
			for _, d := range dirs {
				nc := c.Add(d)
				if nc.X < 0 || nc.X >= w || nc.Y < 0 || nc.Y >= h || visited[nc.Y][nc.X] {
					continue
				}
				if elevationMap[nc.Y][nc.X]-currentElev > 1 {
					continue
				}
				visited[nc.Y][nc.X] = true
				next = append(next, nc)
			}
		}
		current = next
	}

	return steps
}
