package raygui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	AnchorTopLeft  = rl.NewVector2(0, 0)
	AnchorTopRight = rl.NewVector2(1, 0)
	AnchorCenter   = rl.NewVector2(0.5, 0.5)
)

func Pivot(anchor rl.Vector2, r rl.Rectangle) rl.Rectangle {
	return rl.NewRectangle(r.X-r.Width*anchor.X, r.Y-r.Height*anchor.Y, r.Width, r.Height)
}

type canvasLayout struct {
	bounds rl.Rectangle
}

func CanvasLayout(bounds rl.Rectangle) canvasLayout {
	return canvasLayout{
		bounds: bounds,
	}
}

func (cl *canvasLayout) Add(anchor rl.Vector2, pivot rl.Vector2, rect rl.Rectangle) rl.Rectangle {
	anchorp := rl.NewVector2(anchor.X*cl.bounds.Width, anchor.Y*cl.bounds.Height)
	pivotp := rl.NewVector2(pivot.X*rect.Width, pivot.Y*rect.Height)
	return rl.NewRectangle(anchorp.X-pivotp.X+rect.X, anchorp.Y-pivotp.Y+rect.Y, rect.Width, rect.Height)
}

type verticalLayout struct {
	bounds   rl.Rectangle
	spacing  int
	position int
}

func VerticalLayout(bounds rl.Rectangle, spacing int) verticalLayout {
	return verticalLayout{
		bounds:   bounds,
		spacing:  spacing,
		position: 0,
	}
}

func (vl *verticalLayout) Add(wh rl.Vector2) rl.Rectangle {
	r := rl.NewRectangle(vl.bounds.X, vl.bounds.Y+float32(vl.position), wh.X, wh.Y)
	vl.position += int(wh.Y) + vl.spacing
	return r
}
