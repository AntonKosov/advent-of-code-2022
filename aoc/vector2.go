package aoc

type Vector2[T Ordinary] struct {
	X T
	Y T
}

func NewVector2[T Ordinary](x, y T) Vector2[T] {
	return Vector2[T]{X: x, Y: y}
}

func (v Vector2[T]) Add(av Vector2[T]) Vector2[T] {
	return Vector2[T]{X: v.X + av.X, Y: v.Y + av.Y}
}
