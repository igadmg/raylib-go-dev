module github.com/igadmg/raylib-go/raygui

go 1.24

replace (
	deedles.dev/xiter => ../../xiter
	github.com/igadmg/gamemath => ../../gamemath
	github.com/igadmg/goex => ../../goex
	github.com/igadmg/raylib-go/raylib => ../raylib
)

require (
	github.com/igadmg/gamemath v1.6.0
	github.com/igadmg/goex v0.0.0-20250325133153-61aee7990ef8
	github.com/igadmg/raylib-go/raylib v0.0.0-20250327112125-46827dd07487
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
