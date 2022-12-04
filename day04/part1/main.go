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

func read() []pair {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	var res []pair
	for _, line := range lines {
		lists := strings.Split(line, ",")
		parseList := func(str string) list {
			numbers := strings.Split(str, "-")
			return list{
				from: aoc.StrToInt(numbers[0]),
				to:   aoc.StrToInt(numbers[1]),
			}
		}
		res = append(res, pair{
			first:  parseList(lists[0]),
			second: parseList(lists[1]),
		})
	}

	return res
}

func process(data []pair) int {
	count := 0
	for _, p := range data {
		if p.first.contains(p.second) || p.second.contains(p.first) {
			count++
		}
	}

	return count
}

type list struct {
	from int
	to   int
}

func (l list) contains(al list) bool {
	return l.from <= al.from && l.to >= al.to
}

type pair struct {
	first  list
	second list
}
