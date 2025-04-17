package main

import (
	"github.com/igadmg/goex/image/colorex"
	rl "github.com/igadmg/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [textures] example - texture from raw data")

	// NOTE: Textures MUST be loaded after Window initialization (OpenGL context is required)

	// Load RAW image data (384x512, 32bit RGBA, no file header)
	fudesumiRaw := rl.LoadImageRaw("texture_formats/fudesumi.raw", 384, 512, rl.UncompressedR8g8b8a8, 0)
	fudesumi := rl.LoadTextureFromImage(fudesumiRaw) // Upload CPU (RAM) image to GPU (VRAM)
	rl.UnloadImage(&fudesumiRaw)                     // Unload CPU (RAM) image data

	// Generate a checked texture by code (1024x1024 pixels)
	width := 1024
	height := 1024

	// Dynamic memory allocation to store pixels data (Color type)
	pixels := make([]colorex.RGBA, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if ((x/32+y/32)/1)%2 == 0 {
				pixels[y*height+x] = rl.Orange
			} else {
				pixels[y*height+x] = rl.Gold
			}
		}
	}

	// LoadImageEx was removed from raylib
	// Load pixels data into an image structure and create texture
	// checkedIm := rl.LoadImageEx(pixels, int32(width), int32(height))
	// checked := rl.LoadTextureFromImage(checkedIm)
	// rl.UnloadImage(checkedIm) // Unload CPU (RAM) image data

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		//rl.DrawTexture(&checked, screenWidth/2-checked.Width/2, screenHeight/2-checked.Height/2, rl.Fade(rl.White, 0.5))
		rl.DrawTexture(fudesumi, 430, -30, rl.White)

		rl.DrawText("CHECKED TEXTURE ", 84, 100, 30, rl.Brown)
		rl.DrawText("GENERATED by CODE", 72, 164, 30, rl.Brown)
		rl.DrawText("and RAW IMAGE LOADING", 46, 226, 30, rl.Brown)

		rl.EndDrawing()
	}

	rl.UnloadTexture(&fudesumi) // Texture unloading
	//rl.UnloadTexture(&checked)  // Texture unloading

	rl.CloseWindow()
}
