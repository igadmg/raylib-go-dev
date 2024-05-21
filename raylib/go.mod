module github.com/igadmg/raylib-go/raylib

go 1.21

replace github.com/igadmg/raylib-go/raymath => ../raymath

require (
	github.com/ebitengine/purego v0.7.1
	github.com/igadmg/raylib-go/raymath v1.6.0
	golang.org/x/sys v0.19.0
)

require golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
