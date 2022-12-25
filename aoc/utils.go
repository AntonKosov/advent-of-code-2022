package aoc

import "regexp"

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ParseInts(str string) []int {
	r := regexp.MustCompile(`[\d]+`)
	matches := r.FindAllString(str, -1)

	res := make([]int, len(matches))
	for i, m := range matches {
		res[i] = StrToInt(m)
	}

	return res
}
