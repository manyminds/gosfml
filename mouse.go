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

func Mouse_IsButtonPressed(button MouseButton) bool {
	return sfBool2Go(C.sfMouse_isButtonPressed(C.sfMouseButton(button)))
}

func Mouse_GetPosition(relativeTo *Window) (pos Vector2i) {
	pos.fromC(C.sfMouse_getPosition(relativeTo.cptr))
	return
}

func Mouse_SetPosition(position Vector2i, relativeTo *Window) {
	C.sfMouse_setPosition(position.toC(), relativeTo.cptr)
}
