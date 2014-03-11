// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Rect.h>
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
