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

// #include <SFML/Graphics/Texture.h> 
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Texture struct {
	cptr *C.sfTexture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewTextureFromFile(file string) (texture *Texture, err error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	texture = &Texture{C.sfTexture_createFromFile(cFile, nil)}
	runtime.SetFinalizer(texture, (*Texture).Destroy)

	if texture.cptr == nil {
		err = errors.New("NewTextureFromFile: Cannot load texture " + file)
	}

	return
}

func NewTextureFromMemory(data []byte, area *IntRect) (texture *Texture, err error) {
	if len(data) > 0 {
		texture = &Texture{C.sfTexture_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)), area.toCPtr())}
		runtime.SetFinalizer(texture, (*Texture).Destroy)
	}
	err = errors.New("NewTextureFromMemory: no data")
	return
}

func NewTextureFromImage(image *Image, area *IntRect) (texture *Texture, err error) {
	texture = &Texture{C.sfTexture_createFromImage(image.toCPtr(), area.toCPtr())}
	runtime.SetFinalizer(texture, (*Texture).Destroy)

	if texture.cptr == nil {
		err = errors.New("NewTextureFromFile: Cannot create texture from image")
	}

	return
}

func (this *Texture) Copy() *Texture {
	texture := &Texture{C.sfTexture_copy(this.cptr)}
	runtime.SetFinalizer(texture, (*Texture).Destroy)
	return texture
}

func (this *Texture) Destroy() {
	C.sfTexture_destroy(this.cptr)
	this.cptr = nil
}

func (this *Texture) GetSize() (size Vector2u) {
	size.fromC(C.sfTexture_getSize(this.cptr))
	return
}

func (this *Texture) UpdateFromWindow(window *Window, x, y uint) {
	C.sfTexture_updateFromWindow(this.cptr, window.cptr, C.uint(x), C.uint(y))
}

func (this *Texture) UpdateFromRenderWindow(window *RenderWindow, x, y uint) {
	C.sfTexture_updateFromRenderWindow(this.cptr, window.cptr, C.uint(x), C.uint(y))
}

func (this *Texture) UpdateFromImage(image *Image, x, y uint) {
	C.sfTexture_updateFromImage(this.cptr, image.toCPtr(), C.uint(x), C.uint(y))
}

func (this *Texture) SetSmooth(smooth bool) {
	C.sfTexture_setSmooth(this.cptr, goBool2C(smooth))
}

func (this *Texture) IsSmooth() bool {
	return sfBool2Go(C.sfTexture_isSmooth(this.cptr))
}

func (this *Texture) SetRepeated(repeated bool) {
	C.sfTexture_setRepeated(this.cptr, goBool2C(repeated))
}

func (this *Texture) IsRepeated() bool {
	return sfBool2Go(C.sfTexture_isRepeated(this.cptr))
}

func TextureGetMaximumSize() uint {
	return uint(C.sfTexture_getMaximumSize())
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Texture) toCPtr() *C.sfTexture {
	if this != nil {
		return this.cptr
	}
	return nil
}
