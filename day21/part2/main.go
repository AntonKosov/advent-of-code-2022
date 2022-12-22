package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

const me = "humn"
const root = "root"

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (values map[string]int, expressions map[string]expression) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	values = map[string]int{}
	expressions = map[string]expression{}
	for _, line := range lines {
		variable := line[:4]
		if variable == me {
			values[me] = 0
			continue
		}
		parts := strings.Split(line[6:], " ")
		if len(parts) == 1 {
			values[variable] = aoc.StrToInt(parts[0])
		} else {
			op := parts[1]
			if variable == root {
				op = "="
			}
			expressions[variable] = expression{
				args:      [2]string{parts[0], parts[2]},
				operation: operation(op),
			}
		}
	}

	return values, expressions
}

func process(values map[string]int, expressions map[string]expression) int {
	indexInRoot, ok := findMe(root, expressions)
	if !ok {
		panic("not found")
	}
	rootExp := expressions[root]
	requiredValue := value(rootExp.args[1-indexInRoot], values, expressions)

	return findMyValue(rootExp.args[indexInRoot], requiredValue, values, expressions)
}

func findMyValue(variable string, requiredResult int, values map[string]int, expressions map[string]expression) int {
	if variable == me {
		return requiredResult
	}
	myIndex, ok := findMe(variable, expressions)
	if !ok {
		panic("not found")
	}
	exp := expressions[variable]
	anotherValue := value(exp.args[1-myIndex], values, expressions)
	myArg := exp.args[myIndex]
	switch exp.operation {
	case sum:
		return findMyValue(myArg, requiredResult-anotherValue, values, expressions)
	case mul:
		return findMyValue(myArg, requiredResult/anotherValue, values, expressions)
	case sub:
		if myIndex == 0 {
			return findMyValue(myArg, requiredResult+anotherValue, values, expressions)
		} else {
			return findMyValue(myArg, anotherValue-requiredResult, values, expressions)
		}
	case div:
		if myIndex == 0 {
			return findMyValue(myArg, requiredResult*anotherValue, values, expressions)
		} else {
			return findMyValue(myArg, anotherValue/requiredResult, values, expressions)
		}
	default:
		panic(fmt.Sprintf("unknown operation: %v", exp.operation))
	}
}

func value(variable string, values map[string]int, expressions map[string]expression) int {
	if v, ok := values[variable]; ok {
		return v
	}
	v := calc(variable, values, expressions)
	values[variable] = v
	return v
}

func calc(variable string, values map[string]int, expressions map[string]expression) int {
	exp := expressions[variable]
	arg0 := value(exp.args[0], values, expressions)
	arg1 := value(exp.args[1], values, expressions)
	switch exp.operation {
	case sum:
		return arg0 + arg1
	case sub:
		return arg0 - arg1
	case div:
		return arg0 / arg1
	case mul:
		return arg0 * arg1
	default:
		panic(fmt.Sprintf("unknown operation: %v", exp.operation))
	}
}

func findMe(variable string, expressions map[string]expression) (int, bool) {
	if _, ok := expressions[variable]; !ok {
		return 0, false
	}

	exp := expressions[variable]

	for i, arg := range exp.args {
		if arg == me {
			return i, true
		}
		if _, ok := findMe(arg, expressions); ok {
			return i, true
		}
	}

	return 0, false
}

type expression struct {
	args      [2]string
	operation operation
}

type operation string

const (
	sum operation = "+"
	sub operation = "-"
	mul operation = "*"
	div operation = "/"
)
