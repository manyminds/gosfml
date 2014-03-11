// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Glyph.h>
import "C"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Glyph struct {
	Advance     int     ///< Offset to move horizontically to the next character
	Bounds      IntRect ///< Bounding rectangle of the glyph, in coordinates relative to the baseline
	TextureRect IntRect ///< Texture coordinates of the glyph inside the font's image
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Glyph) fromC(glyph C.sfGlyph) {
	this.Advance = int(glyph.advance)
	this.Bounds.fromC(glyph.bounds)
	this.TextureRect.fromC(glyph.textureRect)
}
