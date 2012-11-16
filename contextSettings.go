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

// #include <SFML/Window/Window.h>
// int getSizeContextSettings() { return sizeof(sfContextSettings); }
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
///		GO <-> C
/////////////////////////////////////

func (this *ContextSettings) fromC(csettings C.sfContextSettings) {
	this.DepthBits = uint(csettings.depthBits)
	this.StencilBits = uint(csettings.stencilBits)
	this.AntialiasingLevel = uint(csettings.antialiasingLevel)
	this.MajorVersion = uint(csettings.majorVersion)
	this.MinorVersion = uint(csettings.minorVersion)
}

func (this *ContextSettings) toC() C.sfContextSettings {
	return C.sfContextSettings{
		depthBits:         C.uint(this.DepthBits),
		stencilBits:       C.uint(this.StencilBits),
		antialiasingLevel: C.uint(this.AntialiasingLevel),
		majorVersion:      C.uint(this.MajorVersion),
		minorVersion:      C.uint(this.MinorVersion)}
}

/////////////////////////////////////
///		Testing
/////////////////////////////////////

func sizeofContextSettings() int {
	return int(C.getSizeContextSettings())
}
