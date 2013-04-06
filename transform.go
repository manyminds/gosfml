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

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Return the 4x4 matrix of a transform
//
// This function fills an array of 16 floats with the transform
// converted as a 4x4 matrix, which is directly compatible with
// OpenGL functions.
func (this *Transform) GetMatrix() (matrix [14]float32) {
	C.sfTransform_getMatrix(this.toCPtr(), (*C.float)(unsafe.Pointer(&matrix)))
	return
}

// Return the inverse of a transform
//
// If the inverse cannot be computed, a new identity transform
// is returned.
func (this *Transform) GetInverse() (inverse Transform, exists bool) {
	inv := C.sfTransform_getInverse(this.toCPtr())
	inverse.fromC(inv)
	exists = (inverse != TransformIdentity())
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
func (this *Transform) Combine(other *Transform) (newTrans *Transform) {
	C.sfTransform_combine(this.toCPtr(), other.toCPtr())
	newTrans = this
	return
}

// Combine a transform with a translation
//
// 	x: Offset to apply on X axis
// 	y: Offset to apply on Y axis
func (this *Transform) Translate(x, y float32) (newTrans *Transform) {
	C.sfTransform_translate(this.toCPtr(), C.float(x), C.float(y))
	newTrans = this
	return
}

// Combine the current transform with a rotation
//
// 	angle: Rotation angle, in degrees
func (this *Transform) Rotate(angle float32) (newTrans *Transform) {
	C.sfTransform_rotate(this.toCPtr(), C.float(angle))
	newTrans = this
	return
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
func (this *Transform) RotateWithCenter(angle, centerX, centerY float32) (newTrans *Transform) {
	C.sfTransform_rotateWithCenter(this.toCPtr(), C.float(angle), C.float(centerX), C.float(centerY))
	newTrans = this
	return
}

// Combine the current transform with a scaling
//
// 	scaleX: Scaling factor on the X axis
// 	scaleY: Scaling factor on the Y axis
func (this *Transform) Scale(scaleX, scaleY float32) (newTrans *Transform) {
	C.sfTransform_scale(this.toCPtr(), C.float(scaleX), C.float(scaleY))
	newTrans = this
	return
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
