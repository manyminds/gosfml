package GoSFML2

/*
 #include <SFML/Graphics.h> 
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
