package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []string {
	lines := aoc.ReadAllInput()
	return lines[:len(lines)-1]
}

func process(snafuNumbers []string) string {
	sum := 0
	for _, sn := range snafuNumbers {
		sum += snafuToDecimal(sn)
	}

	return decimalToSnafu(sum)
}

func snafuToDecimal(number string) int {
	d := 1
	res := 0
	digits := []byte(number)
	for i := len(digits) - 1; i >= 0; i-- {
		switch digit := digits[i]; digit {
		case '2':
			res += d * 2
		case '1':
			res += d
		case '-':
			res -= d
		case '=':
			res -= d * 2
		}
		d *= 5
	}

	return res
}

func decimalToSnafu(number int) string {
	var snafuNumber []byte
	for number > 0 {
		var sd byte
		switch r := number % 5; r {
		case 0:
			sd = '0'
			number /= 5
		case 1:
			sd = '1'
			number /= 5
		case 2:
			sd = '2'
			number /= 5
		case 3:
			sd = '='
			number = (number + 2) / 5
		case 4:
			sd = '-'
			number = (number + 1) / 5
		}
		snafuNumber = append(snafuNumber, sd)
	}

	for i := 0; i < len(snafuNumber)/2; i++ {
		si := len(snafuNumber) - i - 1
		snafuNumber[i], snafuNumber[si] = snafuNumber[si], snafuNumber[i]
	}

	return string(snafuNumber)
}
