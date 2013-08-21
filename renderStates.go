// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/RenderStates.h>
import "C"

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
///		VARS
/////////////////////////////////////

var (
	defaultRenderStates = MakeRenderStates(BlendAlpha, TransformIdentity(), nil, nil)
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
	return defaultRenderStates
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Initializes a RenderStates object.
func MakeRenderStates(blendMode BlendMode, transform Transform, texture *Texture, shader *Shader) (rt RenderStates) {
	rt.SetBlendMode(blendMode)
	rt.SetTransform(transform)
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
func (this *RenderStates) SetTransform(transform Transform) {
	this.cRenderStates.transform = transform.toC()
}

// Sets the blend mode of the RenderStates.
func (this *RenderStates) SetBlendMode(blendMode BlendMode) {
	this.cRenderStates.blendMode = C.sfBlendMode(blendMode)
}

// Gets the blend mode of the RenderStates.
func (this *RenderStates) GetBlendMode() BlendMode {
	return BlendMode(this.cRenderStates.blendMode)
}

// Gets the shader of the RenderStates
func (this *RenderStates) GetShader() *Shader {
	return this.shader
}

// Gets the texture of the RenderStates
func (this *RenderStates) GetTexture() *Texture {
	return this.texture
}

// Gets the transform of the RenderStates
func (this *RenderStates) GetTransform() (trans Transform) {
	trans.fromC(this.cRenderStates.transform)
	return
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *RenderStates) toCPtr() *C.sfRenderStates {
	return &this.cRenderStates
}
