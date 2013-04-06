// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, 
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
//		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. 	This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/Graphics/Color.h>
// int getSizeColor() { return sizeof(sfColor); }
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Color struct {
	R byte //<< Red component
	G byte //<< Green component
	B byte //<< Blue component
	A byte //<< Alpha component (0 = transparent)
}

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

func ColorBlack() Color       { return Color{0, 0, 0, 255} }
func ColorWhite() Color       { return Color{255, 255, 255, 255} }
func ColorRed() Color         { return Color{255, 0, 0, 255} }
func ColorGreen() Color       { return Color{0, 255, 0, 255} }
func ColorBlue() Color        { return Color{0, 0, 255, 255} }
func ColorYellow() Color      { return Color{255, 255, 0, 255} }
func ColorMagenta() Color     { return Color{255, 0, 255, 255} }
func ColorCyan() Color        { return Color{0, 255, 255, 255} }
func ColorTransparent() Color { return Color{0, 0, 0, 0} }

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

//Component-wise saturated addition of the two colors
func (this Color) Add(other Color) (newColor Color) {
	newColor.fromC(C.sfColor_add(this.toC(), other.toC()))
	return
}

//Component-wise multiplication of the two colors
func (this Color) Modulate(other Color) (newColor Color) {
	newColor.fromC(C.sfColor_modulate(this.toC(), other.toC()))
	return
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Color) fromC(color C.sfColor) {
	this.R = byte(color.r)
	this.G = byte(color.g)
	this.B = byte(color.b)
	this.A = byte(color.a)
}

func (this Color) toC() C.sfColor {
	return C.sfColor{r: C.sfUint8(this.R), g: C.sfUint8(this.G), b: C.sfUint8(this.B), a: C.sfUint8(this.A)}
}

/////////////////////////////////////
///		Testing
/////////////////////////////////////

func sizeofColor() int {
	return int(C.getSizeColor())
}
