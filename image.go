// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Image.h>
// #include <stdlib.h>
// sfUint8 sfImage_getPixelsPtrValue(const sfImage* image, int index) { return sfImage_getPixelsPtr(image)[index]; }
import "C"

import (
	"errors"
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
	runtime.SetFinalizer(image, (*Image).destroy)
	return image
}

// Create an image from a file on disk
//
// The supported image formats are bmp, png, tga, jpg, gif,
// psd, hdr and pic. Some format options are not supported,
// like progressive jpeg.
//
// file: Path of the image file to load
func NewImageFromFile(file string) (*Image, error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	if cptr := C.sfImage_createFromFile(cFile); cptr != nil {
		image := &Image{cptr}
		runtime.SetFinalizer(image, (*Image).destroy)
		return image, nil
	}

	return nil, genericError
}

// Create an image
//
// This image is filled with black pixels.
//
// 	width:  Width of the image
// 	height: Height of the image
func NewImage(width, height uint) (*Image, error) {
	if cptr := C.sfImage_create(C.uint(width), C.uint(height)); cptr != nil {
		image := &Image{cptr}
		runtime.SetFinalizer(image, (*Image).destroy)
		return image, nil
	}

	return nil, genericError
}

// Create an image and fill it with a unique color
//
// 	width:  Width of the image
// 	height: Height of the image
// 	color:  Fill color
func NewImageFromColor(width, height uint, color Color) (*Image, error) {
	if cptr := C.sfImage_createFromColor(C.uint(width), C.uint(height), color.toC()); cptr != nil {
		image := &Image{cptr}
		runtime.SetFinalizer(image, (*Image).destroy)
		return image, nil
	}

	return nil, genericError
}

// Create an image from an array of pixels
//
// The pixel array is assumed to contain 32-bits RGBA pixels,
// and have the given width and height. If not, this is
// an undefined behaviour.
// If pixels is nil, an empty image is created.
//
// 	width:  Width of the image
// 	height: Height of the image
// 	pixels: Slice of pixels to copy to the image
func NewImageFromPixels(width, height uint, data []byte) (*Image, error) {
	if len(data) == 0 {
		return nil, errors.New("NewImageFromPixels: len(data)==0")
	}

	if cptr := C.sfImage_createFromPixels(C.uint(width), C.uint(height), (*C.sfUint8)(&data[0])); cptr != nil {
		image := &Image{cptr}
		runtime.SetFinalizer(image, (*Image).destroy)
		return image, nil
	}

	return nil, genericError
}

// Create an image from a file in memory
//
// The supported image formats are bmp, png, tga, jpg, gif,
// psd, hdr and pic. Some format options are not supported,
// like progressive jpeg.
//
// 	data: Slice containing the file data
func NewImageFromMemory(data []byte) (*Image, error) {
	if len(data) == 0 {
		return nil, errors.New("NewImageFromMemory: len(data)==0")
	}

	if cptr := C.sfImage_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data))); cptr != nil {
		image := &Image{cptr}
		runtime.SetFinalizer(image, (*Image).destroy)
		return image, nil
	}
	return nil, genericError
}

// Copy an existing image
func (this *Image) Copy() *Image {
	image := &Image{C.sfImage_copy(this.cptr)}
	runtime.SetFinalizer(image, (*Image).destroy)
	return image
}

// Destroy an existing image
func (this *Image) destroy() {
	C.sfImage_destroy(this.cptr)
}

// Save an image to a file on disk
//
// The format of the image is automatically deduced from
// the extension. The supported image formats are bmp, png,
// tga and jpg. The destination file is overwritten
// if it already exists. This function fails if the image is empty.
//
// 	filename: Path of the file to save
func (this *Image) SaveToFile(file string) error {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	if !sfBool2Go(C.sfImage_saveToFile(this.cptr, cFile)) {
		return genericError
	}

	return nil
}

// Return the size of an image
func (this *Image) GetSize() (size Vector2u) {
	size.fromC(C.sfImage_getSize(this.cptr))
	return
}

// Create a transparency mask from a specified color-key
//
// This function sets the alpha value of every pixel matching
// the given color to alpha (0 by default), so that they
// become transparent.
//
// 	color: Color to make transparent
// 	alpha: Alpha value to assign to transparent pixels
func (this *Image) CreateMaskFromColor(color Color, alpha byte) {
	C.sfImage_createMaskFromColor(this.cptr, color.toC(), C.sfUint8(alpha))
}

// Copy pixels from an image onto another
//
// This function does a slow pixel copy and should not be
// used intensively. It can be used to prepare a complex
// static image from several others, but if you need this
// kind of feature in real-time you'd better use RenderTexture.
//
// If sourceRect is empty, the whole image is copied.
// If applyAlpha is set to true, the transparency of
// source pixels is applied. If it is false, the pixels are
// copied unchanged with their alpha value.
//
// 	source:     Source image to copy
// 	destX:      X coordinate of the destination position
// 	destY:      Y coordinate of the destination position
// 	sourceRect: Sub-rectangle of the source image to copy
// 	applyAlpha: Should the copy take in account the source transparency?
func (this *Image) CopyImage(source *Image, destX, destY uint, sourceRect IntRect, applyAlpha bool) {
	C.sfImage_copyImage(this.cptr, source.cptr, C.uint(destX), C.uint(destY), sourceRect.toC(), goBool2C(applyAlpha))
}

// Change the color of a pixel in an image
//
// This function doesn't check the validity of the pixel
// coordinates, using out-of-range values will result in
// an undefined behaviour.
//
// 	x:     X coordinate of pixel to change
// 	y:     Y coordinate of pixel to change
// 	color: New color of the pixel
func (this *Image) SetPixel(x, y uint, color Color) {
	C.sfImage_setPixel(this.cptr, C.uint(x), C.uint(y), color.toC())
}

// Get the color of a pixel in an image
//
// This function doesn't check the validity of the pixel
// coordinates, using out-of-range values will result in
// an undefined behaviour.
//
// 	x:     X coordinate of pixel to get
// 	y:     Y coordinate of pixel to get
func (this *Image) GetPixel(x, y uint) (color Color) {
	color.fromC(C.sfImage_getPixel(this.cptr, C.uint(x), C.uint(y)))
	return
}

// Get a slice of pixels of an image
//
// The length of the slice is width * height * 4 (RGBA).
func (this *Image) GetPixelData() []byte {
	data := make([]byte, this.GetSize().X*this.GetSize().Y*4)
	for i := 0; i < len(data); i++ {
		data[i] = byte(C.sfImage_getPixelsPtrValue(this.cptr, C.int(i)))
	}
	return data
}

// Flip an image horizontally (left <-> right)
func (this *Image) FlipHorizontally() {
	C.sfImage_flipHorizontally(this.cptr)
}

// Flip an image vertically (top <-> bottom)
func (this *Image) FlipVertically() {
	C.sfImage_flipVertically(this.cptr)
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Image) toCPtr() *C.sfImage {
	if this != nil {
		return this.cptr
	}
	return nil
}
