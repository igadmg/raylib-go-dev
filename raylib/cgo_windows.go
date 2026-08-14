//go:build windows && !rgfw && !sdl && !sdl3

package rl

/*
#cgo CFLAGS: -I${SRCDIR}/../external/raylib/src -I${SRCDIR}/../external/raylib/src/external/glfw/include -std=gnu99 -Wno-missing-braces -Wno-unused-result -Wno-implicit-function-declaration -DPLATFORM_DESKTOP -D_GLFW_WIN32 -Wno-stringop-overflow
#cgo LDFLAGS: -lgdi32 -lwinmm -lole32
#cgo !es2,!es3,!angle LDFLAGS: -lopengl32

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3,!angle CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
// angle: link against Google's ANGLE (libEGL/libGLESv2, D3D11 backend) instead of native WGL;
// see third_party/angle in the main repo for build scripts. No rpath needed on Windows:
// the default DLL search order already checks the executable's own directory first.
#cgo angle,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES3

#include "external/glfw/src/context.c"
#include "external/glfw/src/init.c"
#include "external/glfw/src/input.c"
#include "external/glfw/src/monitor.c"
#include "external/glfw/src/platform.c"
#include "external/glfw/src/vulkan.c"
#include "external/glfw/src/window.c"

#include "external/glfw/src/win32_init.c"
#include "external/glfw/src/win32_joystick.c"
#include "external/glfw/src/win32_module.c"
#include "external/glfw/src/win32_monitor.c"
#include "external/glfw/src/win32_thread.c"
#include "external/glfw/src/win32_time.c"
#include "external/glfw/src/win32_window.c"
#include "external/glfw/src/wgl_context.c"
#include "external/glfw/src/egl_context.c"
#include "external/glfw/src/osmesa_context.c"

GLFWbool _glfwConnectNull(int platformID, _GLFWplatform* platform) {
	return GLFW_TRUE;
}
*/
import "C"
