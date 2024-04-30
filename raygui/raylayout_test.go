package raygui_test

import (
	"testing"

	"github.com/EliCDavis/vector/test"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestCanvasLayout(t *testing.T) {
	cl := rg.CanvasLayout(rl.NewRectangle(-50, -50, 200, 200))
	{
		r := cl.Layout(rl.AnchorTopLeft, rl.AnchorTopLeft, rl.NewVector2(20, 20))
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -50, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorTopRight, rl.AnchorTopRight, rl.NewVector2(20, 20))
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(130, -50, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorTopLeft, rl.AnchorCenter, rl.NewVector2(20, 20))
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(-60, -60, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorBottomLeft, rl.AnchorBottomLeft, rl.NewVector2(20, 20))
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, 130, 20, 20), 0.00001)
	}
	{
		r := cl.Layout(rl.AnchorBottomRight, rl.AnchorBottomRight, rl.NewVector2(20, 20))
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(130, 130, 20, 20), 0.00001)
	}
}

func TestVerticalLayout(t *testing.T) {
	vl := rg.VerticalLayout(rl.NewRectangle(-50, -50, 200, 200), 0)
	{
		r := vl.Layout(rl.NewVector2(100, 10), rg.JustifyLeft)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -50, 100, 10), 0.00001)
	}
	{
		r := vl.Layout(rl.NewVector2(100, 15), rg.JustifyRight)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(50, -40, 100, 15), 0.00001)
	}
	{
		r := vl.Layout(rl.NewVector2(100, 20), rg.JustifyCenter)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(0, -25, 100, 20), 0.00001)
	}
	{
		r := vl.Layout(rl.NewVector2(100, 25), rg.JustifyFill)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -5, 200, 25), 0.00001)
	}
}

func TestVerticalLayoutSpacing5(t *testing.T) {
	vl := rg.VerticalLayout(rl.NewRectangle(-50, -50, 200, 200), 5)
	{
		r := vl.Layout(rl.NewVector2(100, 10), rg.JustifyLeft)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, -50, 100, 10), 0.00001)
	}
	{
		r := vl.Layout(rl.NewVector2(100, 15), rg.JustifyRight)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(50, -35, 100, 15), 0.00001)
	}
	{
		r := vl.Layout(rl.NewVector2(100, 20), rg.JustifyCenter)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(0, -15, 100, 20), 0.00001)
	}
	{
		r := vl.Layout(rl.NewVector2(100, 25), rg.JustifyFill)
		test.AssertRectangle2InDelta(t, r, rl.NewRectangle(-50, 10, 200, 25), 0.00001)
	}
}
