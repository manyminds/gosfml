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
 #include <SFML/Audio.h> 
*/
import "C"

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func Listener_SetGlobalVolume(volume float32) {
	C.sfListener_setGlobalVolume(C.float(volume))
}

func Listener_GetGlobalVolume() float32 {
	return float32(C.sfListener_getGlobalVolume())
}

func Listener_SetPosition(pos Vector3f) {
	C.sfListener_setPosition(pos.toC())
}

func Listener_GetPosition() (pos Vector3f) {
	pos.fromC(C.sfListener_getPosition())
	return
}

func Listener_SetDirection(dir Vector3f) {
	C.sfListener_setPosition(dir.toC())
}

func Listener_GetDirection() (dir Vector3f) {
	dir.fromC(C.sfListener_getDirection())
	return
}
