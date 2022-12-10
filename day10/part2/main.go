package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	data := read()
	lines := process(data)
	for _, line := range lines {
		fmt.Println(line)
	}
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

func process(commands []command) []string {
	const w, h = 40, 6
	lines := make([]string, 0, h)
	register := 1
	var sb strings.Builder
	sb.WriteRune('#')
	cycle := 0
	for _, command := range commands {
		for completed := false; !completed; {
			cycle++
			if cycle%w == 0 && cycle > 1 {
				lines = append(lines, sb.String())
				if len(lines) == h {
					return lines
				}
				sb = strings.Builder{}
			}
			completed = command.execute(&register)
			if aoc.Abs((cycle)%w-register) <= 1 {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
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
