module github.com/igadmg/raylib-go/raylib

go 1.24

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace github.com/igadmg/goex => ../../goex

replace deedles.dev/xiter => ../../xiter

require (
	github.com/ebitengine/purego v0.8.2
	github.com/igadmg/raylib-go/raymath v1.6.0
	golang.org/x/sys v0.30.0
)

require golang.org/x/exp v0.0.0-20250228200357-dead58393ab7 // indirect
