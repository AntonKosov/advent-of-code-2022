package aoc

func LastItem[T any](slice []T) T {
	return slice[len(slice)-1]
}

func RemoveLastItem[T any](slice *[]T) T {
	v := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]

	return v
}
