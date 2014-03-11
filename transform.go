// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Transform.h>
import "C"

import (
	"unsafe"
)

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

type Transformer interface {
	SetPosition(Vector2f)
	SetScale(Vector2f)
	SetRotation(float32)
	SetOrigin(Vector2f)

	GetRotation() float32
	GetPosition() Vector2f
	GetScale() Vector2f
	GetOrigin() Vector2f

	Move(Vector2f)
	Scale(Vector2f)
	Rotate(float32)

	GetTransform() Transform
	GetInverseTransform() Transform
}

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

func TransformIdentity() Transform { return Transform{1, 0, 0, 0, 1, 0, 0, 0, 1} }

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Transform [9]float32
type Matrix [16]float32 // 4x4 matrix

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Return the 4x4 matrix of a transform
//
// This function fills an array of 16 floats with the transform
// converted as a 4x4 matrix, which is directly compatible with
// OpenGL functions.
func (this *Transform) GetMatrix() (matrix Matrix) {
	C.sfTransform_getMatrix(this.toCPtr(), (*C.float)(unsafe.Pointer(&matrix)))
	return
}

// Return the inverse of a transform
//
// If the inverse cannot be computed, a new identity transform
// is returned.
func (this *Transform) GetInverse() (inverse Transform) {
	inverse.fromC(C.sfTransform_getInverse(this.toCPtr()))
	return
}

// Apply a transform to a 2D point
//
// 	point: Point to transform
func (this *Transform) TransformPoint(point Vector2f) (tansPoint Vector2f) {
	vec := C.sfTransform_transformPoint(this.toCPtr(), point.toC())
	tansPoint.fromC(vec)
	return
}

// Apply a transform to a rectangle
//
// Since SFML doesn't provide support for oriented rectangles,
// the result of this function is always an axis-aligned
// rectangle. Which means that if the transform contains a
// rotation, the bounding rectangle of the transformed rectangle
// is returned.
//
// 	rect: Rectangle to transform
func (this *Transform) TransformRect(rect FloatRect) (tansRect FloatRect) {
	rec := C.sfTransform_transformRect(this.toCPtr(), rect.toC())
	tansRect.fromC(rec)
	return
}

// Combine two transforms
//
// Mathematically, it is equivalent to a matrix multiplication.
func (this *Transform) Combine(other *Transform) *Transform {
	C.sfTransform_combine(this.toCPtr(), other.toCPtr())
	return this
}

// Combine a transform with a translation
//
// 	x: Offset to apply on X axis
// 	y: Offset to apply on Y axis
func (this *Transform) Translate(x, y float32) *Transform {
	C.sfTransform_translate(this.toCPtr(), C.float(x), C.float(y))
	return this
}

// Combine the current transform with a rotation
//
// 	angle: Rotation angle, in degrees
func (this *Transform) Rotate(angle float32) *Transform {
	C.sfTransform_rotate(this.toCPtr(), C.float(angle))
	return this
}

// Combine the current transform with a rotation
//
// The center of rotation is provided for convenience as a second
// argument, so that you can build rotations around arbitrary points
// more easily (and efficiently) than the usual
// [translate(-center), rotate(angle), translate(center)].
//
// 	angle:     Rotation angle, in degrees
// 	centerX:   X coordinate of the center of rotation
// 	centerY:   Y coordinate of the center of rotation
func (this *Transform) RotateWithCenter(angle, centerX, centerY float32) *Transform {
	C.sfTransform_rotateWithCenter(this.toCPtr(), C.float(angle), C.float(centerX), C.float(centerY))
	return this
}

// Combine the current transform with a scaling
//
// 	scaleX: Scaling factor on the X axis
// 	scaleY: Scaling factor on the Y axis
func (this *Transform) Scale(scaleX, scaleY float32) *Transform {
	C.sfTransform_scale(this.toCPtr(), C.float(scaleX), C.float(scaleY))
	return this
}

// Combine the current transform with a scaling
//
// The center of scaling is provided for convenience as a second
// argument, so that you can build scaling around arbitrary points
// more easily (and efficiently) than the usual
// [translate(-center), scale(factors), translate(center)]
//
// 	scaleX:    Scaling factor on X axis
// 	scaleY:    Scaling factor on Y axis
// 	centerX:   X coordinate of the center of scaling
// 	centerY:   Y coordinate of the center of scaling
func (this *Transform) ScaleWithCenter(scaleX, scaleY, centerX, centerY float32) *Transform {
	C.sfTransform_scaleWithCenter(this.toCPtr(), C.float(scaleX), C.float(scaleY), C.float(centerX), C.float(centerY))
	return this
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Transform) fromC(transform C.sfTransform) {
	for i := 0; i < 9; i++ {
		this[i] = float32(transform.matrix[i])
	}
}

func (this *Transform) toC() (transform C.sfTransform) {
	for i := 0; i < 9; i++ {
		transform.matrix[i] = C.float(this[i])
	}
	return
}

func (this *Transform) toCPtr() *C.sfTransform {
	return (*C.sfTransform)(unsafe.Pointer(this))
}

/////////////////////////////////////
///		TEST
/////////////////////////////////////

var _ Transformer = &Sprite{}
var _ Transformer = &Text{}
var _ Transformer = &RectangleShape{}
var _ Transformer = &CircleShape{}
var _ Transformer = &ConvexShape{}
