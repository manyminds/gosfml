// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/View.h>
import "C"

import (
	"runtime"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type View struct {
	cptr *C.sfView
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a default view
//
// This function creates a default view of (0, 0, 1000, 1000)
func NewView() *View {
	view := &View{C.sfView_create()}
	runtime.SetFinalizer(view, (*View).destroy)
	return view
}

// Construct a view from a rectangle
//
// 	rect: Rectangle defining the zone to display
func NewViewFromRect(rect FloatRect) *View {
	view := &View{C.sfView_createFromRect(rect.toC())}
	runtime.SetFinalizer(view, (*View).destroy)
	return view
}

// Destroy an existing view
func (this *View) destroy() {
	C.sfView_destroy(this.cptr)
}

// Copy an existing view
func (this *View) Copy() *View {
	view := &View{C.sfView_copy(this.cptr)}
	runtime.SetFinalizer(view, (*View).destroy)
	return view
}

// Set the center of a view
//
// 	center: New center
func (this *View) SetCenter(center Vector2f) {
	C.sfView_setCenter(this.cptr, center.toC())
}

// Set the size of a view
//
// 	size: New size of the view
func (this *View) SetSize(size Vector2f) {
	C.sfView_setSize(this.cptr, size.toC())
}

// Set the orientation of a view
//
// The default rotation of a view is 0 degree.
//
// 	rotation: New angle, in degrees
func (this *View) SetRotation(rotation float32) {
	C.sfView_setRotation(this.cptr, C.float(rotation))
}

// Set the target viewport of a view
//
// The viewport is the rectangle into which the contents of the
// view are displayed, expressed as a factor (between 0 and 1)
// of the size of the render target to which the view is applied.
// For example, a view which takes the left side of the target would
// be defined by a rect of (0, 0, 0.5, 1).
// By default, a view has a viewport which covers the entire target.
//
// 	viewport: New viewport rectangle
func (this *View) SetViewport(viewport FloatRect) {
	C.sfView_setViewport(this.cptr, viewport.toC())
}

// Reset a view to the given rectangle
//
// Note that this function resets the rotation angle to 0.
//
// 	rect: Rectangle defining the zone to display
func (this *View) Reset(rect FloatRect) {
	C.sfView_reset(this.cptr, rect.toC())
}

// Get the center of a view
func (this *View) GetCenter() (center Vector2f) {
	center.fromC(C.sfView_getCenter(this.cptr))
	return
}

// Get the size of a view
func (this *View) GetSize() (size Vector2f) {
	size.fromC(C.sfView_getSize(this.cptr))
	return
}

// Get the current orientation of a view
func (this *View) GetRotation() float32 {
	return float32(C.sfView_getRotation(this.cptr))
}

// Get the target viewport rectangle of a view
func (this *View) GetViewport() (rect FloatRect) {
	rect.fromC(C.sfView_getViewport(this.cptr))
	return
}

// Move a view relatively to its current position
func (this *View) Move(offset Vector2f) {
	C.sfView_move(this.cptr, offset.toC())
}

// Rotate a view relatively to its current orientation
//
// 	angle: Angle to rotate, in degrees
func (this *View) Rotate(angle float32) {
	C.sfView_rotate(this.cptr, C.float(angle))
}

// Resize a view rectangle relatively to its current size
//
// Resizing the view simulates a zoom, as the zone displayed on
// screen grows or shrinks.
//
// factor is a multiplier:
// 	1 keeps the size unchanged
// 	> 1 makes the view bigger (objects appear smaller)
// 	< 1 makes the view smaller (objects appear bigger)
//
// 	factor: Zoom factor to apply
func (this *View) Zoom(factor float32) {
	C.sfView_zoom(this.cptr, C.float(factor))
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *View) toCPtr() *C.sfView {
	if this != nil {
		return this.cptr
	}
	return nil
}

func newViewFromPtr(cptr *C.sfView) *View {
	view := &View{C.sfView_copy(cptr)}
	runtime.SetFinalizer(view, (*View).destroy)
	return view
}
