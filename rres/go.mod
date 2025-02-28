module github.com/igadmg/raylib-go/rres

go 1.24

replace github.com/igadmg/raylib-go/raylib => ../raylib

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace deedles.dev/xiter => ../../xiter

require github.com/igadmg/raylib-go/raylib v0.0.0-20240518160852-5314a4a2099a

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	github.com/igadmg/raylib-go/raymath v1.6.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/sys v0.22.0 // indirect
)
