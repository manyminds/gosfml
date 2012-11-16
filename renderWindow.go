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
 #include <SFML/Graphics.h>
 #include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type RenderWindow struct {
	cptr *C.sfRenderWindow
}

/////////////////////////////////////
///		CONTRUCTOR
/////////////////////////////////////

func NewRenderWindow(videoMode VideoMode, title string, style int, contextSettings *ContextSettings) (window *RenderWindow) {
	//transform GoString into CString
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	//create the window
	if contextSettings != nil {
		csettings := contextSettings.toC()
		window = &RenderWindow{C.sfRenderWindow_create(videoMode.toC(), cTitle, C.sfUint32(style), &csettings)}
	} else {
		window = &RenderWindow{C.sfRenderWindow_create(videoMode.toC(), cTitle, C.sfUint32(style), nil)}
	}

	//GC cleanup
	runtime.SetFinalizer(window, (*RenderWindow).Destroy)

	return window
}

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

func (this *RenderWindow) GetSettings() (settings ContextSettings) {
	settings.fromC(C.sfRenderWindow_getSettings(this.cptr))
	return
}

func (this *RenderWindow) SetSize(size Vector2u) {
	C.sfRenderWindow_setSize(this.cptr, size.toC())
}

func (this *RenderWindow) GetSize() (size Vector2u) {
	size.fromC(C.sfRenderWindow_getSize(this.cptr))
	return
}

func (this *RenderWindow) SetPosition(pos Vector2i) {
	C.sfRenderWindow_setPosition(this.cptr, pos.toC())
}

func (this *RenderWindow) GetPosition() (pos Vector2i) {
	pos.fromC(C.sfRenderWindow_getPosition(this.cptr))
	return
}

func (this *RenderWindow) IsOpen() bool {
	return sfBool2Go(C.sfRenderWindow_isOpen(this.cptr))
}

func (this *RenderWindow) Close() {
	C.sfRenderWindow_close(this.cptr)
}

func (this *RenderWindow) Destroy() {
	C.sfRenderWindow_destroy(this.cptr)
	this.cptr = nil
}

func (this *RenderWindow) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	C.sfRenderWindow_setTitle(this.cptr, cTitle)
}

func (this *RenderWindow) SetIcon(width, height uint, data []byte) error {
	if len(data) > 0 {
		C.sfRenderWindow_setIcon(this.cptr, C.uint(width), C.uint(height), (*C.sfUint8)(&data[0]))
		return nil
	}
	return &Error{"SetIcon: no data"}
}

// returns nil if there is no event
func (this *RenderWindow) PollEvent() Event {
	cEvent := C.sfEvent{}
	hasEvent := C.sfRenderWindow_pollEvent(this.cptr, &cEvent)

	if hasEvent != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

func (this *RenderWindow) WaitEvent() Event {
	cEvent := C.sfEvent{}
	hasError := C.sfRenderWindow_waitEvent(this.cptr, &cEvent)

	if hasError != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

func (this *RenderWindow) SetVSyncEnabled(enabled bool) {
	C.sfRenderWindow_setVerticalSyncEnabled(this.cptr, goBool2C(enabled))
}

func (this *RenderWindow) SetMouseCursorVisible(visible bool) {
	C.sfRenderWindow_setMouseCursorVisible(this.cptr, goBool2C(visible))
}

func (this *RenderWindow) SetKeyRepeatEnabled(enabled bool) {
	C.sfRenderWindow_setKeyRepeatEnabled(this.cptr, goBool2C(enabled))
}

func (this *RenderWindow) SetVisible(visible bool) {
	C.sfRenderWindow_setVisible(this.cptr, goBool2C(visible))
}

func (this *RenderWindow) SetActive(active bool) {
	C.sfRenderWindow_setActive(this.cptr, goBool2C(active))
}

func (this *RenderWindow) SetFramerateLimit(limit uint) {
	C.sfRenderWindow_setFramerateLimit(this.cptr, C.uint(limit))
}

func (this *RenderWindow) SetJoystickThreshold(threshold float32) {
	C.sfRenderWindow_setJoystickThreshold(this.cptr, C.float(threshold))
}

func (this *RenderWindow) Display() {
	C.sfRenderWindow_display(this.cptr)
}

func (this *RenderWindow) Clear(color Color) {
	C.sfRenderWindow_clear(this.cptr, color.toC())
}

func (this *RenderWindow) GetView() *View {
	return &View{C.sfRenderWindow_getView(this.cptr)}
}

func (this *RenderWindow) GetDefaultView() *View {
	return &View{C.sfRenderWindow_getDefaultView(this.cptr)}
}

func (this *RenderWindow) SetView(view *View) {
	C.sfRenderWindow_setView(this.cptr, view.cptr)
}

func (this *RenderWindow) Draw(drawable Drawable, renderStates *RenderStates) {
	drawable.Draw(this, renderStates)
}

func (this *RenderWindow) ConvertCoords(pos Vector2i, view *View) (coord Vector2f) {
	if view == nil {
		view = this.GetDefaultView()
	}

	coord.fromC(C.sfRenderWindow_convertCoords(this.cptr, pos.toC(), view.cptr))
	return
}

func (this *RenderWindow) GetViewport(view *View) (viewport Recti) {
	viewport.fromC(C.sfRenderWindow_getViewport(this.cptr, view.cptr))
	return
}

func (this *RenderWindow) PushGLStates() {
	C.sfRenderWindow_pushGLStates(this.cptr)
}

func (this *RenderWindow) PopGLStates() {
	C.sfRenderWindow_popGLStates(this.cptr)
}

func (this *RenderWindow) ResetGLStates() {
	C.sfRenderWindow_resetGLStates(this.cptr)
}

func (this *RenderWindow) Capture() *Image {
	return newImageFromPtr(C.sfRenderWindow_capture(this.cptr))
}
