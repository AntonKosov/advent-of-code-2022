package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []byte {
	return []byte(aoc.ReadAllInput()[0])
}

const (
	freeCell = '.'
	rockCell = '#'
	left     = '<'
)

var rocks = []rock{
	newRock("####"),
	newRock(".#.", "###", ".#."),
	newRock("..#", "..#", "###"),
	newRock("#", "#", "#", "#"),
	newRock("##", "##"),
}

func process(wind []byte) int {
	ch := newChamber()
	windIndex := 0
	rockIndex := 0

	for i := 0; i < 2022; i++ {
		ch.addRock(rocks[rockIndex])
		for placed := false; !placed; {
			placed = ch.move(wind[windIndex])

			windIndex = (windIndex + 1) % len(wind)
		}

		rockIndex = (rockIndex + 1) % len(rocks)
	}

	return ch.highestRock
}

type rock struct {
	cells  []aoc.Vector2
	width  int
	height int
}

func newRock(rows ...string) rock {
	r := rock{height: len(rows)}
	for y, row := range rows {
		r.width = aoc.Max(r.width, len(row))
		for x, v := range row {
			if v != freeCell {
				r.cells = append(r.cells, aoc.NewVector2(x, y))
			}
		}
	}

	return r
}

type chamber struct {
	highestRock  int
	cells        [][]byte
	fallingRock  *rock
	rockPosition aoc.Vector2
}

func newChamber() chamber {
	floor := []byte("#########")
	return chamber{cells: [][]byte{floor}}
}

func (c *chamber) addRock(r rock) {
	reqHeight := c.highestRock + 3 + r.height + 1
	for len(c.cells) > reqHeight {
		c.cells = c.cells[0 : len(c.cells)-1]
	}
	for len(c.cells) < reqHeight {
		wall := []byte("#.......#")
		c.cells = append(c.cells, wall)
	}

	c.fallingRock = &r
	c.rockPosition = aoc.NewVector2(3, 0)
}

func (c *chamber) move(dir byte) (placed bool) {
	horOffset := aoc.NewVector2(1, 0)
	if dir == left {
		horOffset.X *= -1
	}

	c.moveTo(horOffset)

	return !c.moveTo(aoc.NewVector2(0, 1))
}

func (c *chamber) moveTo(offset aoc.Vector2) bool {
	chamberOffset := offset.Add(c.rockPosition)
	for _, cell := range c.fallingRock.cells {
		if !c.isEmpty(cell.Add(chamberOffset)) {
			if offset.Y > 0 {
				c.place()
			}
			return false
		}
	}
	c.rockPosition = c.rockPosition.Add(offset)
	return true
}

func (c *chamber) place() {
	for _, cell := range c.fallingRock.cells {
		c.addRockCell(c.rockPosition.Add(cell))
	}

	c.fallingRock = nil
	c.highestRock = aoc.Max(c.highestRock, len(c.cells)-c.rockPosition.Y-1)
}

func (c *chamber) isEmpty(pos aoc.Vector2) bool {
	return c.cells[len(c.cells)-pos.Y-1][pos.X] == freeCell
}

func (c *chamber) addRockCell(pos aoc.Vector2) {
	c.cells[len(c.cells)-pos.Y-1][pos.X] = rockCell
}

/*func (c *chamber) print() {
	fmt.Println()
	for j := 1; j <= aoc.Min(9, len(c.cells)); j++ {
		fmt.Println(string(c.cells[len(c.cells)-j]))
	}
}*/
