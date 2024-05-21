package vector4

import rm "github.com/igadmg/raylib-go/raymath"

type Serializable[T rm.SignedNumber] struct {
	X T
	Y T
	Z T
	W T
}

func (m Serializable[T]) Immutable() Vector[T] {
	return Vector[T]{
		x: m.X,
		y: m.Y,
		z: m.Z,
		w: m.W,
	}
}
