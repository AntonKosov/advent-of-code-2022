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

func read() []command {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	commands := make([]command, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		commands[i] = command{
			dir:   direction(parts[0][0]),
			steps: aoc.StrToInt(parts[1]),
		}
	}

	return commands
}

func process(commands []command) int {
	knots := make([]aoc.Vector2, 10)
	visited := map[aoc.Vector2]bool{knots[0]: true}
	for _, command := range commands {
		dir := dirs[command.dir]
		for i := 0; i < command.steps; i++ {
			knots[0] = knots[0].Add(dir)
			for i := 1; i < len(knots); i++ {
				diff := knots[i-1].Sub(knots[i])
				if diff.Len() <= 1 {
					break
				}
				norm := diff.Norm()
				knots[i] = knots[i].Add(norm)
				if i == len(knots)-1 {
					visited[knots[i]] = true
				}
			}
		}
	}

	return len(visited)
}

type direction rune

const (
	up    direction = 'U'
	down  direction = 'D'
	left  direction = 'L'
	right direction = 'R'
)

type command struct {
	dir   direction
	steps int
}

var dirs map[direction]aoc.Vector2

func init() {
	dirs = map[direction]aoc.Vector2{
		up:    aoc.NewVector2(0, -1),
		down:  aoc.NewVector2(0, 1),
		left:  aoc.NewVector2(-1, 0),
		right: aoc.NewVector2(1, 0),
	}
}
