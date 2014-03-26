// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Window/Window.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	StyleNone       WindowStyle = C.sfNone         ///< No border / title bar (this flag and all others are mutually exclusive)
	StyleTitlebar   WindowStyle = C.sfTitlebar     ///< Title bar + fixed border
	StyleResize     WindowStyle = C.sfResize       ///< Titlebar + resizable border + maximize button
	StyleClose      WindowStyle = C.sfClose        ///< Titlebar + close button
	StyleFullscreen WindowStyle = C.sfFullscreen   ///< Fullscreen mode (this flag and all others are mutually exclusive)
	StyleDefault    WindowStyle = C.sfDefaultStyle ///< Default window style
)

type WindowStyle int

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Window struct {
	cptr *C.sfWindow
}

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

// Window and RenderWindow are SystemWindows
type SystemWindow interface {
	SetVSyncEnabled(bool)
	SetFramerateLimit(uint)
	SetJoystickThreshold(float32)
	SetKeyRepeatEnabled(bool)
	Display()
	IsOpen() bool
	Close()
	SetTitle(string)
	SetIcon(uint, uint, []byte) error
	SetMouseCursorVisible(bool)
	SetActive(bool) bool
}

//TEST
var _ SystemWindow = &RenderWindow{}
var _ SystemWindow = &Window{}

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

// Construct a new window
//
// 	videoMode:       Video mode to use
// 	title:           Title of the window
// 	style:           Window style
// 	contextSettings: Creation settings (pass nil to use default values)
func NewWindow(videoMode VideoMode, title string, style WindowStyle, contextSettings ContextSettings) (window *Window) {
	//string conversion
	utf32 := strToRunes(title)

	//convert contextSettings to C
	cs := contextSettings.toC()

	//create the window
	window = &Window{C.sfWindow_createUnicode(videoMode.toC(), (*C.sfUint32)(unsafe.Pointer(&utf32[0])), C.sfUint32(style), &cs)}

	//GC cleanup
	runtime.SetFinalizer(window, (*Window).destroy)

	return window
}

// Get the creation settings of a window
func (this *Window) GetSettings() (settings ContextSettings) {
	settings.fromC(C.sfWindow_getSettings(this.cptr))
	return
}

// Change the size of the rendering region of a window
//
// 	size: New size, in pixels
func (this *Window) SetSize(size Vector2u) {
	C.sfWindow_setSize(this.cptr, size.toC())
}

// Get the size of the rendering region of a window
func (this *Window) GetSize() Vector2u {
	size := C.sfWindow_getSize(this.cptr)
	return Vector2u{uint(size.x), uint(size.y)}
}

// Change the position of a window on screen
//
// Only works for top-level windows
//
// 	pos: New position, in pixels
func (this *Window) SetPosition(pos Vector2i) {
	C.sfWindow_setPosition(this.cptr, pos.toC())
}

// Get the position of a render window
func (this *Window) GetPosition() (pos Vector2i) {
	pos.fromC(C.sfWindow_getPosition(this.cptr))
	return
}

// Tell whether or not a window is opened
func (this *Window) IsOpen() bool {
	return sfBool2Go(C.sfWindow_isOpen(this.cptr))
}

// Close a window (but doesn't destroy the internal data)
func (this *Window) Close() {
	C.sfWindow_close(this.cptr)
}

// Destroy an existing window
func (this *Window) destroy() {
	globalMutex.Lock()
	C.sfWindow_destroy(this.cptr)
	globalMutex.Unlock()
}

// Get the event on top of event queue of a window, if any, and pop it
//
// returns nil if there are no events left.
func (this *Window) PollEvent() Event {
	cEvent := C.sfEvent{}

	globalMutex.Lock()
	hasEvent := C.sfWindow_pollEvent(this.cptr, &cEvent)
	globalMutex.Unlock()

	if hasEvent != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

// Wait for an event and return it
func (this *Window) WaitEvent() Event {
	cEvent := C.sfEvent{}

	globalMutex.Lock()
	hasError := C.sfWindow_waitEvent(this.cptr, &cEvent)
	globalMutex.Unlock()

	if hasError != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

// Change the title of a window
//
// 	title: New title
func (this *Window) SetTitle(title string) {
	utf32 := strToRunes(title)

	C.sfWindow_setUnicodeTitle(this.cptr, (*C.sfUint32)(unsafe.Pointer(&utf32[0])))
}

// Change a window's icon
//
// 	width:  Icon's width, in pixels
// 	height: Icon's height, in pixels
// 	pixels: Slice of pixels, format must be RGBA 32 bits
func (this *Window) SetIcon(width, height uint, data []byte) error {
	if len(data) >= int(width*height*4) {
		C.sfWindow_setIcon(this.cptr, C.uint(width), C.uint(height), (*C.sfUint8)(&data[0]))
		return nil
	}
	return errors.New("SetIcon: Slice length does not match specified dimensions")
}

// Limit the framerate to a maximum fixed frequency for a window
//
// 	limit: Framerate limit, in frames per seconds (use 0 to disable limit)
func (this *Window) SetFramerateLimit(limit uint) {
	C.sfWindow_setFramerateLimit(this.cptr, C.uint(limit))
}

///Change the joystick threshold, ie. the value below which no move event will be generated
//
// threshold: New threshold, in range [0, 100]
func (this *Window) SetJoystickThreshold(threshold float32) {
	C.sfWindow_setJoystickThreshold(this.cptr, C.float(threshold))
}

// Enable or disable automatic key-repeat
//
// If key repeat is enabled, you will receive repeated
// KeyPress events while keeping a key pressed. If it is disabled,
// you will only get a single event when the key is pressed.
//
// Key repeat is enabled by default.
func (this *Window) SetKeyRepeatEnabled(enabled bool) {
	C.sfWindow_setKeyRepeatEnabled(this.cptr, goBool2C(enabled))
}

// Display a window on screen
func (this *Window) Display() {
	globalMutex.Lock()
	C.sfWindow_display(this.cptr)
	globalMutex.Unlock()
}

// Enable / disable vertical synchronization on a window
//
// 	enabled: true to enable v-sync, false to deactivate
func (this *Window) SetVSyncEnabled(enabled bool) {
	globalMutex.Lock()
	C.sfWindow_setVerticalSyncEnabled(this.cptr, goBool2C(enabled))
	globalMutex.Unlock()
}

// Activate or deactivate a window as the current target for rendering
//
// 	active: true to activate, false to deactivate
//
// return True if operation was successful, false otherwise
func (this *Window) SetActive(active bool) bool {
	globalMutex.Lock()
	success := sfBool2Go(C.sfWindow_setActive(this.cptr, goBool2C(active)))
	globalMutex.Unlock()
	return success
}

// Show or hide the mouse cursor on a render window
//
// 	visible: true to show, false to hide
func (this *Window) SetMouseCursorVisible(visible bool) {
	C.sfWindow_setMouseCursorVisible(this.cptr, goBool2C(visible))
}
