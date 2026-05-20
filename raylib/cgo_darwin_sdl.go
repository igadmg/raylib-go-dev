//go:build darwin && (sdl || sdl3) && !rgfw

package rl

/*
#cgo LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation
#cgo CFLAGS: -Wno-deprecated-declarations -Wno-implicit-const-int-float-conversion
#cgo sdl CFLAGS: -DPLATFORM_DESKTOP_SDL
#cgo sdl3 CFLAGS: -DPLATFORM_DESKTOP_SDL -DPLATFORM_DESKTOP_SDL3
#cgo sdl pkg-config: sdl2
#cgo sdl3 pkg-config: sdl3

#cgo !es2,!es3 LDFLAGS: -framework OpenGL

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
*/
import "C"
