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

/*
 #include <SFML/Graphics.h>
*/
import "C"
import "unsafe"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Blend_Alpha    = iota ///< Pixel = Src * a + Dest * (1 - a)
	Blend_Add             ///< Pixel = Src + Dest
	Blend_Multiply        ///< Pixel = Src * Dest
	Blend_None            ///< No blending
)

/////////////////////////////////////
///		VARS
/////////////////////////////////////

var RenderStates_Default = NewRenderStates(Blend_Alpha, Transform_Identity, nil, nil)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type BlendMode int

type RenderStates struct {
	cRenderStates C.sfRenderStates
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewRenderStates(blendMode BlendMode, transform Transform, texture *Texture, shader *Shader) (rt RenderStates) {
	rt.cRenderStates.blendMode = C.sfBlendMode(blendMode)
	rt.cRenderStates.transform = transform.toC()
	rt.cRenderStates.shader = shader.toCPtr()
	rt.cRenderStates.texture = texture.toCPtr()
	return
}

// shader can be nil (no shader)
func (this *RenderStates) SetShader(shader *Shader) {
	if shader == nil {
		this.cRenderStates.shader = nil
	} else {
		this.cRenderStates.shader = shader.cptr
	}
}

// texture can be nil (no texture)
func (this *RenderStates) SetTexture(texture *Texture) {
	if texture == nil {
		this.cRenderStates.texture = nil
	} else {
		this.cRenderStates.texture = texture.cptr
	}
}

func (this *RenderStates) SetTramsform(transform Transform) {
	this.cRenderStates.transform = transform.toC()
}

func (this *RenderStates) SetBlendMode(blendMode BlendMode) {
	this.cRenderStates.blendMode = C.sfBlendMode(blendMode)
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *RenderStates) toCPtr() *C.sfRenderStates {
	return (*C.sfRenderStates)(unsafe.Pointer(&this.cRenderStates))
}
