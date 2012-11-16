/*
Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
This software is provided 'as-is', without any express or implied warranty.
In no event will the authors be held liable for any damages arising from the use of this software.
Permission is granted to anyone to use this software for any purpose, including commercial applications, 
and to alter it and redistribute it freely, subject to the following restrictions:
	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
	3. 	This notice may not be removed or altered from any source distribution.
*/

package gosfml2

/*
 #include <SFML/Window/Window.h>
 #include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Style_None         = 0                                           ///< No border / title bar (this flag and all others are mutually exclusive)
	Style_Titlebar     = 1 << 0                                      ///< Title bar + fixed border
	Style_Resize       = 1 << 1                                      ///< Titlebar + resizable border + maximize button
	Style_Close        = 1 << 2                                      ///< Titlebar + close button
	Style_Fullscreen   = 1 << 3                                      ///< Fullscreen mode (this flag and all others are mutually exclusive)
	Style_DefaultStyle = Style_Titlebar | Style_Resize | Style_Close ///< Default window style
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Window struct {
	cptr *C.sfWindow
}

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

//implemented by Window, RenderWindow
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
	SetActive(bool)
}

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

func NewWindow(videoMode VideoMode, title string, style int, contextSettings *ContextSettings) (window *Window) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	//create the window
	if contextSettings != nil {
		csettings := contextSettings.toC()
		window = &Window{C.sfWindow_create(videoMode.toC(), cTitle, C.sfUint32(style), &csettings)}
	} else {
		window = &Window{C.sfWindow_create(videoMode.toC(), cTitle, C.sfUint32(style), nil)}
	}

	//GC cleanup
	runtime.SetFinalizer(window, (*Window).Destroy)

	return window
}

func (this *Window) GetSettings() (settings ContextSettings) {
	settings.fromC(C.sfWindow_getSettings(this.cptr))
	return
}

func (this *Window) SetSize(size Vector2u) {
	C.sfWindow_setSize(this.cptr, size.toC())
}

func (this *Window) GetSize() Vector2u {
	size := C.sfWindow_getSize(this.cptr)
	return Vector2u{uint(size.x), uint(size.y)}
}

func (this *Window) SetPosition(pos Vector2i) {
	C.sfWindow_setPosition(this.cptr, pos.toC())
}

func (this *Window) GetPosition() (pos Vector2i) {
	pos.fromC(C.sfWindow_getPosition(this.cptr))
	return
}

func (this *Window) IsOpen() bool {
	return sfBool2Go(C.sfWindow_isOpen(this.cptr))
}

func (this *Window) Close() {
	C.sfWindow_close(this.cptr)
}

func (this *Window) Destroy() {
	C.sfWindow_destroy(this.cptr)
	this.cptr = nil
}

func (this *Window) PollEvent() Event {
	cEvent := C.sfEvent{}
	hasEvent := C.sfWindow_pollEvent(this.cptr, &cEvent)

	if hasEvent != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

func (this *Window) WaitEvent() Event {
	cEvent := C.sfEvent{}
	hasError := C.sfWindow_waitEvent(this.cptr, &cEvent)

	if hasError != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

func (this *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	C.sfWindow_setTitle(this.cptr, cTitle)
}

func (this *Window) SetIcon(width, height uint, data []byte) error {
	if len(data) > 0 {
		C.sfWindow_setIcon(this.cptr, C.uint(width), C.uint(height), (*C.sfUint8)(&data[0]))
		return nil
	}
	return &Error{"SetIcon: no data"}
}

func (this *Window) SetFramerateLimit(limit uint) {
	C.sfWindow_setFramerateLimit(this.cptr, C.uint(limit))
}

func (this *Window) SetJoystickThreshold(threshold float32) {
	C.sfWindow_setJoystickThreshold(this.cptr, C.float(threshold))
}

func (this *Window) SetKeyRepeatEnabled(enabled bool) {
	C.sfWindow_setKeyRepeatEnabled(this.cptr, goBool2C(enabled))
}

func (this *Window) Display() {
	C.sfWindow_display(this.cptr)
}

func (this *Window) SetVSyncEnabled(enabled bool) {
	C.sfWindow_setVerticalSyncEnabled(this.cptr, goBool2C(enabled))
}

func (this *Window) SetActive(active bool) {
	C.sfWindow_setActive(this.cptr, goBool2C(active))
}

func (this *Window) SetMouseCursorVisible(visible bool) {
	C.sfWindow_setMouseCursorVisible(this.cptr, goBool2C(visible))
}
