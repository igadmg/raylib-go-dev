//go:build windows && rgfw && !sdl && !sdl3

package rl

/*
#cgo LDFLAGS: -lgdi32 -lwinmm
#cgo CFLAGS: -Iexternal -DPLATFORM_DESKTOP_RGFW -Wno-stringop-overflow -Wno-discarded-qualifiers

#cgo !es2,!es3 LDFLAGS: -lopengl32

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
*/
import "C"
