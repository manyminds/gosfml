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

// #include <SFML/Graphics/Shader.h> 
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

type Shader struct {
	cptr *C.sfShader
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewShaderFromFile(vertexShaderFile, fragmentShaderFile string) (shader *Shader, err error) {
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
	if shader.cptr == nil {
		err = errors.New("NewShaderFromFile: Cannot create Shader " + vertexShaderFile + " " + fragmentShaderFile)
	}

	return
}

func NewShaderFromMemory(vertexShader, fragmentShader string) (shader *Shader, err error) {
	var (
		cVShader *C.char = nil
		cFShader *C.char = nil
	)

	if vertexShader != "" {
		cVShader = C.CString(vertexShader)
		defer C.free(unsafe.Pointer(cVShader))
	}

	if fragmentShader != "" {
		cFShader = C.CString(fragmentShader)
		defer C.free(unsafe.Pointer(cFShader))
	}

	shader = &Shader{C.sfShader_createFromMemory(cVShader, cFShader)}
	runtime.SetFinalizer(shader, (*Shader).Destroy)

	//error check
	if shader.cptr == nil {
		err = errors.New("NewShaderFromFile: Cannot create Shader")
	}

	return
}

func (this *Shader) Destroy() {
	C.sfShader_destroy(this.toCPtr())
	this.cptr = nil
}

func (this *Shader) SetColorParameter(name string, color Color) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setColorParameter(this.toCPtr(), cname, color.toC())
}

func (this *Shader) SetTransformParameter(name string, trans Transform) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setTransformParameter(this.toCPtr(), cname, trans.toC())
}

func (this *Shader) SetTextureParameter(name string, texture *Texture) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setTextureParameter(this.toCPtr(), cname, texture.cptr)
}

func (this *Shader) SetCurrentTextureParameter(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setCurrentTextureParameter(this.toCPtr(), cname)
}

func (this *Shader) SetFloatParameter(name string, data ...float32) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	switch len(data) {
	case 1:
		C.sfShader_setFloatParameter(this.toCPtr(), cname, C.float(data[0]))
	case 2:
		C.sfShader_setFloat2Parameter(this.toCPtr(), cname, C.float(data[0]), C.float(data[1]))
	case 3:
		C.sfShader_setFloat3Parameter(this.toCPtr(), cname, C.float(data[0]), C.float(data[1]), C.float(data[2]))
	case 4:
		C.sfShader_setFloat4Parameter(this.toCPtr(), cname, C.float(data[0]), C.float(data[1]), C.float(data[2]), C.float(data[3]))
	default:
		panic("Shader.SetFloatParameter: Invalid amount of data.")
	}
}

func (this *Shader) Bind() {
	C.sfShader_bind(this.toCPtr())
}

func ShaderAvailable() bool {
	return sfBool2Go(C.sfShader_isAvailable())
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Shader) toCPtr() *C.sfShader {
	if this != nil {
		return this.cptr
	}
	return nil
}
