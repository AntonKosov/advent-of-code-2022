package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	stacks, moves := read()
	r := process(stacks, moves)
	fmt.Printf("Answer: %v\n", r)
}

func read() ([][]rune, []move) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	stacks := [][]rune{}
	addCrate := func(str string) {
		stacks = append(stacks, []rune(str))
	}
	addCrate("HRBDZFLS")
	addCrate("TBMZR")
	addCrate("ZLCHNS")
	addCrate("SCFJ")
	addCrate("PGHWRZB")
	addCrate("VJZGDNMT")
	addCrate("GLNWFSPQ")
	addCrate("MZR")
	addCrate("MCLGVRT")

	var moves []move
	for i := 10; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Split(line, " ")
		moves = append(moves, move{
			count: aoc.StrToInt(parts[1]),
			from:  aoc.StrToInt(parts[3]) - 1,
			to:    aoc.StrToInt(parts[5]) - 1,
		})
	}

	return stacks, moves
}

func process(stacks [][]rune, moves []move) string {
	for _, m := range moves {
		stackFrom, stackTo := stacks[m.from], stacks[m.to]
		stacks[m.to] = append(stackTo, stackFrom[len(stackFrom)-m.count:]...)
		stacks[m.from] = stackFrom[:len(stackFrom)-m.count]
	}

	var res strings.Builder
	for _, crate := range stacks {
		res.WriteRune(crate[len(crate)-1])
	}

	return res.String()
}

type move struct {
	from, to, count int
}
