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
import "C"

import (
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Transform [9]float32

var Transform_Identity Transform = Transform{1, 0, 0, 0, 1, 0, 0, 0, 1}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func (this *Transform) GetMatrix() (matrix [14]float32) {
	C.sfTransform_getMatrix(this.toCPtr(), (*C.float)(unsafe.Pointer(&matrix)))
	return
}

func (this *Transform) GetInverse() (inverse Transform, exists bool) {
	inv := C.sfTransform_getInverse(this.toCPtr())
	inverse.fromC(inv)
	exists = (inverse != Transform_Identity)
	return
}

func (this *Transform) TransformPoint(point Vector2f) (tansPoint Vector2f) {
	vec := C.sfTransform_transformPoint(this.toCPtr(), point.toC())
	tansPoint.fromC(vec)
	return
}

func (this *Transform) TransformRect(rect Rectf) (tansRect Rectf) {
	rec := C.sfTransform_transformRect(this.toCPtr(), rect.toC())
	tansRect.fromC(rec)
	return
}

func (this *Transform) Combine(other *Transform) (newTrans *Transform) {
	C.sfTransform_combine(this.toCPtr(), other.toCPtr())
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
