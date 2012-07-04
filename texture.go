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

func (this *Texture) Destroy() {
	C.sfTexture_destroy(this.cptr)
	this.cptr = nil
}
