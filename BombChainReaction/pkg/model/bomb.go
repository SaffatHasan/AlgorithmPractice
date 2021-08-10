package model

import "math"

type Bomb struct {
	X, Y, R int
}

func (b *Bomb) Distance(other Bomb) float64 {
	return math.Sqrt(
		math.Pow(
			float64(b.X-other.X),
			2,
		) + math.Pow(
			float64(b.Y-other.Y),
			2,
		),
	)
}
