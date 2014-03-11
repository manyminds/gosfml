// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Window/Joystick.h>
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	JoystickX    = iota ///< The X axis
	JoystickY           ///< The Y axis
	JoystickZ           ///< The Z axis
	JoystickR           ///< The R axis
	JoystickU           ///< The U axis
	JoystickV           ///< The V axis
	JoystickPovX        ///< The X axis of the point-of-view hat
	JoystickPovY        ///< The Y axis of the point-of-view hat
)

const (
	JoystickCount       = C.sfJoystickCount       ///< Maximum number of supported joysticks
	JoystickButtonCount = C.sfJoystickButtonCount ///< Maximum number of supported buttons
	JoystickAxisCount   = C.sfJoystickAxisCount   ///< Maximum number of supported axes
)

type JoystickAxis int

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

// Check if a joystick is connected
//
// joystick: Index of the joystick to check
func JoystickIsConnected(joystick uint) bool {
	return sfBool2Go(C.sfJoystick_isConnected(C.uint(joystick)))
}

// Return the number of buttons supported by a joystick
//
// If the joystick is not connected, this function returns 0.
//
// joystick: Index of the joystick
func JoystickGetButtonCount(joystick uint) uint {
	return uint(C.sfJoystick_getButtonCount(C.uint(joystick)))
}

// Check if a joystick supports a given axis
//
// If the joystick is not connected, this function returns false.
//
// 	joystick: Index of the joystick
// 	axis:     Axis to check
func JoystickHasAxis(joystick uint, axis JoystickAxis) bool {
	return sfBool2Go(C.sfJoystick_hasAxis(C.uint(joystick), C.sfJoystickAxis(axis)))
}

// Check if a joystick button is pressed
//
// If the joystick is not connected, this function returns false.
//
// 	joystick: Index of the joystick
// 	button:   Button to check
func JoystickIsButtonPressed(joystick uint, button uint) bool {
	return sfBool2Go(C.sfJoystick_isButtonPressed(C.uint(joystick), C.uint(button)))
}

// Get the current position of a joystick axis
//
// If the joystick is not connected, this function returns 0.
//
// 	joystick: Index of the joystick
// 	axis:     Axis to check
//
// return Current position of the axis, in range [-100 .. 100]
func JoystickGetAxisPosition(joystick uint, axis JoystickAxis) float32 {
	return float32(C.sfJoystick_getAxisPosition(C.uint(joystick), C.sfJoystickAxis(axis)))
}

// Update the states of all joysticks
//
// This function is used internally by SFML, so you normally
// don't have to call it explicitely. However, you may need to
// call it if you have no window yet (or no window at all):
// in this case the joysticks states are not updated automatically.
func JoystickUpdate() {
	C.sfJoystick_update()
}
