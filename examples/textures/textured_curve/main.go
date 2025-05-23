package main

import (
	"fmt"
	"math"

	"github.com/igadmg/gamemath/vector2"
	rl "github.com/igadmg/raylib-go/raylib"
)

var (
	texRoad       rl.Texture2D
	showCurve     = true
	curveW        = float32(50)
	curveSegments = 24

	curveStartPos, curveStartPosTangent, curveEndPos, curveEndPosTangent vector2.Float32

	curveSelectedPoint *vector2.Float32

	screenW = int32(800)
	screenH = int32(450)
)

func main() {

	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagMsaa4xHint)

	rl.InitWindow(screenW, screenH, "raylib [textures] example - textured curve")

	texRoad = rl.LoadTexture("road.png")
	rl.SetTextureFilter(texRoad, rl.TextureFilterMode(rl.FilterBilinear))

	curveStartPos = vector2.NewFloat32(80, 100)
	curveStartPosTangent = vector2.NewFloat32(100, 300)

	curveEndPos = vector2.NewFloat32(700, 350)
	curveEndPosTangent = vector2.NewFloat32(600, 100)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		upCurve()
		upOptions()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		drawTexturedCurve()
		drawCurve()

		rl.DrawText("Drag points to move curve, press SPACE to show/hide base curve", 10, 10, 10, rl.Black)
		rl.DrawText("Curve width: "+fmt.Sprintf("%.0f", curveW)+" use UP/DOWN arrows to adjust", 10, 30, 10, rl.Black)
		rl.DrawText("Curve segments: "+fmt.Sprint(curveSegments)+" use RIGHT/LEFT arrows to adjust", 10, 50, 10, rl.Black)

		rl.EndDrawing()
	}

	rl.UnloadTexture(&texRoad)

	rl.CloseWindow()
}

func upCurve() {

	if !rl.IsMouseButtonDown(rl.MouseLeftButton) {
		curveSelectedPoint = &vector2.Float32{}
	}

	*curveSelectedPoint = curveSelectedPoint.Add(rl.GetMouseDelta())

	mouse := rl.GetMousePosition()

	if rl.CheckCollisionPointCircle(mouse, curveStartPos, 6) {
		curveSelectedPoint = &curveStartPos
	} else if rl.CheckCollisionPointCircle(mouse, curveStartPosTangent, 6) {
		curveSelectedPoint = &curveStartPosTangent
	} else if rl.CheckCollisionPointCircle(mouse, curveEndPos, 6) {
		curveSelectedPoint = &curveEndPos
	} else if rl.CheckCollisionPointCircle(mouse, curveEndPosTangent, 6) {
		curveSelectedPoint = &curveEndPosTangent
	}

}
func upOptions() {

	if rl.IsKeyPressed(rl.KeySpace) {
		showCurve = !showCurve
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		curveW += 2
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		curveW -= 2
	}
	if curveW < 2 {
		curveW = 2
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		curveSegments -= 2
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		curveSegments += 2
	}
	if curveSegments < 2 {
		curveSegments = 2
	}

}
func drawTexturedCurve() {

	step := float32(1) / float32(curveSegments)
	previous := curveStartPos
	previousTangent := vector2.Zero[float32]()
	previousV := float32(0)
	tangentSet := false
	current := vector2.Zero[float32]()
	t := float32(0)

	for i := 0; i < curveSegments; i++ {
		t = step * float32(i)
		a := float32(math.Pow(1-float64(t), 3))
		b := 3 * float32(math.Pow(1-float64(t), 2)) * t
		c := 3 * (1 - t) * float32(math.Pow(float64(t), 2))
		d := float32(math.Pow(float64(t), 3))

		current.Y = a*curveStartPos.Y + b*curveStartPosTangent.Y + c*curveEndPosTangent.Y + d*curveEndPos.Y
		current.X = a*curveStartPos.X + b*curveStartPosTangent.X + c*curveEndPosTangent.X + d*curveEndPos.X

		delta := vector2.NewFloat32(current.X-previous.X, current.Y-previous.Y)
		normal := vector2.NewFloat32(-delta.Y, delta.X).Normalized()
		v := previousV + delta.LengthF()

		if !tangentSet {
			previousTangent = normal
			tangentSet = true
		}

		prevPosNormal := previous.Add(previousTangent.ScaleF(curveW))
		prevNegNormal := previous.Add(previousTangent.ScaleF(-curveW))

		currentPosNormal := current.Add(normal.ScaleF(curveW))
		currentNegNormal := current.Add(normal.ScaleF(-curveW))

		rl.SetTexture(texRoad.ID)
		rl.Begin(rl.Quads)

		rl.Color4ub(255, 255, 255, 255)
		rl.Normal3f(0, 0, 1)

		rl.TexCoord2f(0, previousV)
		rl.Vertex2f(prevNegNormal.X, prevNegNormal.Y)

		rl.TexCoord2f(1, previousV)
		rl.Vertex2f(prevPosNormal.X, prevPosNormal.Y)

		rl.TexCoord2f(1, v)
		rl.Vertex2f(currentPosNormal.X, currentPosNormal.Y)

		rl.TexCoord2f(0, v)
		rl.Vertex2f(currentNegNormal.X, currentNegNormal.Y)

		rl.End()

		previous = current
		previousTangent = normal
		previousV = v

	}

}
func drawCurve() {

	if showCurve {
		rl.DrawSplineSegmentBezierCubic(curveStartPos, curveEndPos, curveStartPosTangent, curveEndPosTangent, 2, rl.Blue)
	}
	rl.DrawLineV(curveStartPos, curveStartPosTangent, rl.SkyBlue)
	rl.DrawLineV(curveEndPos, curveEndPosTangent, rl.Purple)
	mouse := rl.GetMousePosition()

	if rl.CheckCollisionPointCircle(mouse, curveStartPos, 6) {
		rl.DrawCircleV(curveStartPos, 7, rl.Yellow)
	}
	rl.DrawCircleV(curveStartPos, 5, rl.Red)

	if rl.CheckCollisionPointCircle(mouse, curveStartPosTangent, 6) {
		rl.DrawCircleV(curveStartPosTangent, 7, rl.Yellow)
	}
	rl.DrawCircleV(curveStartPosTangent, 5, rl.Maroon)

	if rl.CheckCollisionPointCircle(mouse, curveEndPos, 6) {
		rl.DrawCircleV(curveEndPos, 7, rl.Yellow)
	}
	rl.DrawCircleV(curveEndPosTangent, 5, rl.Green)

	if rl.CheckCollisionPointCircle(mouse, curveEndPosTangent, 6) {
		rl.DrawCircleV(curveEndPosTangent, 7, rl.Yellow)
	}
	rl.DrawCircleV(curveEndPosTangent, 5, rl.DarkGreen)

}
