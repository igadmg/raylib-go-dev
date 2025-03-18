module github.com/igadmg/raylib-go/raylib

go 1.24

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace github.com/igadmg/goex => ../../goex

replace deedles.dev/xiter => ../../xiter

require (
	github.com/ebitengine/purego v0.8.2
	github.com/igadmg/goex v0.0.0-20250312230527-f6fa5b3c2d75
	github.com/igadmg/raylib-go/raymath v1.6.0
	golang.org/x/sys v0.31.0
)

require (
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
