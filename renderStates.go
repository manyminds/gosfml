// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, 
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
//		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. 	This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/Graphics/RenderStates.h>
import "C"
import "unsafe"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	BlendAlpha    = iota ///< Pixel = Src * a + Dest * (1 - a)
	BlendAdd             ///< Pixel = Src + Dest
	BlendMultiply        ///< Pixel = Src * Dest
	BlendNone            ///< No blending
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type BlendMode int

type RenderStates struct {
	cRenderStates C.sfRenderStates
	shader        *Shader
	texture       *Texture
}

/////////////////////////////////////
///		CONTS
/////////////////////////////////////

func RenderStatesDefault() RenderStates {
	return MakeRenderStates(BlendAlpha, TransformIdentity(), nil, nil)
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Initializes a RenderStates object.
func MakeRenderStates(blendMode BlendMode, transform Transform, texture *Texture, shader *Shader) (rt RenderStates) {
	rt.SetBlendMode(blendMode)
	rt.SetTramsform(transform)
	rt.SetTexture(texture)
	rt.SetShader(shader)
	return
}

// Sets the shader of the RenderStates.
//
// 	shader: can be nil (no shader)
func (this *RenderStates) SetShader(shader *Shader) {
	this.cRenderStates.shader = shader.toCPtr()
	this.shader = shader
}

// Sets the texture of the RenderStates.
//
// 	texture: can be nil (no texture)
func (this *RenderStates) SetTexture(texture *Texture) {
	this.cRenderStates.texture = texture.toCPtr()
	this.texture = texture
}

// Sets the transformation of the RenderStates.
func (this *RenderStates) SetTramsform(transform Transform) {
	this.cRenderStates.transform = transform.toC()
}

// Sets the blend mode of the RenderStates.
func (this *RenderStates) SetBlendMode(blendMode BlendMode) {
	this.cRenderStates.blendMode = C.sfBlendMode(blendMode)
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *RenderStates) toCPtr() *C.sfRenderStates {
	return (*C.sfRenderStates)(unsafe.Pointer(&this.cRenderStates))
}
