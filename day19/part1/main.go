package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

const materials = 4
const minutes = 24
const shrinkThreshold = 10_000

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []blueprint {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	blueprints := make([]blueprint, len(lines))
	for i, line := range lines {
		values := aoc.ParseInts(line)
		req := [materials]vector{
			{int16(values[1]), 0, 0, 0},
			{int16(values[2]), 0, 0, 0},
			{int16(values[3]), int16(values[4]), 0, 0},
			{int16(values[5]), 0, int16(values[6]), 0},
		}
		blueprints[i] = blueprint{id: values[0], req: req}
	}

	return blueprints
}

func process(blueprints []blueprint) int {
	sum := 0
	for _, b := range blueprints {
		mp := int(maxProducts(b))
		sum += b.id * mp
		fmt.Println("id:", b.id, "max:", mp)
	}

	return sum
}

func maxProducts(b blueprint) int16 {
	cycle := map[state]struct{}{
		{robots: vector{1, 0, 0, 0}, minutesLeft: minutes}: {},
	}
	for i := byte(minutes); i > 0; i-- {
		cycle = processCycle(cycle, b)
	}

	max := int16(0)
	for s := range cycle {
		max = aoc.Max(max, s.resources[materials-1])
	}

	return max
}

func processCycle(cycle map[state]struct{}, b blueprint) (nextCycle map[state]struct{}) {
	nextCycle = map[state]struct{}{}
	maxProducts := vector{0, 0, 0, 0xff}
	for i := 0; i < materials-1; i++ {
		for j := 0; j < materials; j++ {
			maxProducts[i] = aoc.Max(maxProducts[i], b.req[j][i])
		}
	}

	for s := range cycle {
		nextStates := map[state]struct{}{
			s: {}, // just collect resources
		}
		for i := 0; i < materials; i++ {
			buildRobots(s, b, nextStates)
		}
		for ns := range nextStates {
			ns.resources = ns.resources.add(s.robots)
			ns.minutesLeft--
			nextCycle[ns] = struct{}{}
		}
	}

	shrink(nextCycle)

	return nextCycle
}

func buildRobots(s state, b blueprint, nextStates map[state]struct{}) {
	for i := 0; i < materials; i++ {
		if resourcesLeft := s.resources.sub(b.req[i]); resourcesLeft.valid() {
			ls := s
			ls.resources = resourcesLeft
			ls.robots[i]++
			nextStates[ls] = struct{}{}
		}
	}
}

func shrink(states map[state]struct{}) {
	stateScore := func(s state) int {
		score := 0
		for i := materials - 1; i >= 0; i-- {
			score = (score << 8) + int(s.resources[i]) + int(s.robots[i])
		}
		return score
	}
	for len(states) > shrinkThreshold {
		sum := 0
		for s := range states {
			sum += stateScore(s)
		}
		threshold := sum / len(states)
		toRemove := make([]state, 0, len(states)/2)
		for s := range states {
			if stateScore(s) < threshold {
				toRemove = append(toRemove, s)
			}
		}
		for _, s := range toRemove {
			delete(states, s)
		}
	}
}

type blueprint struct {
	id  int
	req [materials]vector
}

type state struct {
	resources   vector
	robots      vector
	minutesLeft byte
}

type vector [materials]int16

func (v vector) sub(av vector) vector {
	res := vector{}
	for i := 0; i < materials; i++ {
		res[i] = v[i] - av[i]
	}
	return res
}

func (v vector) add(av vector) vector {
	res := vector{}
	for i := 0; i < materials; i++ {
		res[i] = v[i] + av[i]
	}
	return res
}

func (v vector) valid() bool {
	for _, c := range v {
		if c < 0 {
			return false
		}
	}
	return true
}
