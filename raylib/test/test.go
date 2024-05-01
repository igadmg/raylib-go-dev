package test

import (
	"github.com/EliCDavis/vector"
	. "github.com/EliCDavis/vector/test"
	"github.com/gen2brain/raylib-go/raylib/rect2"
	"github.com/stretchr/testify/assert"
)

func AssertRectangle2InDelta[T vector.Number](t assert.TestingT, expected, actual rect2.Rectangle[T], delta float64) {
	AssertVector2InDelta(t, expected.Position(), actual.Position(), delta)
	AssertVector2InDelta(t, expected.Size(), actual.Size(), delta)
}
