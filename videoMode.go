// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Window/Window.h>
// #include <stdlib.h>
// sfVideoMode videoModeAt(size_t index, sfVideoMode* modes) { return modes[index]; }
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type VideoMode struct {
	Width        uint ///< Video mode width, in pixels
	Height       uint ///< Video mode height, in pixels
	BitsPerPixel uint ///< Video mode pixel depth, in bits per pixels
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Get the current desktop video mode
func GetDesktopVideoMode() (videoMode VideoMode) {
	videoMode.fromC(C.sfVideoMode_getDesktopMode())
	return
}

// Tell whether or not a video mode is valid
//
// The validity of video modes is only relevant when using
// fullscreen windows; otherwise any video mode can be used
// with no restriction.
//
// return true if the video mode is valid for fullscreen mode
func (this *VideoMode) IsValid() bool {
	return sfBool2Go(C.sfVideoMode_isValid(this.toC()))
}

// Retrieve all the video modes supported in fullscreen mode
//
// When creating a fullscreen window, the video mode is restricted
// to be compatible with what the graphics driver and monitor
// support. This function returns the complete list of all video
// modes that can be used in fullscreen mode.
// The returned array is sorted from best to worst, so that
// the first element will always give the best mode (higher
// width, height and bits-per-pixel).
//
// Slice containing all the supported fullscreen modes
func GetFullscreenModes() []VideoMode {
	c := C.size_t(0)
	cVideoModesPtr := C.sfVideoMode_getFullscreenModes(&c)

	modes := make([]VideoMode, c)
	for i := uint(0); i < uint(c); i++ {
		modes[i].fromC(C.videoModeAt(C.size_t(i), cVideoModesPtr))
	}
	return modes
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *VideoMode) fromC(videoMode C.sfVideoMode) {
	this.Width = uint(videoMode.width)
	this.Height = uint(videoMode.height)
	this.BitsPerPixel = uint(videoMode.bitsPerPixel)
}

func (this *VideoMode) toC() C.sfVideoMode {
	return C.sfVideoMode{width: C.uint(this.Width), height: C.uint(this.Height), bitsPerPixel: C.uint(this.BitsPerPixel)}
}
