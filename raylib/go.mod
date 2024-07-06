module github.com/igadmg/raylib-go/raylib

go 1.22.2

toolchain go1.22.4

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace github.com/igadmg/goex => ../../goex

replace github.com/DeedleFake/xiter => ../../xiter

require (
	github.com/DeedleFake/xiter v0.0.0-20240215152241-9fc873aaff80
	github.com/ebitengine/purego v0.7.1
	github.com/igadmg/raylib-go/raymath v1.6.0
	golang.org/x/sys v0.22.0
)

require golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
