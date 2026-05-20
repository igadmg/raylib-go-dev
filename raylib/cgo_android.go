//go:build android

package rl

/*
#cgo LDFLAGS: -llog -landroid -lEGL -lGLESv2 -lOpenSLES -lm -Wl,--wrap=fopen
#cgo CFLAGS: -DPLATFORM_ANDROID -DPLATFORM_ANDROID_GOLANG -DGRAPHICS_API_OPENGL_ES2 -Iexternal/android/native_app_glue -Wno-implicit-const-int-float-conversion

#cgo arm CFLAGS: -march=armv7-a -mfloat-abi=softfp -mfpu=vfpv3-d16

#include "platforms/android_native_app_glue.c"
*/
import "C"
