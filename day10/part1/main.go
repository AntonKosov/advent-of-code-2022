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

	ops := make([]command, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		var cmd command
		switch c := parts[0]; c {
		case "noop":
			cmd = noopCommand{}
		case "addx":
			cmd = &addCommand{num: aoc.StrToInt(parts[1])}
		default:
			panic(fmt.Sprintf("unknown command: %v", c))
		}
		ops[i] = cmd
	}

	return ops
}

func process(commands []command) int {
	cycles := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	register := 1
	cycle := 0
	for _, command := range commands {
		for completed := false; !completed; {
			cycle++
			if cycle == cycles[0] {
				sum += cycle * register
				if len(cycles) == 1 {
					return sum
				}
				cycles = cycles[1:]
			}
			completed = command.execute(&register)
		}
	}

	panic("invalid input")
}

type command interface {
	execute(register *int) (completed bool)
}

type noopCommand struct{}

func (c noopCommand) execute(_ *int) bool {
	return true
}

type addCommand struct {
	inProgress bool
	num        int
}

func (c *addCommand) execute(register *int) bool {
	if c.inProgress {
		*register += c.num
	}

	c.inProgress = !c.inProgress

	return !c.inProgress
}
