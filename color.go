// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Color.h>
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

// RGBA Color
//
// Color{} represents Color{0,0,0,0} i.e. transparent
type Color struct {
	//Created by cgo -godefs - DO NOT EDIT
	R uint8 //<< Red component
	G uint8 //<< Green component
	B uint8 //<< Blue component
	A uint8 //<< Alpha component (0 = transparent)
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

// Component-wise saturated addition of the two colors
func (this Color) Add(other Color) (newColor Color) {
	newColor.fromC(C.sfColor_add(this.toC(), other.toC()))
	return
}

// Component-wise multiplication of the two colors
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
