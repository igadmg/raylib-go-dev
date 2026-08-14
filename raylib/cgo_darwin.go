//go:build darwin && !rgfw && !sdl && !sdl3

package rl

/*
#include "external/glfw/src/context.c"
#include "external/glfw/src/init.c"
#include "external/glfw/src/input.c"
#include "external/glfw/src/monitor.c"
#include "external/glfw/src/platform.c"
#include "external/glfw/src/vulkan.c"
#include "external/glfw/src/window.c"

#include "external/glfw/src/cocoa_init.m"
#include "external/glfw/src/cocoa_joystick.m"
#include "external/glfw/src/cocoa_monitor.m"
#include "external/glfw/src/cocoa_time.c"
#include "external/glfw/src/cocoa_window.m"
#include "external/glfw/src/posix_module.c"
#include "external/glfw/src/posix_thread.c"
#include "external/glfw/src/nsgl_context.m"
#include "external/glfw/src/egl_context.c"
#include "external/glfw/src/osmesa_context.c"

GLFWbool _glfwConnectNull(int platformID, _GLFWplatform* platform) {
	return GLFW_TRUE;
}

#cgo LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation
#cgo CFLAGS: -x objective-c -Iexternal/glfw/include -D_GLFW_COCOA -D_GLFW_USE_CHDIR -D_GLFW_USE_MENUBAR -D_GLFW_USE_RETINA -Wno-deprecated-declarations -Wno-implicit-const-int-float-conversion -DPLATFORM_DESKTOP

#cgo !es2,!es3,!angle LDFLAGS: -framework OpenGL

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3,!angle CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
// angle: link against Google's ANGLE (libEGL/libGLESv2) instead of native OpenGL/Metal;
// see third_party/angle in the main repo for build scripts and cgo_darwin_angle.go for rpath setup.
#cgo angle,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
*/
import "C"
