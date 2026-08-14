//go:build darwin && angle && !rgfw && !sdl && !sdl3

package rl

// GLFW's EGL backend dlopen()s "libEGL.dylib" / "libGLESv2.dylib" by bare name at
// runtime (no link-time dependency). macOS has no system libEGL, so ANGLE's copies
// must ship next to the built binary; this rpath makes dlopen's dyld search find
// them there. See third_party/angle/build-macos.sh (main repo) to produce the libs
// and go.build.copy-angle-libs.sh to place them next to the binary.

/*
#cgo LDFLAGS: -Wl,-rpath,@executable_path
*/
import "C"
