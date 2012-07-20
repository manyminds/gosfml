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

// #include <SFML/Graphics.h>
// int getSizeColor() { return sizeof(sfColor); }
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Color struct {
	R byte
	G byte
	B byte
	A byte //< 0=transparent
}

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

var (
	Color_Black = Color{0, 0, 0, 255}
	Color_Red   = Color{255, 0, 0, 255}
	Color_Green = Color{0, 255, 0, 255}
	Color_Blue  = Color{0, 0, 255, 255}
	Color_White = Color{255, 255, 255, 255}
)

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Color) fromC(color C.sfColor) {
	this.R = byte(color.r)
	this.G = byte(color.g)
	this.B = byte(color.b)
	this.A = byte(color.a)
}

func (this *Color) toC() C.sfColor {
	return C.sfColor{r: C.sfUint8(this.R), g: C.sfUint8(this.G), b: C.sfUint8(this.B), a: C.sfUint8(this.A)}
}

/////////////////////////////////////
///		Testing
/////////////////////////////////////

func sizeofColor() int {
	return int(C.getSizeColor())
}
