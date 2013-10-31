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
	defaultRenderStates = RenderStates{Shader: nil, Texture: nil, BlendMode: BlendAlpha, Transform: TransformIdentity()}
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type BlendMode int

type RenderStates struct {
	BlendMode BlendMode
	Transform Transform
	Shader    *Shader
	Texture   *Texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func RenderStatesDefault() RenderStates {
	return defaultRenderStates
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *RenderStates) toC() C.sfRenderStates {
	return C.sfRenderStates{blendMode: C.sfBlendMode(this.BlendMode), transform: this.Transform.toC(), texture: this.Texture.toCPtr(), shader: this.Shader.toCPtr()}
}
