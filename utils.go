package main

func Clamp(v, a, b int) int {
	if v < a {
		return a
	}
	if v > b {
		return b
	}
	return v
}

type Vector2d struct {
	x, y int
}

func (v *Vector2d) Negate() *Vector2d {
	return &Vector2d{-v.x, -v.y}
}

var V = struct {
	Zero     Vector2d
	Identity Vector2d
	Left     Vector2d
	Right    Vector2d
	Up       Vector2d
	Down     Vector2d
	Negate   func(Vector2d) Vector2d
	Sum      func(Vector2d, Vector2d) Vector2d
	Multiply func(Vector2d, int) Vector2d
	Divide   func(Vector2d, int) Vector2d
}{
	Zero:     Vector2d{0, 0},
	Identity: Vector2d{1, 1},
	Left:     Vector2d{-1, 0},
	Right:    Vector2d{1, 0},
	Up:       Vector2d{0, -1},
	Down:     Vector2d{0, 1},
	Negate: func(a Vector2d) Vector2d {
		return Vector2d{
			-a.x,
			-a.y,
		}
	},
	Sum: func(a Vector2d, b Vector2d) Vector2d {
		return Vector2d{
			a.x + b.x,
			a.y + b.y,
		}
	},
	Multiply: func(a Vector2d, b int) Vector2d {
		return Vector2d{
			a.x * b,
			a.y * b,
		}
	},
	Divide: func(a Vector2d, b int) Vector2d {
		return Vector2d{
			a.x / b,
			a.y / b,
		}
	},
}
