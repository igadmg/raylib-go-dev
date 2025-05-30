//go:build !rgfw && !sdl && !sdl3 && !drm && !android
// +build !rgfw,!sdl,!sdl3,!drm,!android

package rl

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"

import (
	"os"
	"unsafe"
)

// InitWindow - Initialize Window and OpenGL Graphics
func InitWindow[WT, HT IntegerT](width WT, height HT, title string) {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)

	ctitle := textAlloc(title)

	C.InitWindow(cwidth, cheight, ctitle)
}

// SetCallbackFunc - Sets callback function
func SetCallbackFunc(func()) {
}

// ShowCursor - Shows cursor
func ShowCursor() {
	C.ShowCursor()
}

// HideCursor - Hides cursor
func HideCursor() {
	C.HideCursor()
}

// IsCursorHidden - Returns true if cursor is not visible
func IsCursorHidden() bool {
	ret := C.IsCursorHidden()
	v := bool(ret)

	return v
}

// IsCursorOnScreen - Check if cursor is on the current screen.
func IsCursorOnScreen() bool {
	ret := C.IsCursorOnScreen()
	v := bool(ret)

	return v
}

// EnableCursor - Enables cursor
func EnableCursor() {
	C.EnableCursor()
}

// DisableCursor - Disables cursor
func DisableCursor() {
	C.DisableCursor()
}

// IsFileDropped - Check if a file have been dropped into window
func IsFileDropped() bool {
	ret := C.IsFileDropped()
	v := bool(ret)

	return v
}

// LoadDroppedFiles - Load dropped filepaths
func LoadDroppedFiles() []string {
	ret := C.LoadDroppedFiles()
	defer C.UnloadDroppedFiles(ret)

	tmpslice := (*[1 << 24]*C.char)(unsafe.Pointer(ret.paths))[:ret.count:ret.count]
	gostrings := make([]string, ret.count)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}

	return gostrings
}

// UnloadDroppedFiles - Unload dropped filepaths
func UnloadDroppedFiles() {
}

type asset struct {
	*os.File
}

// OpenAsset - Open asset
func OpenAsset(name string) (Asset, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return &asset{f}, nil
}

func (a *asset) Size() int64 {
	fi, err := a.Stat()
	if err != nil {
		return 0
	}
	return fi.Size()
}
