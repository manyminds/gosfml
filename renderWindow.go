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

package GoSFML2

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

type Drawable interface{
	Draw(target RenderTarget, renderStates *RenderStates)	
}

/////////////////////////////////////
///		CONTRUCTOR
/////////////////////////////////////

func NewRenderWindow(videoMode VideoMode, title string, style int, contextSettings *ContextSettings) *RenderWindow {
	//transform GoString into CString
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	//create the window
	window := &RenderWindow{
		C.sfRenderWindow_create(C.sfVideoMode{C.uint(videoMode.Width), C.uint(videoMode.Height), C.uint(videoMode.BitsPerPixel)},
			cTitle,
			C.sfUint32(style),
			(*C.sfContextSettings)(unsafe.Pointer(contextSettings)))}

	//GC cleanup
	runtime.SetFinalizer(window, (*RenderWindow).Destroy)

	return window
}

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

func (this *RenderWindow) GetSettings() ContextSettings {
	csettings := C.sfRenderWindow_getSettings(this.cptr)
	return ContextSettings{uint(csettings.depthBits),
		uint(csettings.stencilBits),
		uint(csettings.antialiasingLevel),
		uint(csettings.majorVersion),
		uint(csettings.minorVersion)}
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

func (this *RenderWindow) PollEvent() Event {
	cEvent := new(RawEvent)
	
	hasEvent := C.sfRenderWindow_pollEvent(this.cptr, (*C.sfEvent)(unsafe.Pointer(cEvent)))

	if hasEvent != 0 {
		return HandleEvent(cEvent)
	}
	return nil
}

func (this *RenderWindow) SetVSyncEnabled(enabled bool) {
	C.sfRenderWindow_setVerticalSyncEnabled(this.cptr, goBool2C(enabled))
}

func (this *RenderWindow) SetMouseCursorVisible(visible bool) {
	C.sfRenderWindow_setMouseCursorVisible(this.cptr, goBool2C(visible))
}

func (this *RenderWindow) SetKeyRepeaterEnabled(enabled bool) {
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
	drawable.Draw(this,renderStates)
}

func (this *RenderWindow) ConvertCoords(pos Vector2i, view *View) (coord Vector2f) {
	if view == nil { view = this.GetDefaultView() }
	
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

//Test
func (this *RenderWindow) AsWindow() *Window {
	return &Window{this.cptr}
}
