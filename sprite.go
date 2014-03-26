// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Sprite.h>
// #include <SFML/Graphics/RenderWindow.h>
// #include <SFML/Graphics/RenderTexture.h>
import "C"
import "runtime"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Sprite struct {
	cptr    *C.sfSprite
	texture *Texture //to prevent the GC from deleting the texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new sprite with a given texture (can be nil to use no texture)
func NewSprite(tex *Texture) (*Sprite, error) {
	if cptr := C.sfSprite_create(); cptr != nil {
		shape := &Sprite{cptr: cptr}
		runtime.SetFinalizer(shape, (*Sprite).destroy)
		shape.SetTexture(tex, true)

		return shape, nil
	}

	return nil, genericError
}

// Copy an existing sprite
func (this *Sprite) Copy() *Sprite {
	sprite := &Sprite{C.sfSprite_copy(this.cptr), this.texture}
	runtime.SetFinalizer(sprite, (*Sprite).destroy)
	return sprite
}

// Destroy an existing sprite
func (this *Sprite) destroy() {
	C.sfSprite_destroy(this.cptr)
}

// Set the position of a sprite
//
// This function completely overwrites the previous position.
// See Sprite.Move to apply an offset based on the previous position instead.
// The default position of a sprite Sprite object is (0, 0).
func (this *Sprite) SetPosition(pos Vector2f) {
	C.sfSprite_setPosition(this.cptr, pos.toC())
}

// Set the scale factors of a sprite
//
// This function completely overwrites the previous scale.
// See sfSprite_scale to add a factor based on the previous scale instead.
// The default scale of a sprite Sprite object is (1, 1).
func (this *Sprite) SetScale(scale Vector2f) {
	C.sfSprite_setScale(this.cptr, scale.toC())
}

// Set the local origin of a sprite
//
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a sprite Sprite object is (0, 0).
func (this *Sprite) SetOrigin(orig Vector2f) {
	C.sfSprite_setOrigin(this.cptr, orig.toC())
}

// Set the orientation of a sprite
//
// This function completely overwrites the previous rotation.
// See Sprite.Rotate to add an angle based on the previous rotation instead.
// The default rotation of a sprite Sprite object is 0.
func (this *Sprite) SetRotation(rot float32) {
	C.sfSprite_setRotation(this.cptr, C.float(rot))
}

// Move a sprite by a given offset
//
// This function adds to the current position of the object,
// unlike Sprite.SetPosition which overwrites it.
func (this *Sprite) Move(offset Vector2f) {
	C.sfSprite_move(this.cptr, offset.toC())
}

// Scale a sprite
//
// This function multiplies the current scale of the object,
// unlike Sprite.SetScale which overwrites it.
func (this *Sprite) Scale(factor Vector2f) {
	C.sfSprite_scale(this.cptr, factor.toC())
}

// Rotate a sprite
//
// This function adds to the current rotation of the object,
// unlike Sprite.SetRotation which overwrites it.
func (this *Sprite) Rotate(angle float32) {
	C.sfSprite_rotate(this.cptr, C.float(angle))
}

// Get the orientation of a sprite
//
// The rotation is always in the range [0, 360].
func (this *Sprite) GetRotation() float32 {
	return float32(C.sfSprite_getRotation(this.cptr))
}

// Get the position of a sprite
func (this *Sprite) GetPosition() (pos Vector2f) {
	pos.fromC(C.sfSprite_getPosition(this.cptr))
	return
}

// Get the current scale of a sprite
func (this *Sprite) GetScale() (scale Vector2f) {
	scale.fromC(C.sfSprite_getScale(this.cptr))
	return
}

// Get the local origin of a sprite
func (this *Sprite) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfSprite_getOrigin(this.cptr))
	return
}

// Change the source texture of a sprite
//
// The texture argument refers to a texture that must
// exist as long as the sprite uses it. Indeed, the sprite
// doesn't store its own copy of the texture, but rather keeps
// a pointer to the one that you passed to this function.
// If the source texture is destroyed and the sprite tries to
// use it, the behaviour is undefined.
// If resetRect is true, the TextureRect property of
// the sprite is automatically adjusted to the size of the new
// texture. If it is false, the texture rect is left unchanged.
//
// 	texture:   New texture
// 	resetRect: Should the texture rect be reset to the size of the new texture?
func (this *Sprite) SetTexture(texture *Texture, resetRect bool) {
	C.sfSprite_setTexture(this.cptr, texture.toCPtr(), goBool2C(resetRect))
	this.texture = texture
}

// Set the sub-rectangle of the texture that a sprite will display
//
// The texture rect is useful when you don't want to display
// the whole texture, but rather a part of it.
// By default, the texture rect covers the entire texture.
//
// 	rect: Rectangle defining the region of the texture to display
func (this *Sprite) SetTextureRect(rect IntRect) {
	C.sfSprite_setTextureRect(this.cptr, rect.toC())
}

// Get the source texture of a sprite
//
// If the sprite has no source texture, nil is returned.
func (this *Sprite) GetTexture() *Texture {
	return this.texture
}

// Get the sub-rectangle of the texture displayed by a sprite
func (this *Sprite) GetTextureRect() (rect IntRect) {
	rect.fromC(C.sfSprite_getTextureRect(this.cptr))
	return
}

// Get the global color of a sprite
func (this *Sprite) GetColor() (color Color) {
	color.fromC(C.sfSprite_getColor(this.cptr))
	return
}

// Set the global color of a sprite
//
// This color is modulated (multiplied) with the sprite's
// texture. It can be used to colorize the sprite, or change
// its global opacity.
// By default, the sprite's color is opaque white.
func (this *Sprite) SetColor(color Color) {
	C.sfSprite_setColor(this.cptr, color.toC())
}

// Get the combined transform of a sprite
func (this *Sprite) GetTransform() (trans Transform) {
	trans.fromC(C.sfSprite_getTransform(this.cptr))
	return
}

// Get the inverse of the combined transform of a sprite
func (this *Sprite) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfSprite_getInverseTransform(this.cptr))
	return
}

// Get the local bounding rectangle of a sprite
//
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
func (this *Sprite) GetLocalBounds() (rect FloatRect) {
	rect.fromC(C.sfSprite_getLocalBounds(this.cptr))
	return
}

// Get the global bounding rectangle of a sprite
//
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// sprite in the global 2D world's coordinate system.
func (this *Sprite) GetGlobalBounds() (rect FloatRect) {
	rect.fromC(C.sfSprite_getGlobalBounds(this.cptr))
	return
}

// Draws a RectangleShape on a render target
func (this *Sprite) Draw(target RenderTarget, renderStates RenderStates) {
	rs := renderStates.toC()
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawSprite(target.(*RenderWindow).cptr, this.cptr, &rs)
	case *RenderTexture:
		C.sfRenderTexture_drawSprite(target.(*RenderTexture).cptr, this.cptr, &rs)
	}
}
