package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() []monkey {
	lines := aoc.ReadAllInput()
	monkeys := make([]monkey, len(lines)/7)
	for i := range monkeys {
		idx := i * 7
		m := monkey{}

		startingItems := strings.Split(lines[idx+1][18:], ", ")
		for _, item := range startingItems {
			m.items = append(m.items, aoc.StrToInt(item))
		}

		ops := strings.Split(lines[idx+2][23:], " ")
		itself := ops[1] == "old"
		var opFunc func(old int) int
		switch op := ops[0]; op {
		case "*":
			if itself {
				opFunc = func(old int) int { return old * old }
			} else {
				v := aoc.StrToInt(ops[1])
				opFunc = func(old int) int { return old * v }
			}
		case "+":
			if itself {
				opFunc = func(old int) int { return old + old }
			} else {
				v := aoc.StrToInt(ops[1])
				opFunc = func(old int) int { return old + v }
			}
		default:
			panic(fmt.Sprintf("unknown operation: %v", op))
		}
		m.operation = opFunc

		m.test = test{
			num:      aoc.StrToInt(lines[idx+3][21:]),
			positive: aoc.StrToInt(lines[idx+4][29:]),
			negative: aoc.StrToInt(lines[idx+5][30:]),
		}

		monkeys[i] = m
	}

	return monkeys
}

func process(monkeys []monkey) int {
	n := len(monkeys)
	// All test numbers are prime. So, it's possible to keep modulus only.
	mod := 1
	for _, m := range monkeys {
		mod *= m.test.num
	}
	inspections := make([]int, n)
	for i := 0; i < 10_000; i++ {
		for j := range monkeys {
			m := &monkeys[j]
			inspections[j] += len(m.items)
			for len(m.items) > 0 {
				item := m.items[0]
				m.items = m.items[1:]
				item = m.operation(item) % mod
				target := m.test.negative
				if item%m.test.num == 0 {
					target = m.test.positive
				}
				monkeys[target].items = append(monkeys[target].items, item)
			}
		}
	}

	sort.Ints(inspections)

	return inspections[n-1] * inspections[n-2]
}

type test struct {
	num      int
	positive int
	negative int
}
type monkey struct {
	items     []int
	operation func(old int) int
	test      test
}
