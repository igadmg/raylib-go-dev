module github.com/igadmg/raylib-go/physics

go 1.24

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace github.com/igadmg/raylib-go/raylib => ../raylib

replace deedles.dev/xiter => ../../xiter

require (
	github.com/igadmg/raylib-go/raylib v0.0.0-20240518160852-5314a4a2099a
	github.com/igadmg/raylib-go/raymath v1.6.0
)

require (
	github.com/ebitengine/purego v0.8.2 // indirect
	github.com/igadmg/goex v0.0.0-20250312230527-f6fa5b3c2d75 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/sys v0.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
