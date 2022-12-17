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

func read() (rates []int, tunnels [][]int, start int) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	valves := make(map[string]int, len(lines))
	for i, line := range lines {
		valve := line[6:8]
		valves[valve] = i
	}

	rates = make([]int, len(valves))
	tunnels = make([][]int, len(valves))

	for valveID, line := range lines {
		parts := strings.Split(line[23:], " ")
		rate := aoc.StrToInt(parts[0][:len(parts[0])-1])
		var connectedTunnels []int
		for i := 5; i < len(parts); i++ {
			c := parts[i]
			connectedTunnels = append(connectedTunnels, valves[c[:2]])
		}

		rates[valveID] = rate
		tunnels[valveID] = connectedTunnels
	}

	return rates, tunnels, valves["AA"]
}

const valvesCount = 60

func process(rates []int, tunnels [][]int, start int) int {
	opened := [valvesCount]bool{}
	memo := map[cache]int{}
	var findMaxPressure func(currentValve int, secondsLeft int) int
	findMaxPressure = func(currentValve int, secondsLeft int) int {
		if secondsLeft <= 1 {
			return 0
		}

		c := cache{valve: currentValve, seconds: secondsLeft, opened: opened}
		if cv, ok := memo[c]; ok {
			return cv
		}

		openCurrentValve := 0
		if !opened[currentValve] && rates[currentValve] > 0 {
			opened[currentValve] = true

			maxFollowing := 0
			for _, tunnel := range tunnels[currentValve] {
				maxFollowing = aoc.Max(maxFollowing, findMaxPressure(tunnel, secondsLeft-2))
			}
			openCurrentValve = maxFollowing + rates[currentValve]*(secondsLeft-1)

			opened[currentValve] = false
		}

		skipCurrentValve := 0
		for _, tunnel := range tunnels[currentValve] {
			skipCurrentValve = aoc.Max(skipCurrentValve, findMaxPressure(tunnel, secondsLeft-1))
		}

		maxPressure := aoc.Max(openCurrentValve, skipCurrentValve)
		memo[c] = maxPressure

		return maxPressure
	}

	return findMaxPressure(start, 30)
}

type cache struct {
	valve   int
	seconds int
	opened  [valvesCount]bool
}
