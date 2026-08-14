//go:build linux && angle && !rgfw && !drm && !sdl && !sdl3 && !android

package rl

// GLFW's EGL backend dlopen()s "libEGL.so.1" / "libGLESv2.so.1" by bare name at
// runtime (no link-time dependency). System Mesa usually ships its own libEGL, so
// without this rpath the dynamic linker would find Mesa's copy before ANGLE's.
// ANGLE's libs must ship next to the built binary; $ORIGIN makes dlopen's search
// find them there ahead of the system search path. See third_party/angle/build-linux.sh
// (main repo) to produce the libs and go.build.copy-angle-libs.sh to place them.

/*
#cgo LDFLAGS: -Wl,-rpath,$ORIGIN
*/
import "C"
