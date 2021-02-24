package raygui

import rl "github.com/gen2brain/raylib-go/raylib"

// Label - Label element, show text
func Label(bounds rl.Rectangle, text string) {
	LabelEx(bounds, text, rl.GetColor(int32(style[LabelTextColor])), rl.NewColor(0, 0, 0, 0), rl.NewColor(0, 0, 0, 0))
}

// LabelEx - Label element extended, configurable colors
func LabelEx(bounds rl.Rectangle, text string, textColor, border, inner rl.Color) {
	// Update control
	textHeight := GetStyle32(GlobalTextFontsize)
	textWidth := rl.MeasureText(text, textHeight)

	b := bounds.ToInt32()
	if b.Width < textWidth {
		b.Width = textWidth + GetStyle32(LabelTextPadding)
	}
	if b.Height < textHeight {
		b.Height = textHeight + GetStyle32(LabelTextPadding)/2
	}

	// Draw control
	DrawBorderedRectangle(b, GetStyle32(LabelBorderWidth), border, inner)
	rl.DrawText(text, b.X+((b.Width/2)-(textWidth/2)), b.Y+((b.Height/2)-(textHeight/2)), textHeight, textColor)
}
