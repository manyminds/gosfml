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

func JoystickIsConnected(joystick uint) bool {
	return sfBool2Go(C.sfJoystick_isConnected(C.uint(joystick)))
}

func JoystickGetButtonCount(joystick uint) uint {
	return uint(C.sfJoystick_getButtonCount(C.uint(joystick)))
}

func JoystickHasAxis(joystick uint, axis JoystickAxis) bool {
	return sfBool2Go(C.sfJoystick_hasAxis(C.uint(joystick), C.sfJoystickAxis(axis)))
}

func JoystickIsButtonPressed(joystick uint, button uint) bool {
	return sfBool2Go(C.sfJoystick_isButtonPressed(C.uint(joystick), C.uint(button)))
}

func JoystickGetAxisPosition(joystick uint, axis JoystickAxis) float32 {
	return float32(C.sfJoystick_getAxisPosition(C.uint(joystick), C.sfJoystickAxis(axis)))
}

func JoystickUpdate() {
	C.sfJoystick_update()
}
