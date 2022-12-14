package main

import (
	"fmt"
	"sort"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	nodes := read()
	r := process(nodes)
	fmt.Printf("Answer: %v\n", r)
}

func read() []node {
	lines := aoc.ReadAllInput()

	var nodes []node
	for _, line := range lines {
		if line == "" {
			continue
		}

		line = line[1 : len(line)-1]
		nodes = append(nodes, parse([]byte(line)))
	}

	divNode1, divNode2 := parse([]byte("[2]")), parse([]byte("[6]"))
	divNode1.divider = true
	divNode2.divider = true
	nodes = append(nodes, divNode1, divNode2)

	return nodes
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

func process(nodes []node) int {
	mul := 1
	sort.Slice(nodes, func(i, j int) bool {
		n1, n2 := nodes[i], nodes[j]
		return n1.compare(n2) == cmpRight
	})

	for i, n := range nodes {
		if n.divider {
			mul *= i + 1
		}
	}

	return mul
}

type node struct {
	children []*node
	val      *int
	divider  bool
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

type cmp int

const (
	cmpRight = 1
	cmpWrong = -1
	cmpEqual = 0
)
