// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Window/Window.h>
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type ContextSettings struct {
	DepthBits         uint ///< Bits of the depth buffer
	StencilBits       uint ///< Bits of the stencil buffer
	AntialiasingLevel uint ///< Level of antialiasing
	MajorVersion      uint ///< Major number of the context version to create
	MinorVersion      uint ///< Minor number of the context version to create
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func DefaultContextSettings() ContextSettings {
	return ContextSettings{DepthBits: 0, StencilBits: 0, AntialiasingLevel: 0, MajorVersion: 2, MinorVersion: 0}
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *ContextSettings) fromC(csettings C.sfContextSettings) {
	this.DepthBits = uint(csettings.depthBits)
	this.StencilBits = uint(csettings.stencilBits)
	this.AntialiasingLevel = uint(csettings.antialiasingLevel)
	this.MajorVersion = uint(csettings.majorVersion)
	this.MinorVersion = uint(csettings.minorVersion)
}

//allocates memory!
func (this *ContextSettings) toC() C.sfContextSettings {
	cs := C.sfContextSettings{}
	cs.depthBits = C.uint(this.DepthBits)
	cs.stencilBits = C.uint(this.StencilBits)
	cs.antialiasingLevel = C.uint(this.AntialiasingLevel)
	cs.majorVersion = C.uint(this.MajorVersion)
	cs.minorVersion = C.uint(this.MinorVersion)

	return cs
}
