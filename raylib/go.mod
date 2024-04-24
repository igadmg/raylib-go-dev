module github.com/gen2brain/raylib-go/raylib

go 1.21

replace github.com/EliCDavis/vector => ../vector

require (
	github.com/EliCDavis/vector v1.6.0
	github.com/ebitengine/purego v0.7.1
	golang.org/x/sys v0.19.0
)

require golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
