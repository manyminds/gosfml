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

type Shader struct {
	cptr *C.sfShader
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateShaderFromFile(vertexShaderFile, fragmentShaderFile string) (shader *Shader, err error) {
	var (
		cVShader *C.char = nil
		cFShader *C.char = nil
	)

	if vertexShaderFile != "" {
		cVShader = C.CString(vertexShaderFile)
		defer C.free(unsafe.Pointer(cVShader))
	}

	if fragmentShaderFile != "" {
		cFShader = C.CString(fragmentShaderFile)
		defer C.free(unsafe.Pointer(cFShader))
	}

	shader = &Shader{C.sfShader_createFromFile(cVShader, cFShader)}
	runtime.SetFinalizer(shader, (*Shader).Destroy)
	
	//error check
	if shader.cptr == nil { err = &Error{"Cannot create Shader"} }
	
	return
}

func (this *Shader) Destroy() {
	C.sfShader_destroy(this.cptr)
	this.cptr = nil
}

func (this *Shader) SetFloatParameter(name string, x float32) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setFloatParameter(this.cptr, cname, C.float(x))
}

func (this *Shader) SetFloat2Parameter(name string, x, y float32) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setFloat2Parameter(this.cptr, cname, C.float(x), C.float(y))
}

func (this *Shader) SetFloat3Parameter(name string, x, y, z float32) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setFloat3Parameter(this.cptr, cname, C.float(x), C.float(y), C.float(z))
}

func (this *Shader) SetFloat4Parameter(name string, x, y, z, w float32) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setFloat4Parameter(this.cptr, cname, C.float(x), C.float(y), C.float(z), C.float(w))
}

func (this *Shader) SetColorParameter(name string, color Color) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setColorParameter(this.cptr, cname, color.toC())
}

func (this *Shader) SetTransformParameter(name string, trans Transform) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setTransformParameter(this.cptr, cname, trans.toC())
}

func (this *Shader) SetTextureParameter(name string, texture *Texture) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setTextureParameter(this.cptr, cname, texture.cptr)
}

func (this *Shader) SetCurrentTextureParameter(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setCurrentTextureParameter(this.cptr, cname)
}

func (this *Shader) Bind() {
	C.sfShader_bind(this.cptr)
}

func (this *Shader) Unbind() {
	C.sfShader_unbind(this.cptr)
}

func Shader_IsAvailable() bool {
	return sfBool2Go(C.sfShader_isAvailable())
}
