//go:build freebsd && !linux && (sdl || sdl3) && !rgfw && !drm && !android

package rl

/*
#cgo CFLAGS: -I. -I/usr/local/include
#cgo sdl CFLAGS: -DPLATFORM_DESKTOP_SDL
#cgo sdl3 CFLAGS: -DPLATFORM_DESKTOP_SDL -DPLATFORM_DESKTOP_SDL3
#cgo LDFLAGS: -L/usr/local/lib

#cgo sdl pkg-config: sdl2
#cgo sdl3 pkg-config: sdl3

#cgo !es2,!es3 LDFLAGS: -lGL

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
*/
import "C"
