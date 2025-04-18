package vector3

import (
	"math"

	rm "github.com/igadmg/raylib-go/raymath"
)

type Array[T rm.SignedNumber] []Vector[T]

type (
	Float64Array = Array[float64]
	Float32Array = Array[float32]
	IntArray     = Array[int]
	Int64Array   = Array[int64]
)

func (v3a Array[T]) Add(other Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = Vector[T]{
			v.X + other.X,
			v.Y + other.Y,
			v.Z + other.Z,
		}
	}

	return
}

func (v3a Array[T]) AddInplace(other Vector[T]) Array[T] {
	for i, v := range v3a {
		v3a[i] = Vector[T]{
			v.X + other.X,
			v.Y + other.Y,
			v.Z + other.Z,
		}
	}
	return v3a
}

func (v3a Array[T]) Sub(other Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = Vector[T]{
			v.X - other.X,
			v.Y - other.Y,
			v.Z - other.Z,
		}
	}

	return
}

func (v3a Array[T]) SubInplace(other Vector[T]) Array[T] {
	for i, v := range v3a {
		v3a[i] = Vector[T]{
			v.X - other.X,
			v.Y - other.Y,
			v.Z - other.Z,
		}
	}
	return v3a
}

func (v3a Array[T]) Distance() (total float64) {
	if len(v3a) < 2 {
		return
	}
	for i := 1; i < len(v3a); i++ {
		total += v3a[i].Distance(v3a[i-1])
	}
	return
}

func (v3a Array[T]) Scale(t float64) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = Vector[T]{
			X: T(float64(v.X) * t),
			Y: T(float64(v.Y) * t),
			Z: T(float64(v.Z) * t),
		}
	}

	return
}

func (v3a Array[T]) ScaleInplace(t float64) Array[T] {
	for i, v := range v3a {
		v3a[i] = Vector[T]{
			X: T(float64(v.X) * t),
			Y: T(float64(v.Y) * t),
			Z: T(float64(v.Z) * t),
		}
	}
	return v3a
}

func (v3a Array[T]) DivByConstant(t float64) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = v.DivByConstant(t)
	}

	return
}

func (v3a Array[T]) Normalized() (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = v.Normalized()
	}

	return
}

func (v3a Array[T]) ContainsNaN() bool {
	for _, v := range v3a {
		if v.ContainsNaN() {
			return true
		}
	}
	return false
}

func (v3a Array[T]) MaxLength() float64 {
	max := 0.

	for _, v := range v3a {
		max = math.Max(max, v.Length())
	}

	return max
}

func (v3a Array[T]) Sum() (sum Vector[T]) {
	for _, v := range v3a {
		sum = sum.Add(v)
	}
	return
}

func (v3a Array[T]) Modify(f func(Vector[T]) Vector[T]) (out Array[T]) {
	out = make(Array[T], len(v3a))

	for i, v := range v3a {
		out[i] = f(v)
	}

	return
}

// Average sums all vector3's components together and divides each
// component by the number of values added
func (v3a Array[T]) Average(vectors []Vector[T]) Vector[float64] {
	xTotal := 0.
	yTotal := 0.
	zTotal := 0.

	for _, v := range v3a {
		xTotal += float64(v.X)
		yTotal += float64(v.Y)
		zTotal += float64(v.Z)
	}

	return New(xTotal, yTotal, zTotal).DivByConstant(float64(len(v3a)))
}

// Bounds returns the min and max points of an AABB encompassing
func (v3a Array[T]) Bounds() (Vector[T], Vector[T]) {
	vmin := New(math.Inf(1), math.Inf(1), math.Inf(1))
	vmax := New(math.Inf(-1), math.Inf(-1), math.Inf(-1))

	for _, v := range v3a {
		vmin = New(
			min(float64(v.X), vmin.X),
			min(float64(v.Y), vmin.Y),
			min(float64(v.Z), vmin.Z),
		)

		vmax = New(
			max(float64(v.X), vmax.X),
			max(float64(v.Y), vmax.Y),
			max(float64(v.Z), vmax.Z),
		)
	}

	return New(T(vmin.X), T(vmin.Y), T(vmin.Z)), New(T(vmax.X), T(vmax.Y), T(vmax.Z))
}

// StandardDeviation calculates the population standard deviation on each
// component of the vector
func (v3a Array[T]) StandardDeviation() (mean, deviation Vector[float64]) {
	mean = v3a.Average(v3a)

	xTotal, yTotal, zTotal := 0., 0., 0.
	for _, v := range v3a {
		diff := v.ToFloat64().Sub(mean)
		xTotal += (diff.X * diff.X)
		yTotal += (diff.Y * diff.Y)
		zTotal += (diff.Z * diff.Z)
	}

	deviation = New(
		math.Sqrt(xTotal/float64(len(v3a))),
		math.Sqrt(yTotal/float64(len(v3a))),
		math.Sqrt(zTotal/float64(len(v3a))),
	)
	return
}
