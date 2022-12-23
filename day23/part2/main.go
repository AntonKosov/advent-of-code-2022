package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() map[aoc.Vector2]bool {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	elves := map[aoc.Vector2]bool{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				elves[aoc.NewVector2(x, y)] = true
			}
		}
	}

	return elves
}

var directions = [][]aoc.Vector2{
	[]aoc.Vector2{{-1, -1}, {0, -1}, {1, -1}},
	[]aoc.Vector2{{-1, 1}, {0, 1}, {1, 1}},
	[]aoc.Vector2{{-1, -1}, {-1, 0}, {-1, 1}},
	[]aoc.Vector2{{1, -1}, {1, 0}, {1, 1}},
}

var cellsAround = []aoc.Vector2{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func process(elves map[aoc.Vector2]bool) int {
	startAdjPos := 0
	for i := 1; ; i++ {
		suggestions := map[aoc.Vector2][]aoc.Vector2{}
		for pos := range elves {
			if !hasElves(elves, pos, cellsAround) {
				continue
			}

			for i := 0; i < len(directions); i++ {
				adjPos := directions[(startAdjPos+i)%len(directions)]
				if !hasElves(elves, pos, adjPos) {
					target := pos.Add(adjPos[1])
					suggestions[target] = append(suggestions[target], pos)
					break
				}
			}

		}

		if len(suggestions) == 0 {
			return i
		}

		for to, from := range suggestions {
			if len(from) > 1 {
				continue
			}
			delete(elves, from[0])
			elves[to] = true
		}

		startAdjPos = (startAdjPos + 1) % len(directions)
	}
}

func hasElves(positions map[aoc.Vector2]bool, position aoc.Vector2, offsets []aoc.Vector2) bool {
	for _, offset := range offsets {
		if positions[position.Add(offset)] {
			return true
		}
	}

	return false
}
