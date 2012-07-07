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

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Texture struct {
	cptr *C.sfTexture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateTextureFromFile(file string) *Texture {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	texture := &Texture{C.sfTexture_createFromFile(cFile, nil)}
	runtime.SetFinalizer(texture, (*Texture).Destroy)
	return texture
}

//needs testing
func CreateTextureFromMemory(data []byte) *Texture {
	texture := &Texture{C.sfImage_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))}
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

func Texture_GetMaximumSize() uint {
	return uint(C.sfTexture_getMaximumSize())
}
