package aoc

type Ordinary interface {
	byte | int
}

func Max[T Ordinary](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordinary](a, b T) T {
	if a < b {
		return a
	}
	return b
}
