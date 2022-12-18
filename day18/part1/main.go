package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []aoc.Vector3 {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	cubes := make([]aoc.Vector3, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		cubes[i] = aoc.NewVector3(
			aoc.StrToInt(parts[0]),
			aoc.StrToInt(parts[1]),
			aoc.StrToInt(parts[2]),
		)
	}

	return cubes
}

func process(cubes []aoc.Vector3) int {
	cube := buildCube(cubes)
	edges := 0
	dirs := []aoc.Vector3{
		aoc.NewVector3(0, 0, 1),
		aoc.NewVector3(0, 0, -1),
		aoc.NewVector3(0, 1, 0),
		aoc.NewVector3(0, -1, 0),
		aoc.NewVector3(1, 0, 0),
		aoc.NewVector3(-1, 0, 0),
	}

	for _, c := range cubes {
		for _, dir := range dirs {
			p := c.Add(dir)
			if p.X < 0 || p.Y < 0 || p.Z < 0 {
				edges++
				continue
			}
			if !cube[p.X][p.Y][p.Z] {
				edges++
			}
		}
	}

	return edges
}

func buildCube(cubes []aoc.Vector3) [][][]bool {
	max := aoc.Vector3{}
	for _, c := range cubes {
		max.X = aoc.Max(max.X, c.X)
		max.Y = aoc.Max(max.Y, c.Y)
		max.Z = aoc.Max(max.Z, c.Z)
	}

	cube := make([][][]bool, max.X+2)
	for x := range cube {
		yz := make([][]bool, max.Y+2)
		for y := range yz {
			yz[y] = make([]bool, max.Z+2)
		}

		cube[x] = yz
	}

	for _, c := range cubes {
		cube[c.X][c.Y][c.Z] = true
	}

	return cube
}
