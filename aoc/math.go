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

func Abs[T Ordinary](v T) T {
	if v > 0 {
		return v
	}
	return -v
}
