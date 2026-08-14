//go:build linux && !rgfw && !drm && !sdl && !sdl3 && !android

package rl

/*
#include "external/glfw/src/context.c"
#include "external/glfw/src/init.c"
#include "external/glfw/src/input.c"
#include "external/glfw/src/monitor.c"
#include "external/glfw/src/platform.c"
#include "external/glfw/src/vulkan.c"
#include "external/glfw/src/window.c"

#if defined _GLFW_WAYLAND
#include "external/glfw/src/wl_init.c"
#include "external/glfw/src/wl_monitor.c"
#include "external/glfw/src/wl_window.c"
#endif

#if defined _GLFW_X11
#include "external/glfw/src/x11_init.c"
#include "external/glfw/src/x11_monitor.c"
#include "external/glfw/src/x11_window.c"
#include "external/glfw/src/glx_context.c"
#endif

#include "external/glfw/src/linux_joystick.c"
#include "external/glfw/src/posix_module.c"
#include "external/glfw/src/posix_poll.c"
#include "external/glfw/src/posix_thread.c"
#include "external/glfw/src/posix_time.c"
#include "external/glfw/src/xkb_unicode.c"
#include "external/glfw/src/egl_context.c"
#include "external/glfw/src/osmesa_context.c"

GLFWbool _glfwConnectNull(int platformID, _GLFWplatform* platform) {
	return GLFW_TRUE;
}

#cgo CFLAGS: -Iexternal/glfw/include -DPLATFORM_DESKTOP -Wno-stringop-overflow
#cgo LDFLAGS: -lm -pthread -ldl -lrt -lxkbcommon
#cgo !x11 LDFLAGS: -lwayland-client -lwayland-cursor -lwayland-egl

#cgo x11 CFLAGS: -D_GLFW_X11
#cgo wayland CFLAGS: -D_GLFW_WAYLAND
//cgo !x11,!wayland CFLAGS: -D_GLFW_X11 -D_GLFW_WAYLAND
#cgo !x11,!wayland CFLAGS: -D_GLFW_WAYLAND

#cgo !es2,!es3,!angle,!wayland LDFLAGS: -lGL

#cgo opengl11,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo opengl21,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo opengl43,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_43
#cgo !opengl11,!opengl21,!opengl43,!es2,!es3,!angle CFLAGS: -DGRAPHICS_API_OPENGL_33
#cgo es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES2
#cgo es3,!es2 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
// angle: link against Google's ANGLE (libEGL/libGLESv2, Vulkan backend) instead of native GLX;
// see third_party/angle in the main repo for build scripts and cgo_linux_angle.go for rpath setup.
#cgo angle,!es2,!es3 CFLAGS: -DGRAPHICS_API_OPENGL_ES3
*/
import "C"
