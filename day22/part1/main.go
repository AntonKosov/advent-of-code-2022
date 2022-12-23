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

func read() (grid board, commands []string) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	h := len(lines) - 2
	w := 0
	for i := 0; i < h; i++ {
		w = aoc.Max(w, len(lines[i]))
	}

	grid = make([][]cell, h)
	for i := 0; i < h; i++ {
		row := make([]cell, w)
		line := lines[i]
		for j, r := range line {
			if r != ' ' {
				row[j] = cell(r)
			}
		}
		grid[i] = row
	}

	parts := strings.Split(lines[len(lines)-1], turnLeft)
	for _, part := range parts {
		parts2 := strings.Split(part, turnRight)
		for _, steps := range parts2 {
			commands = append(commands, steps, turnRight)
		}
		aoc.RemoveLastItem(&commands)
		commands = append(commands, turnLeft)
	}
	aoc.RemoveLastItem(&commands)

	return grid, commands
}

func process(grid board, commands []string) int {
	pos := grid.nextPos(aoc.NewVector2(-1, 0), aoc.NewVector2(1, 0))
	orientation := aoc.NewVector2(1, 0)
	for _, command := range commands {
		switch command {
		case turnLeft:
			orientation = orientation.RotateLeft()
		case turnRight:
			orientation = orientation.RotateRight()
		default:
			steps := aoc.StrToInt(command)
			for i := 0; i < steps; i++ {
				nextPos := grid.nextPos(pos, orientation)
				if grid[nextPos.Y][nextPos.X] == cellWall {
					break
				}
				pos = nextPos
			}
		}
	}

	psw := 1000*(pos.Y+1) + 4*(pos.X+1)
	for i, dir := range dirsScore {
		if orientation == dir {
			psw += i
			break
		}
	}

	return psw
}

type cell byte

const (
	cellNone  cell = 0
	cellWall  cell = '#'
	cellEmpty cell = '.'
)

const (
	turnLeft  = "L"
	turnRight = "R"
)

var dirsScore = []aoc.Vector2{
	aoc.NewVector2(1, 0),
	aoc.NewVector2(0, 1),
	aoc.NewVector2(-1, 0),
	aoc.NewVector2(0, -1),
}

type board [][]cell

func (b board) nextPos(currentPos, dir aoc.Vector2) aoc.Vector2 {
	h, w := len(b), len(b[0])
	for {
		currentPos = currentPos.Add(dir)
		if currentPos.X < 0 {
			currentPos.X = w - 1
		}
		if currentPos.Y < 0 {
			currentPos.Y = h - 1
		}
		currentPos.X %= w
		currentPos.Y %= h

		if b[currentPos.Y][currentPos.X] != cellNone {
			return currentPos
		}
	}
}
