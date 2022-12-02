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
		data[i] = round{opponent: shape(parts[0][0]), me: shape(parts[1][0])}
	}

	return data
}

func process(data []round) int {
	shapeScore := map[shape]int{meRock: 1, mePaper: 2, meScissors: 3}
	wins := map[shape]shape{oppRock: mePaper, oppPaper: meScissors, oppScissors: meRock}
	losses := map[shape]shape{oppPaper: meRock, oppRock: meScissors, oppScissors: mePaper}
	totalScore := 0
	for _, r := range data {
		totalScore += shapeScore[r.me]
		if wins[r.opponent] == r.me {
			totalScore += win
		} else if losses[r.opponent] == r.me {
			totalScore += lose
		} else {
			totalScore += draw
		}
	}

	return totalScore
}

type round struct {
	opponent shape
	me       shape
}

const (
	lose = 0
	draw = 3
	win  = 6
)

type shape rune

const (
	oppRock     shape = 'A'
	oppPaper    shape = 'B'
	oppScissors shape = 'C'
	meRock      shape = 'X'
	mePaper     shape = 'Y'
	meScissors  shape = 'Z'
)
