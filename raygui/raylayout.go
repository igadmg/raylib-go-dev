package raygui

import (
	"iter"

	rl "github.com/igadmg/raylib-go/raylib"
	rm "github.com/igadmg/raylib-go/raymath"
	"github.com/igadmg/raylib-go/raymath/rect2"
)

type Justyfy int

const (
	JustifyTop Justyfy = iota
	JustifyCenter
	JustifyBottom
	JustifyFill
	JustifyLeft
	JustifyRight
)

// Justify a segment of width v in a segment of width max
// returns new segmend width amd distance from start
//
//	/--------max-------------/
//	/-----v----/
//	JustifyTop, JustifyLeft:
//	/-----v----/                 (v, 0)
//	JustifyCenter:
//	        /-----v----/         (v, (max - v) / 2)
//	JustifyBottom, JustifyRight:
//	              /-----v----/   (v, max - v)
//	JustifyFill:
//	/---------v--------------/   (max, 0)
func (j Justyfy) Justyfy(v, max float32) (nv, dv float32) {
	switch j {
	case JustifyTop:
	case JustifyLeft:
		return v, 0
	case JustifyCenter:
		return v, (max - v) / 2
	case JustifyBottom:
	case JustifyRight:
		return v, max - v
	case JustifyFill:
		return max, 0
	}

	return v, 0
}

func Spread[S ~[]E, E any](s S, a, b rl.Vector2, maxd, mind float32, justify Justyfy) iter.Seq2[rl.Vector2, E] {
	count := len(s)
	if count == 0 {
		return func(yield func(rl.Vector2, E) bool) {}
	} else if count == 1 {
		return func(yield func(rl.Vector2, E) bool) { yield(a, s[0]) }
	}

	return func(yield func(rl.Vector2, E) bool) {
		ab := b.Sub(a)
		ab_length := ab.LengthF()
		direction := ab.Normalized()
		distance := rm.Clamp(ab_length/(float32(count)-1), mind, maxd)

		_, dv := justify.Justyfy(min(distance*(float32(count)-1), ab_length), ab_length)
		position := a.Add(direction.ScaleF(dv))
		delta := direction.ScaleF(distance)
		for _, e := range s {
			if !yield(position, e) {
				return
			}

			position = position.Add(delta)
		}
	}
}

type layout struct {
	Bounds rl.Rectangle
}

func (l *layout) Anchor(xy rl.Vector2) rl.Vector2 {
	return l.Bounds.Size().NormalizeF(xy)
}

type canvasLayout struct {
	layout
}

func CanvasLayout(bounds rl.Rectangle) canvasLayout {
	return canvasLayout{
		layout: layout{Bounds: bounds},
	}
}

func (cl *canvasLayout) Layout(anchor rl.Vector2, pivot rl.Vector2, wh rl.Vector2) rl.Rectangle {
	anchorp := anchor.MultByVector(cl.Bounds.Size())
	pivotp := pivot.MultByVector(wh)
	return rect2.NewFloat32(cl.Bounds.Position().Add(anchorp).Sub(pivotp), wh)
}

type horizontalLayout struct {
	layout
	spacing  int
	position int
}

func HorizontalLayout(bounds rl.Rectangle, spacing int) horizontalLayout {
	return horizontalLayout{
		layout:   layout{Bounds: bounds},
		spacing:  spacing,
		position: 0,
	}
}

func (hl *horizontalLayout) Layout(wh rl.Vector2, justify Justyfy) rl.Rectangle {
	whY, dy := justify.Justyfy(wh.Y(), hl.Bounds.Height())
	if wh.X() < 0 {
		wh = wh.SetX(hl.Bounds.Width() - float32(hl.position) + wh.X() - float32(hl.spacing))
	}
	r := rl.NewRectangle(hl.Bounds.X()+float32(hl.position), hl.Bounds.Y()+dy, wh.X(), whY)
	hl.position += int(wh.X()) + hl.spacing
	return r
}

func (hl *horizontalLayout) Fill(height float32, justify Justyfy) rl.Rectangle {
	whY, dy := justify.Justyfy(height, hl.Bounds.Height())
	r := rl.NewRectangle(hl.Bounds.X()+float32(hl.position), hl.Bounds.Y()+dy, hl.Bounds.Width()-float32(hl.position), whY)
	hl.position = int(hl.Bounds.Width())
	return r
}

func (hl *horizontalLayout) Pie(percent float32) rl.Rectangle {
	r := hl.Bounds.ShrinkXYWH(float32(hl.position), 0, 0, 0)
	r = r.ShrinkXYWH(0, 0, r.Width()*(1-percent), 0).Round()
	hl.position += int(r.Width())
	return r
}

type verticalLayout struct {
	layout
	spacing  int
	position int
}

func VerticalLayout(bounds rl.Rectangle, spacing int) verticalLayout {
	return verticalLayout{
		layout:   layout{Bounds: bounds},
		spacing:  spacing,
		position: 0,
	}
}

func (vl *verticalLayout) Layout(wh rl.Vector2, justify Justyfy) rl.Rectangle {
	whX, dx := justify.Justyfy(wh.X(), vl.Bounds.Width())
	if wh.Y() < 0 {
		wh = wh.SetY(vl.Bounds.Height() - float32(vl.position) + wh.Y() - float32(vl.spacing))
	}
	r := rl.NewRectangle(vl.Bounds.X()+dx, vl.Bounds.Y()+float32(vl.position), whX, wh.Y())
	vl.position += int(wh.Y()) + vl.spacing
	return r
}

func (vl *verticalLayout) Fill(width float32, justify Justyfy) rl.Rectangle {
	whX, dx := justify.Justyfy(width, vl.Bounds.Width())
	r := rl.NewRectangle(vl.Bounds.X()+dx, vl.Bounds.Y()+float32(vl.position), whX, vl.Bounds.Height()-float32(vl.position))
	vl.position = int(vl.Bounds.Height())
	return r
}

func (vl *verticalLayout) Pie(percent float32) rl.Rectangle {
	r := vl.Bounds.ShrinkXYWH(0, float32(vl.position), 0, 0)
	r = r.ShrinkXYWH(0, 0, 0, r.Height()*(1-percent)).Round()
	vl.position += int(r.Height())
	return r
}
