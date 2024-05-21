package raygui_test

import (
	"testing"

	. "github.com/igadmg/raylib-go/raymath/test"
	rg "github.com/igadmg/raylib-go/raygui"
	rl "github.com/igadmg/raylib-go/raylib"
)

func TestCanvasLayout(t *testing.T) {
	cl := rg.CanvasLayout(rl.NewRectangle(-50, -50, 200, 200))
	{
		r := cl.Layout(rl.AnchorTopLeft, rl.AnchorTopLeft, vector2.NewFloat32(20, 20))
		AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -50, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorTopRight, rl.AnchorTopRight, vector2.NewFloat32(20, 20))
		AssertRectangle2InDelta(t, r, rl.NewRectangle(130, -50, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorTopLeft, rl.AnchorCenter, vector2.NewFloat32(20, 20))
		AssertRectangle2InDelta(t, r, rl.NewRectangle(-60, -60, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorBottomLeft, rl.AnchorBottomLeft, vector2.NewFloat32(20, 20))
		AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, 130, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorBottomRight, rl.AnchorBottomRight, vector2.NewFloat32(20, 20))
		AssertRectangle2InDelta(t, r, rl.NewRectangle(130, 130, 20, 20), 0.00001)
	}
}

func TestVerticalLayout(t *testing.T) {
	vl := rg.VerticalLayout(rl.NewRectangle(-50, -50, 200, 200), 0)
	{
		r := vl.Layout(vector2.NewFloat32(100, 10), rg.JustifyLeft)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -50, 100, 10), 0.00001)
	}
	{
		r := vl.Layout(vector2.NewFloat32(100, 15), rg.JustifyRight)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(50, -40, 100, 15), 0.00001)
	}
	{
		r := vl.Layout(vector2.NewFloat32(100, 20), rg.JustifyCenter)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(0, -25, 100, 20), 0.00001)
	}
	{
		r := vl.Layout(vector2.NewFloat32(100, 25), rg.JustifyFill)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -5, 200, 25), 0.00001)
	}
}

func TestVerticalLayoutSpacing5(t *testing.T) {
	vl := rg.VerticalLayout(rl.NewRectangle(-50, -50, 200, 200), 5)
	{
		r := vl.Layout(vector2.NewFloat32(100, 10), rg.JustifyLeft)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -50, 100, 10), 0.00001)
	}
	{
		r := vl.Layout(vector2.NewFloat32(100, 15), rg.JustifyRight)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(50, -35, 100, 15), 0.00001)
	}
	{
		r := vl.Layout(vector2.NewFloat32(100, 20), rg.JustifyCenter)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(0, -15, 100, 20), 0.00001)
	}
	{
		r := vl.Layout(vector2.NewFloat32(100, 25), rg.JustifyFill)
		AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, 10, 200, 25), 0.00001)
	}
}
