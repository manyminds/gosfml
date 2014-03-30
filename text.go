// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Text.h>
// #include <SFML/Graphics/RenderWindow.h>
// #include <SFML/Graphics/RenderTexture.h>
// #include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	TextRegular    = C.sfTextRegular    ///< Regular characters, no style
	TextBold       = C.sfTextBold       ///< Characters are bold
	TextItalic     = C.sfTextItalic     ///< Characters are in italic
	TextUnderlined = C.sfTextUnderlined ///< Characters are underlined
)

type TextStyle uint32

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Text struct {
	cptr *C.sfText
	font *Font
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new text with a given font (can be nil).
func NewText(font *Font) (*Text, error) {
	if cptr := C.sfText_create(); cptr != nil {
		text := &Text{cptr: cptr}
		runtime.SetFinalizer(text, (*Text).destroy)
		text.SetFont(font)

		return text, nil
	}

	return nil, genericError
}

// Destroy an existing text
func (this *Text) destroy() {
	C.sfText_destroy(this.cptr)
}

// Copy an existing text
func (this *Text) Copy() *Text {
	text := &Text{C.sfText_copy(this.cptr), this.font}
	runtime.SetFinalizer(text, (*Text).destroy)
	return text
}

// Set the position of a text
//
// This function completely overwrites the previous position.
// See Text.Move to apply an offset based on the previous position instead.
// The default position of a text Text object is (0, 0).
//
// 	position: New position
func (this *Text) SetPosition(pos Vector2f) {
	C.sfText_setPosition(this.cptr, pos.toC())
}

// Set the scale factors of a text
//
// This function completely overwrites the previous scale.
// See Text.Scale to add a factor based on the previous scale instead.
// The default scale of a text Text object is (1, 1).
//
// 	scale: New scale factors
func (this *Text) SetScale(scale Vector2f) {
	C.sfText_setScale(this.cptr, scale.toC())
}

// Set the local origin of a text
//
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a text object is (0, 0).
//
// 	origin: New origin
func (this *Text) SetOrigin(orig Vector2f) {
	C.sfText_setOrigin(this.cptr, orig.toC())
}

// Set the orientation of a text
//
// This function completely overwrites the previous rotation.
// See sfText_rotate to add an angle based on the previous rotation instead.
// The default rotation of a text Text object is 0.
//
// 	rot: New rotation, in degrees
func (this *Text) SetRotation(rot float32) {
	C.sfText_setRotation(this.cptr, C.float(rot))
}

// Move a text by a given offset
//
// This function adds to the current position of the object,
// unlike Text.SetPosition which overwrites it.
//
// 	offset: Offset
func (this *Text) Move(offset Vector2f) {
	C.sfText_move(this.cptr, offset.toC())
}

// Scale a text
//
// This function multiplies the current scale of the object,
// unlike Text.SetScale which overwrites it.
//
// 	factor: Scale factors
func (this *Text) Scale(factor Vector2f) {
	C.sfText_scale(this.cptr, factor.toC())
}

// Rotate a text
//
// This function adds to the current rotation of the object,
// unlike Text.SetRotation which overwrites it.
//
// 	angle: Angle of rotation, in degrees
func (this *Text) Rotate(angle float32) {
	C.sfText_rotate(this.cptr, C.float(angle))
}

// Get the orientation of a text
//
// The rotation is always in the range [0, 360].
func (this *Text) GetRotation() float32 {
	return float32(C.sfText_getRotation(this.cptr))
}

// Get the position of a text
func (this *Text) GetPosition() (pos Vector2f) {
	pos.fromC(C.sfText_getPosition(this.cptr))
	return
}

// Get the current scale of a text
func (this *Text) GetScale() (scale Vector2f) {
	scale.fromC(C.sfText_getScale(this.cptr))
	return
}

// Get the local origin of a text
func (this *Text) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfText_getOrigin(this.cptr))
	return
}

// Get the combined transform of a text
func (this *Text) GetTransform() (trans Transform) {
	trans.fromC(C.sfText_getTransform(this.cptr))
	return
}

// Get the inverse of the combined transform of a text
func (this *Text) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfText_getInverseTransform(this.cptr))
	return
}

// Set the string of a text (from a unicode string)
func (this *Text) SetString(text string) {
	runes := strToRunes(text)
	C.sfText_setUnicodeString(this.cptr, (*C.sfUint32)(unsafe.Pointer(&runes[0])))
}

// Set the font of a text
func (this *Text) SetFont(font *Font) {
	C.sfText_setFont(this.cptr, font.toCPtr())
	this.font = font
}

// Set the character size of a text
//
// The default size is 30.
func (this *Text) SetCharacterSize(size uint) {
	C.sfText_setCharacterSize(this.cptr, C.uint(size))
}

// Set the style of a text
//
// You can pass a combination of one or more styles, for
// example TextBold | TextItalic.
// The default style is TextRegular.
func (this *Text) SetStyle(style TextStyle) {
	C.sfText_setStyle(this.cptr, C.sfUint32(style))
}

// Set the global color of a text
//
// By default, the text's color is opaque white.
func (this *Text) SetColor(color Color) {
	C.sfText_setColor(this.cptr, color.toC())
}

// Get the string of a text (returns a unicode string)
func (this *Text) GetString() string {
	cstr := C.sfText_getUnicodeString(this.cptr)
	return utf32CString2Go(cstr)
}

// Get the font used by a text
//
// If the text has no font attached, a nil pointer is returned.
func (this *Text) GetFont() *Font {
	return this.font
}

// Get the size of the characters of a text
func (this *Text) GetCharacterSize() uint {
	return uint(C.sfText_getCharacterSize(this.cptr))
}

// Get the style of a text
func (this *Text) GetStyle() TextStyle {
	return TextStyle(C.sfText_getStyle(this.cptr))
}

// Get the global color of a text
func (this *Text) GetColor() (color Color) {
	color.fromC(C.sfText_getColor(this.cptr))
	return
}

// Return the position of the index-th character in a text
//
// This function computes the visual position of a character
// from its index in the string. The returned position is
// in global coordinates (translation, rotation, scale and
// origin are applied).
// If index is out of range, the position of the end of
// the string is returned.
func (this *Text) FindCharacterPos(index uint) (pos Vector2f) {
	pos.fromC(C.sfText_findCharacterPos(this.cptr, C.size_t(index)))
	return
}

// Get the local bounding rectangle of a text
//
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
func (this *Text) GetLocalBounds() (rect FloatRect) {
	rect.fromC(C.sfText_getLocalBounds(this.cptr))
	return
}

// Get the global bounding rectangle of a text
//
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// text in the global 2D world's coordinate system.
func (this *Text) GetGlobalBounds() (rect FloatRect) {
	rect.fromC(C.sfText_getGlobalBounds(this.cptr))
	return
}

// Draws a Text on a render target
func (this *Text) Draw(target RenderTarget, renderStates RenderStates) {
	rs := renderStates.toC()
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawText(target.(*RenderWindow).cptr, this.cptr, &rs)
	case *RenderTexture:
		C.sfRenderTexture_drawText(target.(*RenderTexture).cptr, this.cptr, &rs)
	}
}
