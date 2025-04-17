package main

import (
	"github.com/igadmg/gamemath/vector2"
	"github.com/igadmg/gamemath/vector3"
	"github.com/igadmg/goex/image/colorex"
	rl "github.com/igadmg/raylib-go/raylib"
)

const (
	maxColumns = 20
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - 3d camera first person")

	camera := rl.Camera3D{}
	camera.Position = vector3.NewFloat32(4.0, 2.0, 4.0)
	camera.Target = vector3.NewFloat32(0.0, 1.8, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	// Generates some random columns
	heights := make([]float32, maxColumns)
	positions := make([]vector3.Float32, maxColumns)
	colors := make([]colorex.RGBA, maxColumns)

	for i := 0; i < maxColumns; i++ {
		heights[i] = float32(rl.GetRandomValue(1, 12))
		positions[i] = vector3.NewFloat32(float32(rl.GetRandomValue(-15, 15)), heights[i]/2, float32(rl.GetRandomValue(-15, 15)))
		colors[i] = rl.NewColor(uint8(rl.GetRandomValue(20, 255)), uint8(rl.GetRandomValue(10, 55)), 30, 255)
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson) // Update camera with first person mode

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawPlane(vector3.NewFloat32(0.0, 0.0, 0.0), vector2.NewFloat32(32.0, 32.0), rl.LightGray) // Draw ground
		rl.DrawCube(vector3.NewFloat32(-16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Blue)                     // Draw a blue wall
		rl.DrawCube(vector3.NewFloat32(16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Lime)                      // Draw a green wall
		rl.DrawCube(vector3.NewFloat32(0.0, 2.5, 16.0), 32.0, 5.0, 1.0, rl.Gold)                      // Draw a yellow wall

		// Draw some cubes around
		for i := 0; i < maxColumns; i++ {
			rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
			rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Maroon)
		}

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)

		rl.DrawText("First person camera default controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Move with keys: W, A, S, D", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse move to look around", 40, 60, 10, rl.DarkGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
