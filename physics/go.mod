module github.com/igadmg/raylib-go/physics

go 1.22.2

toolchain go1.22.4

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace github.com/igadmg/raylib-go/raylib => ../raylib

replace github.com/DeedleFake/xiter => ../../xiter

require (
	github.com/igadmg/raylib-go/raylib v0.0.0-20240518160852-5314a4a2099a
	github.com/igadmg/raylib-go/raymath v1.6.0
)

require (
	github.com/DeedleFake/xiter v0.0.0-20240215152241-9fc873aaff80 // indirect
	github.com/ebitengine/purego v0.7.1 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/sys v0.22.0 // indirect
)
