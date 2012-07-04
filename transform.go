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

var (
	Transform_Identity = Transform{1, 0, 0, 0, 1, 0, 0, 0, 1}
)

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func (this *Transform) GetMatrix() (matrix [14]float32) {
	C.sfTransform_getMatrix(this.toCPtr(), (*C.float)(unsafe.Pointer(&matrix)))
	return
}

func (this *Transform) GetInverse() (exists bool, inverse Transform) {
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
