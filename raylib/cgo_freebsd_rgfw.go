//go:build freebsd && rgfw && !linux && !drm && !sdl && !sdl3 && !android

package rl

/*
#cgo CFLAGS: -I. -I/usr/local/include -DPLATFORM_DESKTOP_RGFW
#cgo LDFLAGS: -L/usr/local/lib -lX11 -lXrandr -lXinerama -lXi -lXxf86vm -lXcursor -lm -lpthread -ldl -lrt

#cgo !es2,!es3 LDFLAGS: -lGL

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
*/
import "C"
