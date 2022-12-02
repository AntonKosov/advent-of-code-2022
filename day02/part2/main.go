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

func read() []round {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	data := make([]round, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		data[i] = round{opponent: shape(parts[0][0]), result: result(parts[1][0])}
	}

	return data
}

func process(data []round) int {
	shapeScore := map[shape]int{rock: 1, paper: 2, scissors: 3}
	wins := map[shape]shape{rock: paper, paper: scissors, scissors: rock}
	losses := map[shape]shape{paper: rock, rock: scissors, scissors: paper}
	totalScore := 0
	for _, r := range data {
		switch r.result {
		case win:
			totalScore += shapeScore[wins[r.opponent]] + winScore
		case draw:
			totalScore += shapeScore[r.opponent] + drawScore
		case lose:
			totalScore += shapeScore[losses[r.opponent]] + loseScore
		default:
			panic("unknown result")
		}
	}

	return totalScore
}

type round struct {
	opponent shape
	result   result
}

const (
	loseScore = 0
	drawScore = 3
	winScore  = 6
)

type shape rune

const (
	rock     shape = 'A'
	paper    shape = 'B'
	scissors shape = 'C'
)

type result rune

const (
	lose result = 'X'
	draw result = 'Y'
	win  result = 'Z'
)
