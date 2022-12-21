package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	moves := make([]int, len(lines))
	for i, line := range lines {
		move := aoc.StrToInt(line)
		moves[i] = move
	}

	return moves
}

func process(data []int) int {
	nodes := buildLinkedList(data)

	for _, n := range nodes {
		count := n.val
		switch {
		case count > 0:
			n.moveForward(count, len(nodes))
		case count < 0:
			n.moveBackward(-count, len(nodes))
		}
	}

	for _, n := range nodes {
		if n.val == 0 {
			sum := 0
			for i := 0; i < 3; i++ {
				n = n.move(1000, len(nodes))
				sum += n.val
			}
			return sum
		}
	}

	panic("zero value not found")
}

func buildLinkedList(data []int) []*node {
	nodes := make([]*node, len(data))
	start := &node{}
	tail := start
	for i, n := range data {
		tail.next = &node{val: n, prev: tail}
		tail = tail.next
		nodes[i] = tail
	}
	tail.next = start.next
	start.next.prev = tail

	return nodes
}

type node struct {
	val  int
	next *node
	prev *node
}

func (n *node) move(count int, length int) *node {
	for i := 0; i < count; i++ {
		n = n.next
	}

	return n
}

func (n *node) moveForward(count, length int) {
	for ; count > length; count -= length {
		movingNode := n.next
		movingNode.delete()
		movingNode.insertAfter(n.prev)
	}
	for i := 0; i < count; i++ {
		next := n.next
		n.delete()
		n.insertAfter(next)
	}
}

func (n *node) moveBackward(count, length int) {
	for ; count > length; count -= length {
		movingNode := n.prev
		movingNode.delete()
		movingNode.insertAfter(n)
	}
	for i := 0; i < count; i++ {
		prev := n.prev.prev
		n.delete()
		n.insertAfter(prev)
	}
}

func (n *node) delete() {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
}

func (n *node) insertAfter(an *node) {
	n.next = an.next
	n.prev = an
	n.next.prev = n
	n.prev.next = n
}
