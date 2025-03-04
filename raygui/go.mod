module github.com/igadmg/raylib-go/raygui

go 1.24

replace github.com/igadmg/raylib-go/raylib => ../raylib

replace github.com/igadmg/raylib-go/raymath => ../raymath

replace deedles.dev/xiter => ../../xiter

require github.com/igadmg/raylib-go/raylib v0.0.0-20240518160852-5314a4a2099a

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.2 // indirect
	github.com/igadmg/raylib-go/raymath v1.6.0
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20250228200357-dead58393ab7 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
