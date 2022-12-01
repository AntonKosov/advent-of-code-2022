package aoc

import "strconv"

func StrToInt(str string) int {
	r, err := strconv.Atoi(str)
	Must(err)
	return r
}
