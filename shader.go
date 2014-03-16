// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Shader.h>
// #include <stdlib.h>
import "C"

import (
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

// Load both the vertex and fragment shaders from files
//
// This function can load both the vertex and the fragment
// shaders, or only one of them: pass "" (empty string) if you don't want to load
// either the vertex shader or the fragment shader.
// The sources must be text files containing valid shaders
// in GLSL language. GLSL is a C-like language dedicated to
// OpenGL shaders; you'll probably need to read a good documentation
// for it before writing your own shaders.
//
// 	vertexShaderFile:   Path of the vertex shader file to load, or "" to skip this shader
// 	fragmentShaderFile: Path of the fragment shader file to load, or "" to skip this shader
func NewShaderFromFile(vertexShaderFile, fragmentShaderFile string) (*Shader, error) {
	var (
		cVShader *C.char = nil
		cFShader *C.char = nil
	)

	if len(vertexShaderFile) > 0 {
		cVShader = C.CString(vertexShaderFile)
		defer C.free(unsafe.Pointer(cVShader))
	}

	if len(fragmentShaderFile) > 0 {
		cFShader = C.CString(fragmentShaderFile)
		defer C.free(unsafe.Pointer(cFShader))
	}

	if cptr := C.sfShader_createFromFile(cVShader, cFShader); cptr != nil {
		shader := &Shader{cptr}
		runtime.SetFinalizer(shader, (*Shader).destroy)

		return shader, nil
	}

	return nil, genericError
}

// Load both the vertex and fragment shaders from source codes in memory
//
// This function can load both the vertex and the fragment
// shaders, or only one of them: pass "" (empty string) if you don't want to load
// either the vertex shader or the fragment shader.
// The sources must be valid shaders in GLSL language. GLSL is
// a C-like language dedicated to OpenGL shaders; you'll
// probably need to read a good documentation for it before
// writing your own shaders.
//
// 	vertexShader:   String containing the source code of the vertex shader, or "" to skip this shader
// 	fragmentShader: String containing the source code of the fragment shader, or "" to skip this shader
func NewShaderFromMemory(vertexShader, fragmentShader string) (*Shader, error) {
	var (
		cVShader *C.char = nil
		cFShader *C.char = nil
	)

	if len(vertexShader) > 0 {
		cVShader = C.CString(vertexShader)
		defer C.free(unsafe.Pointer(cVShader))
	}

	if len(fragmentShader) > 0 {
		cFShader = C.CString(fragmentShader)
		defer C.free(unsafe.Pointer(cFShader))
	}

	if cptr := C.sfShader_createFromMemory(cVShader, cFShader); cptr != nil {
		shader := &Shader{cptr}
		runtime.SetFinalizer(shader, (*Shader).destroy)
		return shader, nil
	}

	return nil, genericError
}

// Destroy an existing shader
func (this *Shader) destroy() {
	globalCtxSetActive(true)
	C.sfShader_destroy(this.toCPtr())
	globalCtxSetActive(false)
}

// Change a color parameter of a shader
//
// name is the name of the variable to change in the shader.
// The corresponding parameter in the shader must be a 4x1 vector
// (vec4 GLSL type).
//
// It is important to note that the components of the color are
// normalized before being passed to the shader. Therefore,
// they are converted from range [0 .. 255] to range [0 .. 1].
// For example, a sf::Color(255, 125, 0, 255) will be transformed
// to a vec4(1.0, 0.5, 0.0, 1.0) in the shader.
//
// 	name:   Name of the parameter in the shader
// 	color:  Color to assign
func (this *Shader) SetColorParameter(name string, color Color) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setColorParameter(this.toCPtr(), cname, color.toC())
}

// Change a matrix parameter of a shader
//
// name is the name of the variable to change in the shader.
// The corresponding parameter in the shader must be a 4x4 matrix
// (mat4 GLSL type).
//
// 	name:      Name of the parameter in the shader
// 	transform: Transform to assign
func (this *Shader) SetTransformParameter(name string, trans Transform) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setTransformParameter(this.toCPtr(), cname, trans.toC())
}

// Change a texture parameter of a shader
//
// name is the name of the variable to change in the shader.
// The corresponding parameter in the shader must be a 2D texture
// (sampler2D GLSL type).
//
// 	name:    Name of the texture in the shader
// 	texture: Texture to assign
func (this *Shader) SetTextureParameter(name string, texture *Texture) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setTextureParameter(this.toCPtr(), cname, texture.cptr)
}

// Change a texture parameter of a shader
//
// This function maps a shader texture variable to the
// texture of the object being drawn, which cannot be
// known in advance.
// The corresponding parameter in the shader must be a 2D texture
// (sampler2D GLSL type).
//
// 	name:   Name of the texture in the shader
func (this *Shader) SetCurrentTextureParameter(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	C.sfShader_setCurrentTextureParameter(this.toCPtr(), cname)
}

// Change a n-components vector parameter of a shader
//
// name is the name of the variable to change in the shader.
// The corresponding parameter in the shader must be a n x 1 vector with n = 1 ... 4.
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

// Bind a shader for rendering (activate it)
//
// This function is not part of the graphics API, it mustn't be
// used when drawing SFML entities. It must be used only if you
// mix sfShader with OpenGL code.
//
// 	shader: Shader to bind, can be nil to use no shader
func BindShader(shader *Shader) {
	C.sfShader_bind(shader.toCPtr())
}

// Tell whether or not the system supports shaders
//
// This function should always be called before using
// the shader features. If it returns false, then
// any attempt to use shaders will fail.
func ShadersAvailable() bool {
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
