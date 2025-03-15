package rect2

import (
	"iter"

	"github.com/igadmg/raylib-go/raymath/vector2"
)

func (r Rectangle[T]) EachUnitCell() iter.Seq[vector2.Vector[T]] {
	return func(yield func(vector2.Vector[T]) bool) {
		for x := T(0); x < r.size.X; x++ {
			for y := T(0); y < r.size.Y; y++ {
				if !yield(vector2.New[T](r.position.X+x, r.position.X+y)) {
					return
				}
			}
		}
	}
}
