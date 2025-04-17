package main

import (
	"github.com/igadmg/gamemath/vector3"
	rl "github.com/igadmg/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - 3d camera free")

	camera := rl.Camera3D{}
	camera.Position = vector3.NewFloat32(10.0, 10.0, 10.0)
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	cubePosition := vector3.NewFloat32(0.0, 0.0, 0.0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree) // Update camera with free camera mode

		if rl.IsKeyDown(rl.KeyZ) {
			camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 320, 133, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 320, 133, rl.Blue)

		rl.DrawText("Free camera default controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse Wheel Pressed to Pan", 40, 60, 10, rl.DarkGray)
		rl.DrawText("- Alt + Mouse Wheel Pressed to Rotate", 40, 80, 10, rl.DarkGray)
		rl.DrawText("- Alt + Ctrl + Mouse Wheel Pressed for Smooth Zoom", 40, 100, 10, rl.DarkGray)
		rl.DrawText("- Z to zoom to (0, 0, 0)", 40, 120, 10, rl.DarkGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
