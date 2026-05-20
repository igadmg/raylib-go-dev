//go:build darwin && rgfw && !sdl && !sdl3

package rl

/*
#cgo LDFLAGS: -framework Foundation -framework AppKit -framework CoreVideo
#cgo CFLAGS: -DPLATFORM_DESKTOP_RGFW -Wno-deprecated-declarations -Wno-implicit-const-int-float-conversion -Wno-typedef-redefinition -Wno-extern-initializer -Wno-unused-value
#cgo CFLAGS: -Wno-incompatible-pointer-types -Wno-incompatible-function-pointer-types -Wno-incompatible-pointer-types-discards-qualifiers -Wno-macro-redefined

#cgo !es2,!es3 LDFLAGS: -framework OpenGL

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
*/
import "C"
