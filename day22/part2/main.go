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
	pos, orientation := aoc.NewVector2(cubeSize, 0), right

	for _, command := range commands {
		switch command {
		case turnLeft:
			orientation = orientation.RotateLeft()
		case turnRight:
			orientation = orientation.RotateRight()
		default:
			steps := aoc.StrToInt(command)
			for i := 0; i < steps; i++ {
				nextPos, nextDir := grid.nextPos(pos, orientation)
				if grid[nextPos.Y][nextPos.X] == cellWall {
					break
				}
				pos, orientation = nextPos, nextDir
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

const cubeSize = 50

var dirsScore = []aoc.Vector2{
	aoc.NewVector2(1, 0),
	aoc.NewVector2(0, 1),
	aoc.NewVector2(-1, 0),
	aoc.NewVector2(0, -1),
}

type board [][]cell

func (b board) nextPos(pos, dir aoc.Vector2) (nextPos, nextDir aoc.Vector2) {
	wr := wrapper{
		direction: dir,
		sideX:     pos.X / cubeSize,
		sideY:     pos.Y / cubeSize,
	}
	if f := wrapperFuncs[wr]; f != nil {
		pos, dir = f(pos)
	} else {
		pos = pos.Add(dir)
	}

	return pos, dir
}

var (
	up, down, left, right aoc.Vector2
)

func init() {
	up = aoc.NewVector2(0, -1)
	down = aoc.NewVector2(0, 1)
	left = aoc.NewVector2(-1, 0)
	right = aoc.NewVector2(1, 0)
}

type wrapper struct {
	direction aoc.Vector2
	sideX     int
	sideY     int
}

type wrapperFunc func(pos aoc.Vector2) (nextPos, nextDir aoc.Vector2)

var wrapperFuncs map[wrapper]wrapperFunc

func init() {
	wrapperFuncs = map[wrapper]wrapperFunc{
		{up, 1, 0}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.Y != 0 {
				return pos.Add(up), up
			}
			return aoc.NewVector2(0, cubeSize*3+pos.X-cubeSize), right
		},
		{up, 2, 0}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.Y != 0 {
				return pos.Add(up), up
			}
			return aoc.NewVector2(pos.X-cubeSize*2, cubeSize*4-1), up
		},
		{up, 0, 2}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.Y != cubeSize*2 {
				return pos.Add(up), up
			}
			return aoc.NewVector2(cubeSize, cubeSize+pos.X), right
		},
		{down, 2, 0}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.Y != cubeSize-1 {
				return pos.Add(down), down
			}
			return aoc.NewVector2(cubeSize*2-1, cubeSize+pos.X-cubeSize*2), left
		},
		{down, 1, 2}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.Y != cubeSize*3-1 {
				return pos.Add(down), down
			}
			return aoc.NewVector2(cubeSize-1, cubeSize*3+pos.X-cubeSize), left
		},
		{down, 0, 3}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.Y != cubeSize*4-1 {
				return pos.Add(down), down
			}
			return aoc.NewVector2(cubeSize*2+pos.Y-cubeSize*3, 0), down
		},
		{left, 1, 0}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != cubeSize {
				return pos.Add(left), left
			}
			return aoc.NewVector2(0, cubeSize*2+(cubeSize-pos.Y-1)), right
		},
		{left, 1, 1}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != cubeSize {
				return pos.Add(left), left
			}
			return aoc.NewVector2(pos.Y-cubeSize, cubeSize*2), down
		},
		{left, 0, 2}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != 0 {
				return pos.Add(left), left
			}
			return aoc.NewVector2(cubeSize, cubeSize-(pos.Y-cubeSize*2)-1), right
		},
		{left, 0, 3}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != 0 {
				return pos.Add(left), left
			}
			return aoc.NewVector2(cubeSize+pos.Y-cubeSize*3, 0), down
		},
		{right, 2, 0}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != cubeSize*3-1 {
				return pos.Add(right), right
			}
			return aoc.NewVector2(cubeSize*2-1, cubeSize*2+(cubeSize-pos.Y-1)), left
		},
		{right, 1, 1}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != cubeSize*2-1 {
				return pos.Add(right), right
			}
			return aoc.NewVector2(cubeSize*2+pos.Y-cubeSize, cubeSize-1), up
		},
		{right, 1, 2}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != cubeSize*2-1 {
				return pos.Add(right), right
			}
			return aoc.NewVector2(cubeSize*3-1, cubeSize-(pos.Y-cubeSize*2)-1), left
		},
		{right, 0, 3}: func(pos aoc.Vector2) (nextPos aoc.Vector2, nextDir aoc.Vector2) {
			if pos.X != cubeSize-1 {
				return pos.Add(right), right
			}
			return aoc.NewVector2(cubeSize*2+pos.X, 0), down
		},
	}
}
