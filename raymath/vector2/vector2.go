package vector2

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"

	rm "github.com/igadmg/raylib-go/raymath"
)

type Vector[T rm.SignedNumber] struct {
	x T
	y T
}

type (
	Float64 = Vector[float64]
	Float32 = Vector[float32]
	Int     = Vector[int]
	Int64   = Vector[int64]
	Int32   = Vector[int32]
	Int16   = Vector[int16]
	Int8    = Vector[int8]
)

func New[T rm.SignedNumber](x T, y T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
	}
}

func NewT[T rm.SignedNumber, XT, YT rm.Number](x XT, y YT) Vector[T] {
	return Vector[T]{
		x: T(x),
		y: T(y),
	}
}

func NewFloat64[XT, YT rm.Number](x XT, y YT) Float64 {
	return NewT[float64](x, y)
}

func NewFloat32[XT, YT rm.Number](x XT, y YT) Float32 {
	return NewT[float32](x, y)
}

func NewInt[XT, YT rm.Number](x XT, y YT) Int {
	return NewT[int](x, y)
}

func NewInt64[XT, YT rm.Number](x XT, y YT) Int64 {
	return NewT[int64](x, y)
}

func NewInt32[XT, YT rm.Number](x XT, y YT) Int32 {
	return NewT[int32](x, y)
}

func NewInt16[XT, YT rm.Number](x XT, y YT) Int16 {
	return NewT[int16](x, y)
}

func NewInt8[XT, YT rm.Number](x XT, y YT) Int8 {
	return NewT[int8](x, y)
}

// Fill creates a vector where each component is equal to v
func Fill[T rm.SignedNumber](v T) Vector[T] {
	return Vector[T]{
		x: v,
		y: v,
	}
}

func Zero[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: 0,
	}
}

func Up[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: 1,
	}
}

func Down[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: -1,
	}
}

func Left[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		x: -1,
		y: 0,
	}
}

func Right[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		x: 1,
		y: 0,
	}
}

func One[T rm.SignedNumber]() Vector[T] {
	return Vector[T]{
		x: 1,
		y: 1,
	}
}

// Lerp linearly interpolates between a and b by t
func Lerp[T rm.SignedNumber](t float64, a, b Vector[T]) Vector[T] {
	return Vector[T]{
		x: rm.Lerp(t, a.x, b.x),
		y: rm.Lerp(t, a.y, b.y),
	}
}

func Min[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	return New(
		min(a.x, b.x),
		min(a.y, b.y),
	)
}

func Max[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	return New(
		max(a.x, b.x),
		max(a.y, b.y),
	)
}

func MaxX[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.x, b.x)
}

func MaxY[T rm.SignedNumber](a, b Vector[T]) T {
	return max(a.y, b.y)
}

func MinX[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.x, b.x)
}

func MinY[T rm.SignedNumber](a, b Vector[T]) T {
	return min(a.y, b.y)
}

func Less[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.x < b.x && a.y < b.y
}

func LessEq[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.x <= b.x && a.y <= b.y
}

func Greater[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.x > b.x && a.y > b.y
}

func GreaterEq[T rm.SignedNumber](a, b Vector[T]) bool {
	return a.x >= b.x && a.y >= b.y
}

func Midpoint[T rm.SignedNumber](a, b Vector[T]) Vector[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Vector[T]{
		x: T(float64(a.x+b.x) * 0.5),
		y: T(float64(a.y+b.y) * 0.5),
	}
}

func Index(xy Int, i int) Int {
	return Int{
		i % xy.X(),
		i / xy.X(),
	}
}

// Builds a vector from the data found from the passed in array to the best of
// it's ability. If the length of the array is smaller than the vector itself,
// only those values will be used to build the vector, and the remaining vector
// components will remain the default value of the vector's data type (some
// version of 0).
func FromArray[T rm.SignedNumber](data []T) Vector[T] {
	v := Vector[T]{}

	if len(data) > 0 {
		v.x = data[0]
	}

	if len(data) > 1 {
		v.y = data[1]
	}

	return v
}

func Rand(r *rand.Rand) Vector[float64] {
	return Vector[float64]{
		x: r.Float64(),
		y: r.Float64(),
	}
}

func (v Vector[T]) MinComponent() T {
	return min(v.x, v.y)
}

func (v Vector[T]) MaxComponent() T {
	return max(v.x, v.y)
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: float64(v.x),
		Y: float64(v.y),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}{
		X: 0,
		Y: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.x = T(aux.X)
	v.y = T(aux.Y)
	return nil
}

func (v Vector[T]) Format(format string) string {
	return fmt.Sprintf(format, v.x, v.y)
}

// Sqrt applies the math.Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		rm.Sqrt(v.x),
		rm.Sqrt(v.y),
	)
}

func (v Vector[T]) Clamp(vmin, vmax T) Vector[T] {
	return Vector[T]{
		x: rm.Clamp(v.x, vmin, vmax),
		y: rm.Clamp(v.y, vmin, vmax),
	}
}

func (v Vector[T]) ClampV(vmin, vmax Vector[T]) Vector[T] {
	return Vector[T]{
		x: rm.Clamp(v.x, vmin.x, vmax.x),
		y: rm.Clamp(v.y, vmin.y, vmax.y),
	}
}

func (v Vector[T]) Clamp0V(vmax Vector[T]) Vector[T] {
	return Vector[T]{
		x: rm.Clamp(v.x, 0, vmax.x),
		y: rm.Clamp(v.y, 0, vmax.y),
	}
}

func (v Vector[T]) ToNpot() Vector[T] {
	return Vector[T]{
		x: rm.Npot(v.x),
		y: rm.Npot(v.y),
	}
}

func To[T, OT rm.SignedNumber](v Vector[OT]) Vector[T] {
	return Vector[T]{
		x: T(v.x),
		y: T(v.y),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return To[float64](v)
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return To[float32](v)
}

func (v Vector[T]) ToInt() Vector[int] {
	return To[int](v)
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return To[int64](v)
}

func (v Vector[T]) ToInt32() Vector[int32] {
	return To[int32](v)
}

func (v Vector[T]) ToInt16() Vector[int16] {
	return To[int16](v)
}

func (v Vector[T]) ToInt8() Vector[int8] {
	return To[int8](v)
}

func (v Vector[T]) X() T {
	return v.x
}

// SetX changes the x component of the vector
func (v Vector[T]) SetX(X T) Vector[T] {
	return Vector[T]{
		x: X,
		y: v.y,
	}
}

// AddX adds to the x component of the vector
func (v Vector[T]) AddX(dX T) Vector[T] {
	return Vector[T]{
		x: v.x + dX,
		y: v.y,
	}
}

func (v Vector[T]) Y() T {
	return v.y
}

// SetY changes the y component of the vector
func (v Vector[T]) SetY(Y T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: Y,
	}
}

// AddY adds to the y component of the vector
func (v Vector[T]) AddY(dY T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y + dY,
	}
}

func (v Vector[T]) YX() Vector[T] {
	return Vector[T]{
		x: v.y,
		y: v.x,
	}
}

// Angle return angle in radians between vector and other vector [float64]
func (v Vector[T]) Angle(other Vector[T]) float64 {
	denominator := rm.Sqrt(float64(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return rm.Acos(rm.Clamp(float64(v.Dot(other))/denominator, -1., 1.))
}

// AngleF return angle in radians between vector and other vector [float32]
func (v Vector[T]) AngleF(other Vector[T]) float32 {
	denominator := rm.Sqrt(float32(v.LengthSquared() * other.LengthSquared()))
	if denominator < 1e-15 {
		return 0.
	}
	return rm.Acos(rm.Clamp(float32(v.Dot(other))/denominator, -1., 1.))
}

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Vector[T]) Midpoint(o Vector[T]) Vector[T] {
	return o.Add(v).Scale(0.5)
}

// Dot return dot product between vector and other vector
func (v Vector[T]) Dot(other Vector[T]) T {
	return v.x*other.x + v.y*other.y
}

// Perpendicular creates a vector perpendicular to the one passed in with the
// same magnitude
func (v Vector[T]) Perpendicular() Vector[T] {
	return Vector[T]{
		x: v.y,
		y: -v.x,
	}
}

// Add returns a vector that is the result of two vectors added together
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector[T]) AddXY(x, y T) Vector[T] {
	return Vector[T]{
		x: v.x + x,
		y: v.y + y,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v Vector[T]) SubXY(x, y T) Vector[T] {
	return Vector[T]{
		x: v.x - x,
		y: v.y - y,
	}
}

func (v Vector[T]) ReciprocalF() Vector[float32] {
	return Vector[float32]{
		x: 1.0 / float32(v.x),
		y: 1.0 / float32(v.y),
	}
}

func (v Vector[T]) Reciprocal() Vector[float64] {
	return Vector[float64]{
		x: 1.0 / float64(v.x),
		y: 1.0 / float64(v.y),
	}
}

func (v Vector[T]) Product() T {
	return v.x * v.y
}

func (v Vector[T]) LengthSquared() T {
	return v.x*v.x + v.y*v.y
}

func (v Vector[T]) Length() float64 {
	return math.Sqrt((float64)(v.LengthSquared()))
}

func (v Vector[T]) LengthF() float32 {
	return rm.Sqrt((float32)(v.LengthSquared()))
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

func (v Vector[T]) NormalizeF(a Vector[T]) Float32 {
	return Float32{
		x: rm.NormalizeF(a.x, 0, v.x),
		y: rm.NormalizeF(a.y, 0, v.y),
	}
}

func (v Vector[T]) Normalize(a Vector[T]) Float64 {
	return Float64{
		x: rm.Normalize(a.x, 0, v.x),
		y: rm.Normalize(a.y, 0, v.y),
	}
}

func (v Vector[T]) Negated() Vector[T] {
	return Vector[T]{
		x: -v.x,
		y: -v.y,
	}
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * t),
		y: T(float64(v.y) * t),
	}
}

func (v Vector[T]) ScaleF(t float32) Vector[T] {
	return Vector[T]{
		x: T(float32(v.x) * t),
		y: T(float32(v.y) * t),
	}
}

func (v Vector[T]) ScaleByVector(o Float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * o.x),
		y: T(float64(v.y) * o.y),
	}
}

func (v Vector[T]) ScaleByVectorF(o Float32) Vector[T] {
	return Vector[T]{
		x: T(float32(v.x) * o.x),
		y: T(float32(v.y) * o.y),
	}
}

func (v Vector[T]) ScaleByVectorI(o Int) Vector[T] {
	return Vector[T]{
		x: v.x * T(o.x),
		y: v.y * T(o.y),
	}
}

func (v Vector[T]) ScaleByXY(x, y float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * x),
		y: T(float64(v.y) * y),
	}
}

func (v Vector[T]) ScaleByXYF(x, y float32) Vector[T] {
	return Vector[T]{
		x: T(float32(v.x) * x),
		y: T(float32(v.y) * y),
	}
}

func (v Vector[T]) ScaleByXYI(x, y int) Vector[T] {
	return Vector[T]{
		x: v.x * T(x),
		y: v.y * T(y),
	}
}

func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x * o.x,
		y: v.y * o.y,
	}
}

func (v Vector[T]) DivByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x / o.x,
		y: v.y / o.y,
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return v.Scale(1.0 / t)
}

func (v Vector[T]) Project(normal Vector[T]) Vector[T] {
	vdn := float64(v.Dot(normal))
	ndn := float64(normal.Dot(normal))
	mag := vdn / ndn
	return normal.Scale(mag)
}

func (v Vector[T]) Reject(normal Vector[T]) Vector[T] {
	return v.Sub(v.Project(normal))
}

func (v Vector[T]) Reflect(normal Vector[T]) Vector[T] {
	return v.Sub(normal.Scale(2. * float64(v.Dot(normal))))
}

func (v Vector[T]) DistanceSquared(other Vector[T]) T {
	xDist := other.x - v.x
	yDist := other.y - v.y
	return (xDist * xDist) + (yDist * yDist)
}

// Distance is the euclidean distance between two points
func (v Vector[T]) Distance(other Vector[T]) float64 {
	return math.Sqrt((float64)(v.DistanceSquared(other)))
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return Vector[T]{
		x: rm.Round(v.x),
		y: rm.Round(v.y),
	}
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(rm.Round(v.x)),
		int(rm.Round(v.y)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return Vector[T]{
		x: rm.Ceil(v.x),
		y: rm.Ceil(v.y),
	}
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(rm.Ceil(v.x)),
		int(rm.Ceil(v.y)),
	)
}

func (v Vector[T]) Floor() Vector[T] {
	return Vector[T]{
		x: rm.Floor(v.x),
		y: rm.Floor(v.y),
	}
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(rm.Floor(v.x)),
		int(rm.Floor(v.y)),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return Vector[T]{
		x: rm.Abs(v.x),
		y: rm.Abs(v.y),
	}
}

func (v Vector[T]) NearZero() bool {
	return rm.NearZero(v.x) && rm.NearZero(v.y)
}

func (v Vector[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.x)) {
		return true
	}

	if math.IsNaN(float64(v.y)) {
		return true
	}

	return false
}

func (v Vector[T]) Flip() Vector[T] {
	return Vector[T]{
		x: v.x * -1,
		y: v.y * -1,
	}
}

func (v Vector[T]) FlipX() Vector[T] {
	return Vector[T]{
		x: v.x * -1,
		y: v.y,
	}
}

func (v Vector[T]) FlipY() Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y * -1,
	}
}

func (v Vector[T]) Pivot(anchor Vector[T], wh Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x - wh.x*anchor.x,
		y: v.y - wh.y*anchor.y,
	}
}

// Log returns the natural logarithm for each component
func (v Vector[T]) Log() Vector[T] {
	return Vector[T]{
		x: T(math.Log(float64(v.x))),
		y: T(math.Log(float64(v.y))),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Vector[T]) Log10() Vector[T] {
	return Vector[T]{
		x: T(math.Log10(float64(v.x))),
		y: T(math.Log10(float64(v.y))),
	}
}

// Log2 returns the binary logarithm for each component
func (v Vector[T]) Log2() Vector[T] {
	return Vector[T]{
		x: T(math.Log2(float64(v.x))),
		y: T(math.Log2(float64(v.y))),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Vector[T]) Exp2() Vector[T] {
	return Vector[T]{
		x: T(math.Exp2(float64(v.x))),
		y: T(math.Exp2(float64(v.y))),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Vector[T]) Exp() Vector[T] {
	return Vector[T]{
		x: T(math.Exp(float64(v.x))),
		y: T(math.Exp(float64(v.y))),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Vector[T]) Expm1() Vector[T] {
	return Vector[T]{
		x: T(math.Expm1(float64(v.x))),
		y: T(math.Expm1(float64(v.y))),
	}
}

func (v Vector[T]) Values() (T, T) {
	return v.x, v.y
}
