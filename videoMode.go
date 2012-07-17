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
 #include <stdlib.h>

 sfVideoMode videoModeAt(size_t index, sfVideoMode* modes) {
		return modes[index];
 }

*/
import "C"
import "unsafe"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type VideoMode struct {
	Width        uint ///< Video mode width, in pixels
	Height       uint ///< Video mode height, in pixels
	BitsPerPixel uint ///< Video mode pixel depth, in bits per pixels
}

type fullscreenModesCount struct {
	count uint
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func (this *VideoMode) SetAsDesktopVideoMode() {
	this.fromC(C.sfVideoMode_getDesktopMode())
}

func (this *VideoMode) IsValid() bool {
	return sfBool2Go(C.sfVideoMode_isValid(this.toC()))
}

func (this *VideoMode) GetFullscreenModes() []VideoMode {
	c := &fullscreenModesCount{}
	cVideoModes := C.sfVideoMode_getFullscreenModes((*C.size_t)(unsafe.Pointer(c)))

	modes := make([]VideoMode, c.count)
	for i := uint(0); i < c.count; i++ {
		modes[i].fromC(C.videoModeAt(C.size_t(i), cVideoModes))
	}
	return modes
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *VideoMode) fromC(videoMode C.sfVideoMode) {
	this.Width = uint(videoMode.width)
	this.Height = uint(videoMode.height)
	this.BitsPerPixel = uint(videoMode.bitsPerPixel)
}

func (this *VideoMode) toC() C.sfVideoMode {
	return C.sfVideoMode{width: C.uint(this.Width), height: C.uint(this.Height), bitsPerPixel: C.uint(this.BitsPerPixel)}
}
