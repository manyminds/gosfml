/*
	COMPLETE: YES (4.7.2012)
*/

package GoSFML2

/*
 #include <SFML/Window.h>
*/
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Joystick_X    = iota ///< The X axis
	Joystick_Y           ///< The Y axis
	Joystick_Z           ///< The Z axis
	Joystick_R           ///< The R axis
	Joystick_U           ///< The U axis
	Joystick_V           ///< The V axis
	Joystick_PovX        ///< The X axis of the point-of-view hat
	Joystick_PovY        ///< The Y axis of the point-of-view hat
)

const (
	Joystick_Count       = 8  ///< Maximum number of supported joysticks
	Joystick_ButtonCount = 32 ///< Maximum number of supported buttons
	Joystick_AxisCount   = 8  ///< Maximum number of supported axes
)

type JoystickAxis int

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

func Joystick_IsConnected(joystick uint) bool {
	return sfBool2Go(C.sfJoystick_isConnected(C.uint(joystick)))
}

func Joystick_GetButtonCount(joystick uint) uint {
	return uint(C.sfJoystick_getButtonCount(C.uint(joystick)))
}

func Joystick_HasAxis(joystick uint, axis JoystickAxis) bool {
	return sfBool2Go(C.sfJoystick_hasAxis(C.uint(joystick), C.sfJoystickAxis(axis)))
}

func Joystick_IsButtonPressed(joystick uint, button uint) bool {
	return sfBool2Go(C.sfJoystick_isButtonPressed(C.uint(joystick), C.uint(button)))
}

func Joystick_GetAxisPosition(joystick uint, axis JoystickAxis) float32 {
	return float32(C.sfJoystick_getAxisPosition(C.uint(joystick), C.sfJoystickAxis(axis)))
}

func Joystick_Update() {
	C.sfJoystick_update()
}
