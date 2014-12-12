// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/RenderStates.h>
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

// Blend equations
const (
	EquationAdd      BlendEquation = C.sfBlendEquationAdd
	EquationSubtract BlendEquation = C.sfBlendEquationSubtract
)

type BlendEquation C.sfBlendEquation

// Blend factors
const (
	FactorZero             BlendFactor = C.sfBlendFactorZero             ///< (0, 0, 0, 0)
	FactorOne              BlendFactor = C.sfBlendFactorOne              ///< (1, 1, 1, 1)
	FactorSrcColor         BlendFactor = C.sfBlendFactorSrcColor         ///< (src.r, src.g, src.b, src.a)
	FactorOneMinusSrcColor BlendFactor = C.sfBlendFactorOneMinusSrcColor ///< (1, 1, 1, 1) - (src.r, src.g, src.b, src.a)
	FactorDstColor         BlendFactor = C.sfBlendFactorDstColor         ///< (dst.r, dst.g, dst.b, dst.a)
	FactorOneMinusDstColor BlendFactor = C.sfBlendFactorOneMinusDstColor ///< (1, 1, 1, 1) - (dst.r, dst.g, dst.b, dst.a)
	FactorSrcAlpha         BlendFactor = C.sfBlendFactorSrcAlpha         ///< (src.a, src.a, src.a, src.a)
	FactorOneMinusSrcAlpha BlendFactor = C.sfBlendFactorOneMinusSrcAlpha ///< (1, 1, 1, 1) - (src.a, src.a, src.a, src.a)
	FactorDstAlpha         BlendFactor = C.sfBlendFactorDstAlpha         ///< (dst.a, dst.a, dst.a, dst.a)
	FactorOneMinusDstAlpha BlendFactor = C.sfBlendFactorOneMinusDstAlpha ///< (1, 1, 1, 1) - (dst.a, dst.a, dst.a, dst.a)
)

type BlendFactor C.sfBlendFactor

var (
	// Pixel = Src * a + Dest * (1 - a)
	BlendAlpha = BlendMode{
		ColorSrcFactor: FactorSrcAlpha,
		ColorDstFactor: FactorOneMinusSrcAlpha,
		ColorEquation:  EquationAdd,
		AlphaSrcFactor: FactorOne,
		AlphaDstFactor: FactorOneMinusSrcAlpha,
		AlphaEquation:  EquationAdd,
	}

	// Pixel = Src + Dest
	BlendAdd = BlendMode{
		ColorSrcFactor: FactorSrcAlpha,
		ColorDstFactor: FactorOne,
		ColorEquation:  EquationAdd,
		AlphaSrcFactor: FactorOne,
		AlphaDstFactor: FactorOne,
		AlphaEquation:  EquationAdd,
	}

	// Pixel = Src * Dest
	BlendMultiply = BlendMode{
		ColorSrcFactor: FactorDstColor,
		ColorDstFactor: FactorZero,
		ColorEquation:  EquationAdd,
		AlphaSrcFactor: FactorDstColor,
		AlphaDstFactor: FactorZero,
		AlphaEquation:  EquationAdd,
	}

	// No blending
	BlendNone = BlendMode{
		ColorSrcFactor: FactorOne,
		ColorDstFactor: FactorZero,
		ColorEquation:  EquationAdd,
		AlphaSrcFactor: FactorOne,
		AlphaDstFactor: FactorZero,
		AlphaEquation:  EquationAdd,
	}
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type BlendMode struct {
	ColorSrcFactor BlendFactor   ///< Source blending factor for the color channels
	ColorDstFactor BlendFactor   ///< Destination blending factor for the color channels
	ColorEquation  BlendEquation ///< Blending equation for the color channels
	AlphaSrcFactor BlendFactor   ///< Source blending factor for the alpha channel
	AlphaDstFactor BlendFactor   ///< Destination blending factor for the alpha channel
	AlphaEquation  BlendEquation ///< Blending equation for the alpha channel
}

type RenderStates struct {
	BlendMode BlendMode
	Transform Transform
	Shader    *Shader
	Texture   *Texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func DefaultRenderStates() RenderStates {
	return RenderStates{Shader: nil, Texture: nil, BlendMode: BlendAlpha, Transform: TransformIdentity()}
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *BlendMode) toC() C.sfBlendMode {
	return C.sfBlendMode{
		colorSrcFactor: C.sfBlendFactor(this.ColorSrcFactor),
		colorDstFactor: C.sfBlendFactor(this.ColorDstFactor),
		colorEquation:  C.sfBlendEquation(this.ColorEquation),
		alphaSrcFactor: C.sfBlendFactor(this.AlphaSrcFactor),
		alphaDstFactor: C.sfBlendFactor(this.AlphaDstFactor),
		alphaEquation:  C.sfBlendEquation(this.AlphaEquation),
	}
}

func (this *RenderStates) toC() C.sfRenderStates {
	return C.sfRenderStates{blendMode: this.BlendMode.toC(), transform: this.Transform.toC(), texture: this.Texture.toCPtr(), shader: this.Shader.toCPtr()}
}
