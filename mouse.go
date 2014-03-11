// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Window/Mouse.h>
// #include <SFML/Window/Window.h>
// #include <SFML/Graphics/RenderWindow.h>
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	MouseLeft     = iota ///< The left mouse button
	MouseRight           ///< The right mouse button
	MouseMiddle          ///< The middle (wheel) mouse button
	MouseXButton1        ///< The first extra mouse button
	MouseXButton2        ///< The second extra mouse button

	MouseButtonCount ///< Keep last -- the total number of mouse buttons
)

type MouseButton int

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

// Check if a mouse button is pressed
//
// 	button: Button to check
func IsMouseButtonPressed(button MouseButton) bool {
	return sfBool2Go(C.sfMouse_isButtonPressed(C.sfMouseButton(button)))
}

// Set the current position of the mouse
//
// This function sets the current position of the mouse
// cursor relative to the given window, or desktop if nil is passed.
//
// 	position:   New position of the mouse
// 	relativeTo: Reference window
func MouseSetPosition(position Vector2i, relativeTo SystemWindow) {
	switch relativeTo.(type) {
	case *RenderWindow:
		C.sfMouse_setPositionRenderWindow(position.toC(), relativeTo.(*RenderWindow).cptr)
	case *Window:
		C.sfMouse_setPosition(position.toC(), relativeTo.(*Window).cptr)
	default:
	}
}

// Get the current position of the mouse
//
// This function returns the current position of the mouse
// cursor relative to the given window, or desktop if nil is passed.
//
// 	relativeTo: Reference window
func MouseGetPosition(relativeTo SystemWindow) (pos Vector2i) {
	switch relativeTo.(type) {
	case *RenderWindow:
		pos.fromC(C.sfMouse_getPositionRenderWindow(relativeTo.(*RenderWindow).cptr))
	case *Window:
		pos.fromC(C.sfMouse_getPosition(relativeTo.(*Window).cptr))
	default:
	}
	return
}
