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
	head := aoc.NewVector2(0, 0)
	tail := head
	visited := map[aoc.Vector2]bool{tail: true}
	for _, command := range commands {
		dir := dirs[command.dir]
		for i := 0; i < command.steps; i++ {
			head = head.Add(dir)
			diff := head.Sub(tail)
			if diff.Len() > 1 {
				norm := diff.Norm()
				tail = tail.Add(norm)
				visited[tail] = true
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
