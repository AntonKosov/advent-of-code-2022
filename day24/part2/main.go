package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (width, height int, entrance, exit aoc.Vector2, blizzards []blizzard) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	width, height = len(lines[0])-2, len(lines)-2
	entrance, exit = aoc.NewVector2(0, -1), aoc.NewVector2(width-1, height)
	for y := 0; y < height; y++ {
		line := []byte(lines[y+1])
		for x := 0; x < width; x++ {
			var dir aoc.Vector2
			switch v := line[x+1]; v {
			case '.':
				continue
			case '>':
				dir = aoc.NewVector2(1, 0)
			case 'v':
				dir = aoc.NewVector2(0, 1)
			case '<':
				dir = aoc.NewVector2(-1, 0)
			case '^':
				dir = aoc.NewVector2(0, -1)
			default:
				panic(fmt.Sprintf("unknown value: %v", string(v)))
			}
			blizzards = append(blizzards, blizzard{
				pos: aoc.NewVector2(x, y),
				dir: dir,
			})
		}
	}

	return width, height, entrance, exit, blizzards
}

func process(width, height int, entrance, exit aoc.Vector2, blizzards []blizzard) int {
	dirs := []aoc.Vector2{
		aoc.NewVector2(0, 0),
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
	targets := []aoc.Vector2{exit, entrance, exit}
	current := map[aoc.Vector2]struct{}{entrance: {}}
nextMoveBack:
	for minute := 1; ; minute++ {
		next := make(map[aoc.Vector2]struct{}, len(current))
		occupied := moveBlizzards(width, height, blizzards)
		for pos := range current {
			for _, dir := range dirs {
				p := pos.Add(dir)
				if p == exit {
					if len(targets) == 1 {
						return minute
					}
					targets = targets[1:]
					entrance, exit = p, targets[0]
					current = map[aoc.Vector2]struct{}{entrance: {}}
					continue nextMoveBack
				}
				if p.Y >= 0 && p.Y < height && p.X >= 0 && p.X < width && !occupied[p.Y][p.X] {
					next[p] = struct{}{}
				}
			}
		}
		next[entrance] = struct{}{}

		current = next
	}
}

func moveBlizzards(width, height int, blizzards []blizzard) (occupied [][]bool) {
	occupied = make([][]bool, height)
	for i := range occupied {
		occupied[i] = make([]bool, width)
	}

	for i, b := range blizzards {
		b.pos = b.pos.Add(b.dir)
		if b.pos.X < 0 {
			b.pos.X = width - 1
		} else {
			b.pos.X %= width
		}
		if b.pos.Y < 0 {
			b.pos.Y = height - 1
		} else {
			b.pos.Y %= height

		}
		blizzards[i] = b
		occupied[b.pos.Y][b.pos.X] = true
	}

	return occupied
}

type blizzard struct {
	pos aoc.Vector2
	dir aoc.Vector2
}
