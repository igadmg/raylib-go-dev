module github.com/igadmg/raylib-go/physics

go 1.24

replace (
	deedles.dev/xiter => ../../xiter
	github.com/igadmg/gamemath => ../../gamemath
	github.com/igadmg/raylib-go/raylib => ../raylib
)

require (
	github.com/igadmg/gamemath v1.6.0
	github.com/igadmg/raylib-go/raylib v0.0.0-20250327112125-46827dd07487
)

require (
	github.com/chewxy/math32 v1.11.1 // indirect
	github.com/ebitengine/purego v0.8.2 // indirect
	github.com/igadmg/goex v0.0.0-20250407220752-712c023573b8 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/sys v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
