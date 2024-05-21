module github.com/igadmg/raylib-go/raylib

go 1.21

replace github.com/igadmg/raylib-go/raymath => ../raymath

require (
	github.com/igadmg/raylib-go/raymath v1.6.0
	github.com/ebitengine/purego v0.7.1
	github.com/stretchr/testify v1.9.0
	golang.org/x/sys v0.19.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
