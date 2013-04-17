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

// #include <SFML/Graphics/Rect.h>
// int getSizeIntRect() { return sizeof(sfIntRect); }
// int getSizeFloatRect() { return sizeof(sfFloatRect); }
import "C"
import "unsafe"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type FloatRect struct {
	Left   float32
	Top    float32
	Width  float32
	Height float32
}

type IntRect struct {
	Left   int
	Top    int
	Width  int
	Height int
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Check if a point is inside a rectangle's area
//
// 	x: X coordinate of the point to test
// 	y: Y coordinate of the point to test
func (this FloatRect) Contains(x, y float32) bool {
	return sfBool2Go(C.sfFloatRect_contains(this.toCPtr(), C.float(x), C.float(y)))
}

// Check if a point is inside a rectangle's area
//
// 	x: X coordinate of the point to test
// 	y: Y coordinate of the point to test
func (this IntRect) Contains(x, y int) bool {
	return sfBool2Go(C.sfIntRect_contains(this.toCPtr(), C.int(x), C.int(y)))
}

// Check intersection between two rectangles
//
// 	other: Rectangle to test against
// 	intersection: Overlapping rect
func (this FloatRect) Intersects(other FloatRect) (test bool, intersection FloatRect) {
	test = sfBool2Go(C.sfFloatRect_intersects(this.toCPtr(), other.toCPtr(), intersection.toCPtr()))
	return
}

// Check intersection between two rectangles
//
// 	other: Rectangle to test against
// 	intersection: Overlapping rect
func (this IntRect) Intersects(other IntRect) (test bool, intersection IntRect) {
	test = sfBool2Go(C.sfIntRect_intersects(this.toCPtr(), other.toCPtr(), intersection.toCPtr()))
	return
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *FloatRect) fromC(rect C.sfFloatRect) {
	this.Left = float32(rect.left)
	this.Top = float32(rect.top)
	this.Width = float32(rect.width)
	this.Height = float32(rect.height)
}

func (this *IntRect) fromC(rect C.sfIntRect) {
	this.Left = int(rect.left)
	this.Top = int(rect.top)
	this.Width = int(rect.width)
	this.Height = int(rect.height)
}

func (this *IntRect) toC() C.sfIntRect {
	return C.sfIntRect{left: C.int(this.Left), top: C.int(this.Top), width: C.int(this.Width), height: C.int(this.Height)}
}

func (this *FloatRect) toC() C.sfFloatRect {
	return C.sfFloatRect{left: C.float(this.Left), top: C.float(this.Top), width: C.float(this.Width), height: C.float(this.Height)}
}

func (this *IntRect) toCPtr() *C.sfIntRect {
	return (*C.sfIntRect)(unsafe.Pointer(this))
}

func (this *FloatRect) toCPtr() *C.sfFloatRect {
	return (*C.sfFloatRect)(unsafe.Pointer(this))
}

/////////////////////////////////////
///		Testing
/////////////////////////////////////

func sizeofIntRect() int {
	return int(C.getSizeIntRect())
}

func sizeofFloatRect() int {
	return int(C.getSizeFloatRect())
}
