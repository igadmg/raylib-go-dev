package raygui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	AnchorTopLeft     = rl.NewVector2(0, 0)
	AnchorTopRight    = rl.NewVector2(1, 0)
	AnchorCenter      = rl.NewVector2(0.5, 0.5)
	AnchorBottomLeft  = rl.NewVector2(0, 1)
	AnchorBottomRight = rl.NewVector2(1, 1)
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

func (cl *canvasLayout) Layout(anchor rl.Vector2, pivot rl.Vector2, rect rl.Rectangle) rl.Rectangle {
	anchorp := rl.NewVector2(anchor.X*cl.bounds.Width, anchor.Y*cl.bounds.Height)
	pivotp := rl.NewVector2(pivot.X*rect.Width, pivot.Y*rect.Height)
	return rl.NewRectangle(anchorp.X-pivotp.X+rect.X, anchorp.Y-pivotp.Y+rect.Y, rect.Width, rect.Height)
}

type horizontalLayout struct {
	bounds   rl.Rectangle
	spacing  int
	position int
}

func HorizontalLayout(bounds rl.Rectangle, spacing int) horizontalLayout {
	return horizontalLayout{
		bounds:   bounds,
		spacing:  spacing,
		position: 0,
	}
}

func (hl *horizontalLayout) Layout(wh rl.Vector2) rl.Rectangle {
	r := rl.NewRectangle(hl.bounds.X+float32(hl.position), hl.bounds.Y, wh.X, wh.Y)
	hl.position += int(wh.X) + hl.spacing
	return r
}

func (hl *horizontalLayout) Fill(wh rl.Vector2) rl.Rectangle {
	r := rl.NewRectangle(hl.bounds.X+float32(hl.position), hl.bounds.Y, hl.bounds.Width-float32(hl.position), wh.Y)
	hl.position = int(hl.bounds.Width)
	return r
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

func (vl *verticalLayout) Layout(wh rl.Vector2) rl.Rectangle {
	r := rl.NewRectangle(vl.bounds.X, vl.bounds.Y+float32(vl.position), wh.X, wh.Y)
	vl.position += int(wh.Y) + vl.spacing
	return r
}

func (vl *verticalLayout) Fill(wh rl.Vector2) rl.Rectangle {
	r := rl.NewRectangle(vl.bounds.X, vl.bounds.Y+float32(vl.position), wh.X, vl.bounds.Height-float32(vl.position))
	vl.position = int(vl.bounds.Width)
	return r
}
