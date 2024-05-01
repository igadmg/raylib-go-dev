package test

import (
	"github.com/EliCDavis/vector/test"
	"github.com/EliCDavis/vector/vector2"
	"github.com/EliCDavis/vector/vector3"
	"github.com/stretchr/testify/assert"
)

func AssertRectangle2InDelta[T vector.Number](t assert.TestingT, expected, actual rect2.Rectangle[T], delta float64) {
	AssertVector2InDelta(t, expected.Position(), actual.Position(), delta)
	AssertVector2InDelta(t, expected.Size(), actual.Size(), delta)
}
