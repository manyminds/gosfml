package GoSFML2

/*
 #include <SFML/Graphics.h> 
 #include <stdlib.h>
*/
import "C"
import "runtime"
import "unsafe"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Font struct {
	cptr *C.sfFont
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewFontFromFile(filename string) *Font {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	font := &Font{C.sfFont_createFromFile(cFilename)}
	runtime.SetFinalizer(font, (*Font).Destroy)
	return font
}

func NewFontFromMemory(data []byte) *Font {
	//not implemented
	return nil
}

func NewFontFromStream() *Font {
	//not implemented
	return nil
}

func (this *Font) Copy() *Font {
	return &Font{C.sfFont_copy(this.cptr)}
}

func (this *Font) Destroy() {
	C.sfFont_destroy(this.cptr)
	this.cptr = nil
}

func (this *Font) GetGlyph(codePoint uint, characterSize uint32, bold bool) (glyph Glyph) {
	glyph.fromC(C.sfFont_getGlyph(this.cptr, C.sfUint32(codePoint), C.uint(characterSize), goBool2C(bold)))
	return
}

func (this *Font) GetKerning(first uint32, second uint32, characterSize uint) int {
	return int(C.sfFont_getKerning(this.cptr, C.sfUint32(first), C.sfUint32(second), C.uint(characterSize)))
}

func (this *Font) GetLineSpacing(characterSize uint) int {
	return int(C.sfFont_getLineSpacing(this.cptr, C.uint(characterSize)))
}

func (this *Font) GetTexture(characterSize uint) Texture {
	return Texture{C.sfFont_getTexture(this.cptr, C.uint(characterSize))}
}

func GetDefaultFont() *Font {
	return &Font{C.sfFont_getDefaultFont()}
}
