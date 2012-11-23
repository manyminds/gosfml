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

package gosfml2

// #include <SFML/Graphics/Image.h>
// #include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Image struct {
	cptr *C.sfImage
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func newImageFromPtr(cptr *C.sfImage) *Image {
	image := &Image{cptr}
	runtime.SetFinalizer(image, (*Image).Destroy)
	return image
}

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

func NewImageFromMemory(data []byte) (*Image, error) {
	if len(data) > 0 {
		image := &Image{C.sfImage_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))}
		runtime.SetFinalizer(image, (*Image).Destroy)
		return image, nil
	}
	return nil, &Error{"NewImageFromMemory: no data"}
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

func (this *Image) CreateMaskFromColor(color Color, alpha byte) {
	C.sfImage_createMaskFromColor(this.cptr, color.toC(), C.sfUint8(alpha))
}

func (this *Image) CopyImage(source *Image, destX, destY uint, sourceRect IntRect, applyAlpha bool) {
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
