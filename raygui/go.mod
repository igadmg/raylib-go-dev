module github.com/gen2brain/raylib-go/raygui

go 1.21

replace github.com/gen2brain/raylib-go/raylib => ../raylib

replace github.com/EliCDavis/vector => ../vector

require github.com/gen2brain/raylib-go/raylib v0.0.0-20240227114648-c3665eb9abf8

require (
	github.com/EliCDavis/vector v1.6.0 // indirect
	github.com/ebitengine/purego v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)
