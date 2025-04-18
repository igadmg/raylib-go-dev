module github.com/igadmg/raylib-go/raygui

go 1.24

replace (
	deedles.dev/xiter => ../../xiter
	github.com/igadmg/goex => ../../goex
	github.com/igadmg/raylib-go/raylib => ../raylib
	github.com/igadmg/raylib-go/raymath => ../raymath
)

require (
	github.com/igadmg/goex v0.0.0-20250321131421-ccb743b21181
	github.com/igadmg/raylib-go/raylib v0.0.0-20240518160852-5314a4a2099a
	github.com/igadmg/raylib-go/raymath v1.6.0
)

require (
	github.com/chewxy/math32 v1.11.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/sys v0.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
