// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/CircleShape.h>
// #include <SFML/Graphics/RenderWindow.h>
// #include <SFML/Graphics/RenderTexture.h>
import "C"
import "runtime"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type CircleShape struct {
	cptr    *C.sfCircleShape
	texture *Texture //to prevent the GC from deleting the texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new circle shape with a given radius
func NewCircleShape() (*CircleShape, error) {
	if cptr := C.sfCircleShape_create(); cptr != nil {
		shape := &CircleShape{cptr, nil}
		runtime.SetFinalizer(shape, (*CircleShape).destroy)
		return shape, nil
	}
	return nil, genericError
}

// Copy an existing circle shape
func (this *CircleShape) Copy() *CircleShape {
	shape := &CircleShape{C.sfCircleShape_copy(this.cptr), this.texture}
	runtime.SetFinalizer(shape, (*CircleShape).destroy)
	return shape
}

// Destroy an existing circle Shape
func (this *CircleShape) destroy() {
	C.sfCircleShape_destroy(this.cptr)
}

// Set the position of a circle shape
//
// This function completely overwrites the previous position.
// See sfCircleShape_move to apply an offset based on the previous position instead.
// The default position of a circle Shape object is (0, 0).
func (this *CircleShape) SetPosition(pos Vector2f) {
	C.sfCircleShape_setPosition(this.cptr, pos.toC())
}

// Set the scale factors of a circle shape
//
// This function completely overwrites the previous scale.
// See sfCircleShape_scale to add a factor based on the previous scale instead.
// The default scale of a circle Shape object is (1, 1).
func (this *CircleShape) SetScale(scale Vector2f) {
	C.sfCircleShape_setScale(this.cptr, scale.toC())
}

// Set the local origin of a circle shape
//
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a circle Shape object is (0, 0).
func (this *CircleShape) SetOrigin(orig Vector2f) {
	C.sfCircleShape_setOrigin(this.cptr, orig.toC())
}

// Set the orientation of a circle shape
//
// This function completely overwrites the previous rotation.
// See sfCircleShape_rotate to add an angle based on the previous rotation instead.
// The default rotation of a circle Shape object is 0.
func (this *CircleShape) SetRotation(rot float32) {
	C.sfCircleShape_setRotation(this.cptr, C.float(rot))
}

// Get the orientation of a circle shape
//
// The rotation is always in the range [0, 360].
func (this *CircleShape) GetRotation() float32 {
	return float32(C.sfCircleShape_getRotation(this.cptr))
}

// Get the position of a circle shape
func (this *CircleShape) GetPosition() (position Vector2f) {
	position.fromC(C.sfCircleShape_getPosition(this.cptr))
	return
}

// Get the current scale of a circle shape
func (this *CircleShape) GetScale() (scale Vector2f) {
	scale.fromC(C.sfCircleShape_getScale(this.cptr))
	return
}

// Get the local origin of a circle shape
func (this *CircleShape) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfCircleShape_getOrigin(this.cptr))
	return
}

// Move a circle shape by a given offset
//
// This function adds to the current position of the object,
// unlike CircleShape.SetPosition which overwrites it.
func (this *CircleShape) Move(offset Vector2f) {
	C.sfCircleShape_move(this.cptr, offset.toC())
}

// Scale a circle shape
//
// This function multiplies the current scale of the object,
// unlike CircleShape.SetScale which overwrites it.
func (this *CircleShape) Scale(factor Vector2f) {
	C.sfCircleShape_scale(this.cptr, factor.toC())
}

// Rotate a circle shape
//
// This function adds to the current rotation of the object,
// unlike CircleShape.SetRotation which overwrites it.
func (this *CircleShape) Rotate(angle float32) {
	C.sfCircleShape_rotate(this.cptr, C.float(angle))
}

// Change the source texture of a circle shape
//
// texture can be nil to disable texturing.
// If resetRect is true, the TextureRect property of
// the shape is automatically adjusted to the size of the new
// texture. If it is false, the texture rect is left unchanged.
//
// 	texture:   New texture
// 	resetRect: Should the texture rect be reset to the size of the new texture?
func (this *CircleShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfCircleShape_setTexture(this.cptr, texture.toCPtr(), goBool2C(resetRect))
	this.texture = texture
}

// Set the sub-rectangle of the texture that a circle shape will display
//
// The texture rect is useful when you don't want to display
// the whole texture, but rather a part of it.
// By default, the texture rect covers the entire texture.
func (this *CircleShape) SetTextureRect(rect IntRect) {
	C.sfCircleShape_setTextureRect(this.cptr, rect.toC())
}

// Set the fill color of a circle shape
//
// This color is modulated (multiplied) with the shape's
// texture if any. It can be used to colorize the shape,
// or change its global opacity.
// You can use sfTransparent to make the inside of
// the shape transparent, and have the outline alone.
// By default, the shape's fill color is opaque white.
func (this *CircleShape) SetFillColor(color Color) {
	C.sfCircleShape_setFillColor(this.cptr, color.toC())
}

// Set the outline color of a circle shape
//
// You can use sfTransparent to disable the outline.
// By default, the shape's outline color is opaque white.
func (this *CircleShape) SetOutlineColor(color Color) {
	C.sfCircleShape_setOutlineColor(this.cptr, color.toC())
}

// Set the thickness of a circle shape's outline
//
// This number cannot be negative. Using zero disables
// the outline.
// By default, the outline thickness is 0.
func (this *CircleShape) SetOutlineThickness(thickness float32) {
	C.sfCircleShape_setOutlineThickness(this.cptr, C.float(thickness))
}

// Get the source texture of a circle shape
//
// If the shape has no source texture, nil is returned.
// The returned pointer is const, which means that you can't
// modify the texture when you retrieve it with this function.
func (this *CircleShape) GetTexture() *Texture {
	return this.texture
}

// Get the combined transform of a circle shape
func (this *CircleShape) GetTransform() (transform Transform) {
	transform.fromC(C.sfCircleShape_getTransform(this.cptr))
	return
}

// Get the inverse of the combined transform of a circle shape
func (this *CircleShape) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfCircleShape_getInverseTransform(this.cptr))
	return
}

// Get the sub-rectangle of the texture displayed by a circle shape
func (this *CircleShape) GetTextureRect() (rect IntRect) {
	rect.fromC(C.sfCircleShape_getTextureRect(this.cptr))
	return
}

// Get the fill color of a circle shape
func (this *CircleShape) GetFillColor() (color Color) {
	color.fromC(C.sfCircleShape_getFillColor(this.cptr))
	return
}

// Get the outline color of a circle shape
func (this *CircleShape) GetOutlineColor() (color Color) {
	color.fromC(C.sfCircleShape_getOutlineColor(this.cptr))
	return
}

// Get the outline thickness of a circle shape
func (this *CircleShape) GetOutlineThickness() float32 {
	return float32(C.sfCircleShape_getOutlineThickness(this.cptr))
}

func (this *CircleShape) GetPointCount() uint {
	return uint(C.sfCircleShape_getPointCount(this.cptr))
}

// Get the total number of points of a circle shape
func (this *CircleShape) GetPoint(index uint) (point Vector2f) {
	point.fromC(C.sfCircleShape_getPoint(this.cptr, C.uint(index)))
	return
}

// Set the radius of a circle
func (this *CircleShape) SetRadius(radius float32) {
	C.sfCircleShape_setRadius(this.cptr, C.float(radius))
}

// Get the radius of a circle
func (this *CircleShape) GetRadius() float32 {
	return float32(C.sfCircleShape_getRadius(this.cptr))
}

// Set the number of points of a circle
func (this *CircleShape) SetPointCount(count uint) {
	C.sfCircleShape_setPointCount(this.cptr, C.uint(count))
}

// Get the local bounding rectangle of a circle shape
//
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
func (this *CircleShape) GetLocalBounds() (rect FloatRect) {
	rect.fromC(C.sfCircleShape_getLocalBounds(this.cptr))
	return
}

// Get the global bounding rectangle of a circle shape
//
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// sprite in the global 2D world's coordinate system.
func (this *CircleShape) GetGlobalBounds() (rect FloatRect) {
	rect.fromC(C.sfCircleShape_getGlobalBounds(this.cptr))
	return
}

//Draws a CircleShape on a render target
func (this *CircleShape) Draw(target RenderTarget, renderStates RenderStates) {
	rs := renderStates.toC()
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawCircleShape(target.(*RenderWindow).cptr, this.cptr, &rs)
	case *RenderTexture:
		C.sfRenderTexture_drawCircleShape(target.(*RenderTexture).cptr, this.cptr, &rs)
	}
}
