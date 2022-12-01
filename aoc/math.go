package aoc

type Ordinary interface {
	int
}

func Max[T Ordinary](a, b T) T {
	if a > b {
		return a
	}
	return b
}
