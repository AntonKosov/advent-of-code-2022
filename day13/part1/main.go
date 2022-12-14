package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	pairs := read()
	r := process(pairs)
	fmt.Printf("Answer: %v\n", r)
}

func read() []pair {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	n := len(lines) / 3
	pairs := make([]pair, n)
	for i := 0; i < n; i++ {
		line1 := lines[i*3]
		line2 := lines[i*3+1]
		line1 = line1[1 : len(line1)-1]
		line2 = line2[1 : len(line2)-1]
		pairs[i] = pair{
			node1: parse([]byte(line1)),
			node2: parse([]byte(line2)),
		}
	}

	return pairs
}

func parse(str []byte) node {
	if len(str) == 0 {
		return node{children: []*node{}}
	}
	n := node{}
	for i := 0; i < len(str); i++ {
		b := str[i]
		var child node
		if b == '[' {
			child, i = parseBracket(str, i)
		} else {
			child, i = parseNumber(str, i)
		}
		n.children = append(n.children, &child)
	}

	return n
}

func parseBracket(str []byte, pos int) (node, int) {
	openBrackets := 1
	var closeBracket int
	for closeBracket = pos + 1; openBrackets > 0; closeBracket++ {
		switch str[closeBracket] {
		case '[':
			openBrackets++
		case ']':
			openBrackets--
		}
	}
	return parse(str[pos+1 : closeBracket-1]), closeBracket
}

func parseNumber(str []byte, pos int) (node, int) {
	num := 0
	for ; pos < len(str) && str[pos] != ','; pos++ {
		num = num*10 + int(str[pos]-'0')
	}
	return node{val: &num}, pos
}

func process(pairs []pair) int {
	sum := 0
	for i, p := range pairs {
		if p.node1.compare(p.node2) == cmpRight {
			sum += i + 1
		}
	}

	return sum
}

type node struct {
	children []*node
	val      *int
}

func (n node) compare(an node) cmp {
	if n.val != nil && an.val != nil {
		return cmpValues(n.val, an.val)
	}

	getChildren := func(n node) []*node {
		if n.children != nil {
			return n.children
		}
		return []*node{{val: n.val}}
	}
	children1 := getChildren(n)
	children2 := getChildren(an)

	for i := 0; i < len(children1); i++ {
		if i >= len(children2) {
			return cmpWrong
		}
		c := children1[i].compare(*children2[i])
		if c != cmpEqual {
			return c
		}
	}

	if len(children1) == len(children2) {
		return cmpEqual
	}

	return cmpRight
}

func cmpValues(v1, v2 *int) cmp {
	switch {
	case *v1 < *v2:
		return cmpRight
	case *v1 > *v2:
		return cmpWrong
	}
	return cmpEqual
}

type pair struct {
	node1 node
	node2 node
}

type cmp int

const (
	cmpRight = 1
	cmpWrong = -1
	cmpEqual = 0
)
