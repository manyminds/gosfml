package GoSFML2

/*
 #include <SFML/Graphics.h>
 #include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
)

//MISSING: 	sfImage_createFromMemory
//			sfImage_createFromStream
//			sfImage_getPixelsPtr

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Image struct {
	cptr *C.sfImage
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewImageFromFile(file string) *Image {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	image := &Image{C.sfImage_createFromFile(cFile)}
	runtime.SetFinalizer(image, (*Image).Destroy)
	return image
}

func NewImage(width, height uint) *Image {
	image := &Image{C.sfImage_create(C.uint(width), C.uint(height))}
	runtime.SetFinalizer(image, (*Image).Destroy)
	return image
}

func NewImageFromColor(width, height uint, color Color) *Image {
	image := &Image{C.sfImage_createFromColor(C.uint(width), C.uint(height), color.toC())}
	runtime.SetFinalizer(image, (*Image).Destroy)
	return image
}

func NewImageFromPixels(width, height uint, data []byte) *Image {
	image := &Image{C.sfImage_createFromPixels(C.uint(width), C.uint(height), (*C.sfUint8)(&data[0]))}
	runtime.SetFinalizer(image, (*Image).Destroy)
	return image
}

func (this *Image) Copy() *Image {
	image := &Image{C.sfImage_copy(this.cptr)}
	runtime.SetFinalizer(image, (*Image).Destroy)
	return image
}

func (this *Image) Destroy() {
	C.sfImage_destroy(this.cptr)
	this.cptr = nil
}

func (this *Image) SaveToFile(file string) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	C.sfImage_saveToFile(this.cptr, cFile)
}

func (this *Image) GetSize() (size Vector2u) {
	size.fromC(C.sfImage_getSize(this.cptr))
	return
}

//func (this *Image) CreateMaskFromColor(color Color, alpha byte) {
//	C.sfImage_createMaskFromColor(this.cptr, color.toC(), C.sfUint8(alpha))
//}

func (this *Image) CopyImage(source *Image, destX, destY uint, sourceRect Recti, applyAlpha bool) {
	C.sfImage_copyImage(this.cptr, source.cptr, C.uint(destX), C.uint(destY), sourceRect.toC(), goBool2C(applyAlpha))
}

func (this *Image) SetPixel(x, y uint, color Color) {
	C.sfImage_setPixel(this.cptr, C.uint(x), C.uint(y), color.toC())
}

func (this *Image) GetPixel(x, y uint) (color Color) {
	color.fromC(C.sfImage_getPixel(this.cptr, C.uint(x), C.uint(y)))
	return
}

func (this *Image) FlipHorizontally() {
	C.sfImage_flipHorizontally(this.cptr)
}

func (this *Image) FlipVertically() {
	C.sfImage_flipVertically(this.cptr)
}