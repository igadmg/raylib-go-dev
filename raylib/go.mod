module github.com/igadmg/raylib-go/raylib

go 1.22.2

toolchain go1.22.4

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace github.com/igadmg/goex => ../../goex

require (
	github.com/ebitengine/purego v0.7.1
	github.com/igadmg/raylib-go/raymath v1.6.0
	golang.org/x/sys v0.21.0
)

require (
	github.com/igadmg/goex v0.0.0-20240505224334-5f13a70c4df7 // indirect
	golang.org/x/exp v0.0.0-20240604190554-fc45aab8b7f8 // indirect
)
