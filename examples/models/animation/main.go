/*******************************************************************************************
*
*   raylib [models] example - Load 3d model with animations and play them
*
*   Example originally created with raylib 2.5, last time updated with raylib 3.5
*
*   Example contributed by Culacant (@culacant) and reviewed by Ramon Santamaria (@raysan5)
*
*   Example licensed under an unmodified zlib/libpng license, which is an OSI-certified,
*   BSD-like license that allows static linking with closed source software
*
*   Copyright (c) 2019-2024 Culacant (@culacant) and Ramon Santamaria (@raysan5)
*
********************************************************************************************
*
*   NOTE: To export a model from blender, make sure it is not posed, the vertices need to be
*         in the same position as they would be in edit mode and the scale of your models is
*         set to 0. Scaling can be done from the export menu.
*
********************************************************************************************/
package main

import (
	"github.com/igadmg/gamemath/vector3"
	rl "github.com/igadmg/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - model animation")

	camera := rl.Camera{}
	camera.Position = vector3.NewFloat32(10.0, 15.0, 10.0)
	camera.Target = vector3.NewFloat32(0.0, 0.0, 0.0)
	camera.Up = vector3.NewFloat32(0.0, 1.0, 0.0)
	camera.Fovy = 75.0
	camera.Projection = rl.CameraPerspective

	model := rl.LoadModel("guy.iqm")
	texture := rl.LoadTexture("guytex.png")
	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texture)

	position := vector3.NewFloat32(0, 0, 0)

	anims := rl.LoadModelAnimations("guyanim.iqm")
	animFrameCount := 0

	rl.DisableCursor()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraOrbital)

		if rl.IsKeyDown(rl.KeySpace) {
			animFrameCount++
			animCurrent := anims[0]
			animFrameNum := animCurrent.FrameCount

			rl.UpdateModelAnimation(model, anims[0], int32(animFrameCount))
			if animFrameCount >= int(animFrameNum) {
				animFrameCount = 0
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		rl.DrawModelEx(model, position, vector3.NewFloat32(1, 0, 0), -90, vector3.NewFloat32(1, 1, 1), rl.White)
		// Draw translation cubes
		for i := int32(0); i < model.BoneCount; i++ {
			pose := anims[0].GetFramePose(animFrameCount, int(i))
			rl.DrawCube(pose.Translation, 0.2, 0.2, 0.2, rl.Red)
		}
		rl.DrawGrid(10, 1)

		rl.EndMode3D()

		rl.DrawText("PRESS SPACE to PLAY MODEL ANIMATION", 10, 10, 20, rl.Black)
		rl.DrawText("(c) Guy IQM 3D model by @culacant", 10, 30, 10, rl.Black)

		rl.EndDrawing()
	}

	rl.UnloadModel(&model)
	rl.UnloadModelAnimations(anims)
	rl.UnloadTexture(&texture)

	rl.CloseWindow()
}
