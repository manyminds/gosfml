// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/RenderWindow.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type RenderWindow struct {
	cptr *C.sfRenderWindow
	view *View
}

/////////////////////////////////////
///		CONTRUCTOR
/////////////////////////////////////

// Construct a new render window
//
// 	mode:            Video mode to use
// 	title:           Title of the window
// 	style:           Window style
// 	contextSettings: Creation settings
func NewRenderWindow(videoMode VideoMode, title string, style WindowStyle, contextSettings ContextSettings) (window *RenderWindow) {
	//string conversion
	utf32 := strToRunes(title)

	//convert contextSettings to C
	cs := contextSettings.toC()

	//create the window
	window = &RenderWindow{cptr: C.sfRenderWindow_createUnicode(videoMode.toC(), (*C.sfUint32)(unsafe.Pointer(&utf32[0])), C.sfUint32(style), &cs)}

	//create a copy of current view
	window.SetView(newViewFromPtr(C.sfRenderWindow_getView(window.cptr)))

	//GC cleanup
	runtime.SetFinalizer(window, (*RenderWindow).destroy)

	return window
}

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

// Get the creation settings of a render window
func (this *RenderWindow) GetSettings() (settings ContextSettings) {
	settings.fromC(C.sfRenderWindow_getSettings(this.cptr))
	return
}

// Change the size of the rendering region of a render window
//
// 	size: New size, in pixels
func (this *RenderWindow) SetSize(size Vector2u) {
	C.sfRenderWindow_setSize(this.cptr, size.toC())
}

// Get the size of the rendering region of a render window
func (this *RenderWindow) GetSize() (size Vector2u) {
	size.fromC(C.sfRenderWindow_getSize(this.cptr))
	return
}

// Change the position of a render window on screen
//
// Only works for top-level windows
//
// 	pos: New position, in pixels
func (this *RenderWindow) SetPosition(pos Vector2i) {
	C.sfRenderWindow_setPosition(this.cptr, pos.toC())
}

// Get the position of a render window
func (this *RenderWindow) GetPosition() (pos Vector2i) {
	pos.fromC(C.sfRenderWindow_getPosition(this.cptr))
	return
}

// Tell whether or not a render window is opened
func (this *RenderWindow) IsOpen() bool {
	return sfBool2Go(C.sfRenderWindow_isOpen(this.cptr))
}

// Close a render window (but doesn't destroy the internal data)
func (this *RenderWindow) Close() {
	C.sfRenderWindow_close(this.cptr)
}

// Destroy an existing render window
func (this *RenderWindow) destroy() {
	globalMutex.Lock()
	C.sfRenderWindow_destroy(this.cptr)
	globalMutex.Unlock()
}

// Change the title of a render window
//
// 	title: New title
func (this *RenderWindow) SetTitle(title string) {
	utf32 := strToRunes(title)

	C.sfRenderWindow_setUnicodeTitle(this.cptr, (*C.sfUint32)(unsafe.Pointer(&utf32[0])))
}

// Change a render window's icon
//
// 	width:  Icon's width, in pixels
// 	height: Icon's height, in pixels
// 	pixels: Slice of pixels, format must be RGBA 32 bits
func (this *RenderWindow) SetIcon(width, height uint, data []byte) error {
	if len(data) >= int(width*height*4) {
		C.sfRenderWindow_setIcon(this.cptr, C.uint(width), C.uint(height), (*C.sfUint8)(&data[0]))
		return nil
	}
	return errors.New("SetIcon: Slice length does not match specified dimensions")
}

// Get the event on top of event queue of a render window, if any, and pop it
//
// returns nil if there are no events left.
func (this *RenderWindow) PollEvent() Event {
	cEvent := C.sfEvent{}

	globalMutex.Lock()
	hasEvent := C.sfRenderWindow_pollEvent(this.cptr, &cEvent)
	globalMutex.Unlock()

	if hasEvent != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

// Wait for an event and return it
func (this *RenderWindow) WaitEvent() Event {
	cEvent := C.sfEvent{}

	globalMutex.Lock()
	hasError := C.sfRenderWindow_waitEvent(this.cptr, &cEvent)
	globalMutex.Unlock()

	if hasError != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

// Enable / disable vertical synchronization on a render window
//
// 	enabled: true to enable v-sync, false to deactivate
func (this *RenderWindow) SetVSyncEnabled(enabled bool) {
	globalMutex.Lock()
	C.sfRenderWindow_setVerticalSyncEnabled(this.cptr, goBool2C(enabled))
	globalMutex.Unlock()
}

// Show or hide the mouse cursor on a render window
//
// 	visible: true to show, false to hide
func (this *RenderWindow) SetMouseCursorVisible(visible bool) {
	C.sfRenderWindow_setMouseCursorVisible(this.cptr, goBool2C(visible))
}

// Enable or disable automatic key-repeat
//
// If key repeat is enabled, you will receive repeated
// KeyPress events while keeping a key pressed. If it is disabled,
// you will only get a single event when the key is pressed.
//
// Key repeat is enabled by default.
func (this *RenderWindow) SetKeyRepeatEnabled(enabled bool) {
	C.sfRenderWindow_setKeyRepeatEnabled(this.cptr, goBool2C(enabled))
}

// Show or hide a render window
//
// 	visible: true to show the window, false to hide it
func (this *RenderWindow) SetVisible(visible bool) {
	C.sfRenderWindow_setVisible(this.cptr, goBool2C(visible))
}

// Activate or deactivate a render window as the current target for rendering
//
// 	active: true to activate, false to deactivate
//
// return True if operation was successful, false otherwise
func (this *RenderWindow) SetActive(active bool) bool {
	globalMutex.Lock()
	success := sfBool2Go(C.sfRenderWindow_setActive(this.cptr, goBool2C(active)))
	globalMutex.Unlock()
	return success
}

// Limit the framerate to a maximum fixed frequency for a render window
//
// 	limit: Framerate limit, in frames per seconds (use 0 to disable limit)
func (this *RenderWindow) SetFramerateLimit(limit uint) {
	C.sfRenderWindow_setFramerateLimit(this.cptr, C.uint(limit))
}

// Change the joystick threshold, ie. the value below which no move event will be generated
//
// 	threshold: New threshold, in range [0, 100]
func (this *RenderWindow) SetJoystickThreshold(threshold float32) {
	C.sfRenderWindow_setJoystickThreshold(this.cptr, C.float(threshold))
}

// Display a render window on screen
func (this *RenderWindow) Display() {
	globalMutex.Lock()
	C.sfRenderWindow_display(this.cptr)
	globalMutex.Unlock()
}

// Clear a render window with the given color
//
// 	color: Fill color
func (this *RenderWindow) Clear(color Color) {
	C.sfRenderWindow_clear(this.cptr, color.toC())
}

// Get the current active view of a render window
func (this *RenderWindow) GetView() *View {
	return this.view
}

// Get the default view of a render window
func (this *RenderWindow) GetDefaultView() *View {
	return newViewFromPtr(C.sfRenderWindow_getDefaultView(this.cptr))
}

// Change the current active view of a render window
//
// 	view: Pointer to the new view
func (this *RenderWindow) SetView(view *View) {
	this.view = view
	C.sfRenderWindow_setView(this.cptr, view.toCPtr())
}

// Get the viewport of a view applied to this target
//
// 	view: Target view
func (this *RenderWindow) GetViewport(view *View) (viewport IntRect) {
	viewport.fromC(C.sfRenderWindow_getViewport(this.cptr, view.toCPtr()))
	return
}

// Draw a drawable object to the render-target
func (this *RenderWindow) Draw(drawer Drawer, renderStates RenderStates) {
	drawer.Draw(this, renderStates)
}

// Draw primitives defined by a slice of vertices
func (this *RenderWindow) DrawPrimitives(vertices []Vertex, primType PrimitiveType, renderStates RenderStates) {
	if len(vertices) > 0 {
		rs := renderStates.toC()
		C.sfRenderWindow_drawPrimitives(this.cptr, (*C.sfVertex)(unsafe.Pointer(&vertices[0])), C.uint(len(vertices)), C.sfPrimitiveType(primType), &rs)
	}
}

// Convert a point from window coordinates to world coordinates
//
// This function finds the 2D position that matches the
// given pixel of the render-window. In other words, it does
// the inverse of what the graphics card does, to find the
// initial position of a rendered pixel.
//
// Initially, both coordinate systems (world units and target pixels)
// match perfectly. But if you define a custom view or resize your
// render-window, this assertion is not true anymore, ie. a point
// located at (10, 50) in your render-window may map to the point
// (150, 75) in your 2D world -- if the view is translated by (140, 25).
//
// This function is typically used to find which point (or object) is
// located below the mouse cursor.
//
// This version uses a custom view for calculations, see the other
// overload of the function if you want to use the current view of the
// render-window.
//
// 	point: Pixel to convert
// 	view:  The view to use for converting the point
//
// return The converted point, in "world" units
func (this *RenderWindow) MapPixelToCoords(pos Vector2i, view *View) (coords Vector2f) {
	coords.fromC(C.sfRenderWindow_mapPixelToCoords(this.cptr, pos.toC(), view.toCPtr()))
	return
}

// Convert a point from world coordinates to window coordinates
//
// This function finds the pixel of the render-window that matches
// the given 2D point. In other words, it goes through the same process
// as the graphics card, to compute the final position of a rendered point.
//
// Initially, both coordinate systems (world units and target pixels)
// match perfectly. But if you define a custom view or resize your
// render-window, this assertion is not true anymore, ie. a point
// located at (150, 75) in your 2D world may map to the pixel
// (10, 50) of your render-window -- if the view is translated by (140, 25).
//
// This version uses a custom view for calculations, see the other
// overload of the function if you want to use the current view of the
// render-window.
//
// 	point Point to convert
// 	view The view to use for converting the point
//
// return The converted point, in target coordinates (pixels)
func (this *RenderWindow) MapCoordsToPixel(pos Vector2f, view *View) (coords Vector2i) {
	coords.fromC(C.sfRenderWindow_mapCoordsToPixel(this.cptr, pos.toC(), view.toCPtr()))
	return
}

// Save the current OpenGL render states and matrices
//
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering. Combined with PopGLStates,
// it ensures that:
// SFML's internal states are not messed up by your OpenGL code
// your OpenGL states are not modified by a call to a SFML function
//
// Note that this function is quite expensive: it saves all the
// possible OpenGL states and matrices, even the ones you
// don't care about. Therefore it should be used wisely.
// It is provided for convenience, but the best results will
// be achieved if you handle OpenGL states yourself (because
// you know which states have really changed, and need to be
// saved and restored). Take a look at the ResetGLStates
// function if you do so.
func (this *RenderWindow) PushGLStates() {
	C.sfRenderWindow_pushGLStates(this.cptr)
}

// Restore the previously saved OpenGL render states and matrices
//
// See the description of pushGLStates to get a detailed
// description of these functions.
func (this *RenderWindow) PopGLStates() {
	C.sfRenderWindow_popGLStates(this.cptr)
}

// Reset the internal OpenGL states so that the target is ready for drawing
//
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering, if you choose not to use
// PushGLStates/PopGLStates. It makes sure that all OpenGL
// states needed by SFML are set, so that subsequent RenderWindow.Draw
// calls will work as expected.
func (this *RenderWindow) ResetGLStates() {
	C.sfRenderWindow_resetGLStates(this.cptr)
}

// Copy the current contents of a render window to an image
//
// his is a slow operation, whose main purpose is to make
// screenshots of the application. If you want to update an
// image with the contents of the window and then use it for
// drawing, you should rather use a sfTexture and its
// update(sfWindow*) function.
// You can also draw things directly to a texture with the
// sfRenderWindow class.
//
// return New image containing the captured contents
func (this *RenderWindow) Capture() *Image {
	return newImageFromPtr(C.sfRenderWindow_capture(this.cptr))
}
