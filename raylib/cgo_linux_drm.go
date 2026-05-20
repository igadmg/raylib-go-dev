//go:build linux && drm && !rgfw && !sdl && !sdl3 && !android

package rl

/*
#cgo LDFLAGS: -lGLESv2 -lEGL -ldrm -lgbm -lpthread -lrt -lm -ldl
#cgo CFLAGS: -DPLATFORM_DRM -DGRAPHICS_API_OPENGL_ES2 -DEGL_NO_X11 -I/usr/include/libdrm
*/
import "C"
