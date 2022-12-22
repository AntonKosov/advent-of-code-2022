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

func read() (values map[string]int, expressions map[string]expression) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	values = map[string]int{}
	expressions = map[string]expression{}
	for _, line := range lines {
		variable := line[:4]
		parts := strings.Split(line[6:], " ")
		if len(parts) == 1 {
			values[variable] = aoc.StrToInt(parts[0])
		} else {
			expressions[variable] = expression{
				arg1:      parts[0],
				arg2:      parts[2],
				operation: parts[1],
			}
		}
	}

	return values, expressions
}

func process(values map[string]int, expressions map[string]expression) int {
	var value func(string) int
	var calc func(string) int
	value = func(variable string) int {
		if v, ok := values[variable]; ok {
			return v
		}
		return calc(variable)
	}
	calc = func(variable string) int {
		exp := expressions[variable]
		arg1 := value(exp.arg1)
		arg2 := value(exp.arg2)
		switch exp.operation {
		case "+":
			return arg1 + arg2
		case "-":
			return arg1 - arg2
		case "/":
			if variable == "gzjh" {
				fmt.Println(variable, arg1, arg2, arg1%arg2)
			}
			return arg1 / arg2
		case "*":
			return arg1 * arg2
		}
		panic(fmt.Sprintf("unknown operation: %v", exp.operation))
	}

	return value("root")
}

type expression struct {
	arg1      string
	arg2      string
	operation string
}
