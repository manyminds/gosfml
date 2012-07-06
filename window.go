package GoSFML2

/*
 #include <SFML/Window.h>
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

type ContextSettings struct {
	DepthBits         uint ///< Bits of the depth buffer
	StencilBits       uint ///< Bits of the stencil buffer
	AntialiasingLevel uint ///< Level of antialiasing
	MajorVersion      uint ///< Major number of the context version to create
	MinorVersion      uint ///< Minor number of the context version to create
}

type Window struct {
	cptr *C.sfWindow
}

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

func CreateWindow(videoMode VideoMode, title string, style int, contextSettings *ContextSettings) *Window {
	//transform GoString into CString
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	//create the window
	window := &Window{
		C.sfWindow_create(C.sfVideoMode{C.uint(videoMode.Width), C.uint(videoMode.Height), C.uint(videoMode.BitsPerPixel)},
			cTitle,
			C.sfUint32(style),
			(*C.sfContextSettings)(unsafe.Pointer(contextSettings))),
	}

	//GC cleanup
	runtime.SetFinalizer(window, (*Window).Destroy)

	return window
}

func (this *Window) GetSettings() ContextSettings {
	csettings := C.sfWindow_getSettings(this.cptr)
	return ContextSettings{uint(csettings.depthBits),
		uint(csettings.stencilBits),
		uint(csettings.antialiasingLevel),
		uint(csettings.majorVersion),
		uint(csettings.minorVersion)}
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
	cEvent := new(RawEvent)

	runtime.LockOSThread()
	r := C.sfWindow_pollEvent(this.cptr, (*C.sfEvent)(unsafe.Pointer(cEvent)))
	runtime.UnlockOSThread()

	if r != 0 {
		return HandleEvent(cEvent)
	}
	return nil
}
