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
	if relativeTo != nil {
		C.sfMouse_setPosition(position.toC(), relativeTo.cptr)
	}
}
