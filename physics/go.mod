module github.com/igadmg/raylib-go/physics

go 1.22

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace github.com/igadmg/raylib-go/raylib => ../raylib

require (
	github.com/igadmg/raylib-go/raylib v0.0.0-20240518160852-5314a4a2099a
	github.com/igadmg/raylib-go/raymath v1.6.0
)

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	golang.org/x/exp v0.0.0-20240604190554-fc45aab8b7f8 // indirect
	golang.org/x/sys v0.21.0 // indirect
)
