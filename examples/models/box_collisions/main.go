package main

import (
	"github.com/igadmg/gamemath/vector3"
	rl "github.com/igadmg/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - box collisions")

	camera := rl.Camera{}
	camera.Position = vector3.NewFloat32(0.0, 10.0, 10.0)
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	playerPosition := vector3.NewFloat32(0.0, 1.0, 2.0)
	playerSize := vector3.NewFloat32(1.0, 2.0, 1.0)
	playerColor := rl.Green

	enemyBoxPos := vector3.NewFloat32(-4.0, 1.0, 0.0)
	enemyBoxSize := vector3.NewFloat32(2.0, 2.0, 2.0)

	enemySpherePos := vector3.NewFloat32(4.0, 0.0, 0.0)
	enemySphereSize := float32(1.5)

	collision := false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update

		// Move player
		if rl.IsKeyDown(rl.KeyRight) {
			playerPosition.X += 0.2
		} else if rl.IsKeyDown(rl.KeyLeft) {
			playerPosition.X -= 0.2
		} else if rl.IsKeyDown(rl.KeyDown) {
			playerPosition.Z += 0.2
		} else if rl.IsKeyDown(rl.KeyUp) {
			playerPosition.Z -= 0.2
		}

		collision = false

		// Check collisions player vs enemy-box
		if rl.CheckCollisionBoxes(
			rl.NewBoundingBox(
				vector3.NewFloat32(playerPosition.X-playerSize.X/2, playerPosition.Y-playerSize.Y/2, playerPosition.Z-playerSize.Z/2),
				vector3.NewFloat32(playerPosition.X+playerSize.X/2, playerPosition.Y+playerSize.Y/2, playerPosition.Z+playerSize.Z/2)),
			rl.NewBoundingBox(
				vector3.NewFloat32(enemyBoxPos.X-enemyBoxSize.X/2, enemyBoxPos.Y-enemyBoxSize.Y/2, enemyBoxPos.Z-enemyBoxSize.Z/2),
				vector3.NewFloat32(enemyBoxPos.X+enemyBoxSize.X/2, enemyBoxPos.Y+enemyBoxSize.Y/2, enemyBoxPos.Z+enemyBoxSize.Z/2)),
		) {
			collision = true
		}

		// Check collisions player vs enemy-sphere
		if rl.CheckCollisionBoxSphere(
			rl.NewBoundingBox(
				vector3.NewFloat32(playerPosition.X-playerSize.X/2, playerPosition.Y-playerSize.Y/2, playerPosition.Z-playerSize.Z/2),
				vector3.NewFloat32(playerPosition.X+playerSize.X/2, playerPosition.Y+playerSize.Y/2, playerPosition.Z+playerSize.Z/2)),
			enemySpherePos,
			enemySphereSize,
		) {
			collision = true
		}

		if collision {
			playerColor = rl.Red
		} else {
			playerColor = rl.Green
		}

		// Draw

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		// Draw enemy-box
		rl.DrawCube(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, rl.Gray)
		rl.DrawCubeWires(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, rl.DarkGray)

		// Draw enemy-sphere
		rl.DrawSphere(enemySpherePos, enemySphereSize, rl.Gray)
		rl.DrawSphereWires(enemySpherePos, enemySphereSize, 16, 16, rl.DarkGray)

		// Draw player
		rl.DrawCubeV(playerPosition, playerSize, playerColor)

		rl.DrawGrid(10, 1.0) // Draw a grid

		rl.EndMode3D()

		rl.DrawText("Move player with cursors to collide", 220, 40, 20, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
