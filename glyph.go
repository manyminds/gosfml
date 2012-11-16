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
 #include <SFML/Graphics/Glyph.h> 
*/
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Glyph struct {
	Advance     int   ///< Offset to move horizontically to the next character
	Bounds      Recti ///< Bounding rectangle of the glyph, in coordinates relative to the baseline
	TextureRect Recti ///< Texture coordinates of the glyph inside the font's image
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Glyph) fromC(glyph C.sfGlyph) {
	this.Advance = int(glyph.advance)
	this.Bounds.fromC(glyph.bounds)
	this.TextureRect.fromC(glyph.textureRect)
}
