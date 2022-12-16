package main

import (
	"fmt"
	"sort"
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
	const (
		minCoordinate = 0
		maxCoordinate = 4_000_000
	)
	for y := minCoordinate; y <= maxCoordinate; y++ {
		intervals := make([]interval, 0, len(points))
		for _, p := range points {
			dst := p.sensor.Sub(p.beacon).ManhattanDst()
			minY, maxY := p.sensor.Y-dst, p.sensor.Y+dst
			if y < minY || y > maxY {
				continue
			}
			diffX := dst - aoc.Abs(y-p.sensor.Y)
			from, to := p.sensor.X-diffX, p.sensor.X+diffX
			intervals = append(intervals, interval{from: from, to: to})
		}

		sort.Slice(intervals, func(i, j int) bool { return intervals[i].from < intervals[j].from })
		for len(intervals) > 1 {
			if c, ok := intervals[0].combine(intervals[1]); ok {
				intervals[1] = c
				intervals = intervals[1:]
			} else {
				return maxCoordinate*(intervals[0].to+1) + y
			}
		}

		if intervals[0].from > minCoordinate {
			return y
		}

		if intervals[0].to < maxCoordinate {
			return maxCoordinate*maxCoordinate + y
		}
	}

	panic("solution not found")
}

type point struct {
	sensor aoc.Vector2
	beacon aoc.Vector2
}

type interval struct {
	from int
	to   int
}

func (i interval) combine(ai interval) (interval, bool) {
	if i.from > ai.from {
		i, ai = ai, i
	}

	if ai.from > i.to+1 {
		return interval{}, false
	}

	return interval{
		from: i.from,
		to:   aoc.Max(i.to, ai.to),
	}, true
}
