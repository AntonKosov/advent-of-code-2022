package aoc

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) Add(av Vector2) Vector2 {
	return Vector2{X: v.X + av.X, Y: v.Y + av.Y}
}

func (v Vector2) Sub(av Vector2) Vector2 {
	return Vector2{X: v.X - av.X, Y: v.Y - av.Y}
}

func (v Vector2) Len() int {
	return Max(Abs(v.X), Abs(v.Y))
}

func (v Vector2) Norm() Vector2 {
	if v.X != 0 {
		v.X = v.X / Abs(v.X)
	}

	if v.Y != 0 {
		v.Y = v.Y / Abs(v.Y)
	}

	return v
}

func (v Vector2) ManhattanDst() int {
	return Abs(v.X) + Abs(v.Y)
}
