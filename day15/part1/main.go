package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	points := read()
	r := process(points)
	fmt.Printf("Answer: %v\n", r)
}

func read() []point {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	parse := func(str string) aoc.Vector2 {
		parts := strings.Split(str, ", ")
		return aoc.NewVector2(aoc.StrToInt(parts[0][2:]), aoc.StrToInt(parts[1][2:]))
	}

	points := make([]point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line[10:], ": closest beacon is at ")
		points[i] = point{
			sensor: parse(parts[0]),
			beacon: parse(parts[1]),
		}
	}

	return points
}

func process(points []point) int {
	const row = 2_000_000
	beacons := map[aoc.Vector2]bool{}
	for _, p := range points {
		if p.beacon.Y == row {
			beacons[p.beacon] = true
		}
	}
	minX, maxX := detectBoarder(points)
	count := 0
	for x := minX; x <= maxX; x++ {
		pos := aoc.NewVector2(x, row)
		if beacons[pos] {
			continue
		}
		for _, p := range points {
			posDst := p.sensor.Sub(pos).ManhattanDst()
			beaconDst := p.sensor.Sub(p.beacon).ManhattanDst()
			if posDst <= beaconDst {
				count++
				break
			}
		}
	}

	return count
}

func detectBoarder(points []point) (minX, maxX int) {
	minX = points[0].sensor.X
	maxX = minX
	for _, p := range points {
		diffX := p.beacon.Sub(p.sensor).ManhattanDst()
		minX = aoc.Min(minX, p.sensor.X-diffX)
		maxX = aoc.Max(maxX, p.sensor.X+diffX)
	}

	return minX, maxX
}

type point struct {
	sensor aoc.Vector2
	beacon aoc.Vector2
}
