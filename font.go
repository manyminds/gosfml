// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Font.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Font struct {
	cptr *C.sfFont
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Font constructor
// Creates a new font from file
//
// The supported font formats are: TrueType, Type 1, CFF,
// OpenType, SFNT, X11 PCF, Windows FNT, BDF, PFR and Type 42.
// Note that this function know nothing about the standard
// fonts installed on the user's system, thus you can't
// load them directly.
func NewFontFromFile(filename string) (*Font, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	if cptr := C.sfFont_createFromFile(cFilename); cptr != nil {
		font := &Font{cptr}
		runtime.SetFinalizer(font, (*Font).destroy)
		return font, nil
	}

	return nil, genericError
}

// Font constructor
// Creates a new font from memory
//
// The supported font formats are: TrueType, Type 1, CFF,
// OpenType, SFNT, X11 PCF, Windows FNT, BDF, PFR and Type 42.
func NewFontFromMemory(data []byte) (*Font, error) {
	if len(data) == 0 {
		return nil, errors.New("NewFontFromMemory: len(data)==0")
	}

	if cptr := C.sfFont_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data))); cptr != nil {
		font := &Font{cptr}
		runtime.SetFinalizer(font, (*Font).destroy)
		return font, nil
	}
	return nil, genericError
}

func (this *Font) Copy() *Font {
	font := &Font{C.sfFont_copy(this.cptr)}
	runtime.SetFinalizer(font, (*Font).destroy)
	return font
}

func (this *Font) destroy() {
	globalCtxSetActive(true)
	C.sfFont_destroy(this.cptr)
	globalCtxSetActive(false)
}

// Retrieve a glyph of the font
//
// 	codePoint:     Unicode code point of the character to get
// 	characterSize: Reference character size
// 	bold:          Retrieve the bold version or the regular one?
//
// return The glyph corresponding to codePoint and characterSize
func (this *Font) GetGlyph(codePoint uint, characterSize uint32, bold bool) (glyph Glyph) {
	glyph.fromC(C.sfFont_getGlyph(this.cptr, C.sfUint32(codePoint), C.uint(characterSize), goBool2C(bold)))
	return
}

//Get the kerning offset of two glyphs
//
// The kerning is an extra offset (negative) to apply between two
// glyphs when rendering them, to make the pair look more "natural".
// For example, the pair "AV" have a special kerning to make them
// closer than other characters. Most of the glyphs pairs have a
// kerning offset of zero, though.
//
// 	first:         Unicode code point of the first character
// 	second:        Unicode code point of the second character
// 	characterSize: Reference character size
//
// return Kerning value for first and second, in pixels
func (this *Font) GetKerning(first uint32, second uint32, characterSize uint) int {
	return int(C.sfFont_getKerning(this.cptr, C.sfUint32(first), C.sfUint32(second), C.uint(characterSize)))
}

// Get the line spacing
//
// Line spacing is the vertical offset to apply between two
// consecutive lines of text.
//
// 	characterSize: Reference character size
//
// return Line spacing, in pixels
func (this *Font) GetLineSpacing(characterSize uint) int {
	return int(C.sfFont_getLineSpacing(this.cptr, C.uint(characterSize)))
}

// Retrieve the texture containing the loaded glyphs of a certain size
//
// The contents of the returned texture changes as more glyphs
// are requested, thus it is not very relevant.
//
// 	characterSize: Reference character size
//
// Texture containing the glyphs of the requested size
func (this *Font) GetTexture(characterSize uint) Texture {
	return Texture{C.sfFont_getTexture(this.cptr, C.uint(characterSize))}
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Font) toCPtr() *C.sfFont {
	if this != nil {
		return this.cptr
	}
	return nil
}
