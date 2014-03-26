// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/ConvexShape.h>
// #include <SFML/Graphics/RenderWindow.h>
// #include <SFML/Graphics/RenderTexture.h>
import "C"
import "runtime"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type ConvexShape struct {
	cptr    *C.sfConvexShape
	texture *Texture //to prevent the GC from deleting the texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewConvexShape() (*ConvexShape, error) {
	if cptr := C.sfConvexShape_create(); cptr != nil {
		shape := &ConvexShape{cptr, nil}
		runtime.SetFinalizer(shape, (*ConvexShape).destroy)
		return shape, nil
	}
	return nil, genericError
}

//Copy an existing convex shape
func (this *ConvexShape) Copy() *ConvexShape {
	shape := &ConvexShape{C.sfConvexShape_copy(this.cptr), this.texture}
	runtime.SetFinalizer(shape, (*ConvexShape).destroy)
	return shape
}

func (this *ConvexShape) destroy() {
	C.sfConvexShape_destroy(this.cptr)
}

// Set the position of a convex shape
//
// This function completely overwrites the previous position.
// See sfConvexShape_move to apply an offset based on the previous position instead.
// The default position of a circle Shape object is (0, 0).
func (this *ConvexShape) SetPosition(pos Vector2f) {
	C.sfConvexShape_setPosition(this.cptr, pos.toC())
}

// Set the local origin of a convex shape
//
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a circle Shape object is (0, 0).
func (this *ConvexShape) SetScale(scale Vector2f) {
	C.sfConvexShape_setScale(this.cptr, scale.toC())
}

// Set the local origin of a convex shape
//
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a circle Shape object is (0, 0).
func (this *ConvexShape) SetOrigin(orig Vector2f) {
	C.sfConvexShape_setOrigin(this.cptr, orig.toC())
}

// Set the scale factors of a convex shape
//
// This function completely overwrites the previous scale.
// See sfConvexShape_scale to add a factor based on the previous scale instead.
// The default scale of a circle Shape object is (1, 1).
func (this *ConvexShape) SetRotation(rot float32) {
	C.sfConvexShape_setRotation(this.cptr, C.float(rot))
}

// Get the orientation of a convex shape
func (this *ConvexShape) GetRotation() float32 {
	return float32(C.sfConvexShape_getRotation(this.cptr))
}

// Get the position of a convex shape
func (this *ConvexShape) GetPosition() (position Vector2f) {
	position.fromC(C.sfConvexShape_getPosition(this.cptr))
	return
}

// Get the current scale of a convex shape
func (this *ConvexShape) GetScale() (scale Vector2f) {
	scale.fromC(C.sfConvexShape_getScale(this.cptr))
	return
}

// Get the local origin of a convex shape
func (this *ConvexShape) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfConvexShape_getOrigin(this.cptr))
	return
}

// Move a convex shape by a given offset
//
// This function adds to the current position of the object,
// unlike ConvexShape.SetPosition which overwrites it.
func (this *ConvexShape) Move(offset Vector2f) {
	C.sfConvexShape_move(this.cptr, offset.toC())
}

// Scale a convex shape
//
// This function multiplies the current scale of the object,
// unlike ConvexShape.SetScale which overwrites it.
func (this *ConvexShape) Scale(factor Vector2f) {
	C.sfConvexShape_scale(this.cptr, factor.toC())
}

// Rotate a convex shape
//
// This function adds to the current rotation of the object,
// unlike ConvexShape.SetRotation which overwrites it.
func (this *ConvexShape) Rotate(angle float32) {
	C.sfConvexShape_rotate(this.cptr, C.float(angle))
}

// Change the source texture of a convex shape
//
// The texture argument can be nil to disable texturing.
// If resetRect is true, the TextureRect property of
// the shape is automatically adjusted to the size of the new
// texture. If it is false, the texture rect is left unchanged.
func (this *ConvexShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfConvexShape_setTexture(this.cptr, texture.toCPtr(), goBool2C(resetRect))
	this.texture = texture
}

// Set the sub-rectangle of the texture that a convex shape will display
//
// The texture rect is useful when you don't want to display
// the whole texture, but rather a part of it.
// By default, the texture rect covers the entire texture.
func (this *ConvexShape) SetTextureRect(rect IntRect) {
	C.sfConvexShape_setTextureRect(this.cptr, rect.toC())
}

// Set the fill color of a convex shape
//
// This color is modulated (multiplied) with the shape's
// texture if any. It can be used to colorize the shape,
// or change its global opacity.
// You can use sfTransparent to make the inside of
// the shape transparent, and have the outline alone.
// By default, the shape's fill color is opaque white.
func (this *ConvexShape) SetFillColor(color Color) {
	C.sfConvexShape_setFillColor(this.cptr, color.toC())
}

// Set the outline color of a convex shape
//
// You can use sfTransparent to disable the outline.
// By default, the shape's outline color is opaque white.
func (this *ConvexShape) SetOutlineColor(color Color) {
	C.sfConvexShape_setOutlineColor(this.cptr, color.toC())
}

// Set the thickness of a convex shape's outline
//
// This number cannot be negative. Using zero disables
// the outline.
// By default, the outline thickness is 0.
func (this *ConvexShape) SetOutlineThickness(thickness float32) {
	C.sfConvexShape_setOutlineThickness(this.cptr, C.float(thickness))
}

// Get the source texture of a convex shape
//
// If the shape has no source texture, a nil pointer is returned.
func (this *ConvexShape) GetTexture() *Texture {
	return this.texture
}

// Get the sub-rectangle of the texture displayed by a convex shape
func (this *ConvexShape) GetTextureRect() (rect IntRect) {
	rect.fromC(C.sfConvexShape_getTextureRect(this.cptr))
	return
}

// Get the combined transform of a convex shape
func (this *ConvexShape) GetTransform() (transform Transform) {
	transform.fromC(C.sfConvexShape_getTransform(this.cptr))
	return
}

// Get the inverse of the combined transform of a convex shape
func (this *ConvexShape) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfConvexShape_getInverseTransform(this.cptr))
	return
}

// Get the fill color of a convex shape
func (this *ConvexShape) GetFillColor() (color Color) {
	color.fromC(C.sfConvexShape_getFillColor(this.cptr))
	return
}

// Get the outline color of a convex shape
func (this *ConvexShape) GetOutlineColor() (color Color) {
	color.fromC(C.sfConvexShape_getOutlineColor(this.cptr))
	return
}

// Get the outline thickness of a convex shape
func (this *ConvexShape) GetOutlineThickness() float32 {
	return float32(C.sfConvexShape_getOutlineThickness(this.cptr))
}

// Get the total number of points of a convex shape
func (this *ConvexShape) GetPointCount() uint {
	return uint(C.sfConvexShape_getPointCount(this.cptr))
}

// Get a point of a convex shape
//
// The result is undefined if index is out of the valid range.
func (this *ConvexShape) GetPoint(index uint) (point Vector2f) {
	point.fromC(C.sfConvexShape_getPoint(this.cptr, C.uint(index)))
	return
}

// Set the number of points of a convex shap
//
// count must be greater than 2 to define a valid shape.
func (this *ConvexShape) SetPointCount(count uint) {
	C.sfConvexShape_setPointCount(this.cptr, C.uint(count))
}

// Set the position of a point in a convex shape
//
// Don't forget that the polygon must remain convex, and
// the points need to stay ordered!
// SetPointCount must be called first in order to set the total
// number of points. The result is undefined if index is out
// of the valid range.
func (this *ConvexShape) SetPoint(index uint, point Vector2f) {
	C.sfConvexShape_setPoint(this.cptr, C.uint(index), point.toC())
}

// Get the local bounding rectangle of a convex shape
//
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
func (this *ConvexShape) GetLocalBounds() (rect FloatRect) {
	rect.fromC(C.sfConvexShape_getLocalBounds(this.cptr))
	return
}

// Get the global bounding rectangle of a convex shape
//
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// sprite in the global 2D world's coordinate system.
func (this *ConvexShape) GetGlobalBounds() (rect FloatRect) {
	rect.fromC(C.sfConvexShape_getGlobalBounds(this.cptr))
	return
}

// Draws a convex Shape on a render target
func (this *ConvexShape) Draw(target RenderTarget, renderStates RenderStates) {
	rs := renderStates.toC()
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawConvexShape(target.(*RenderWindow).cptr, this.cptr, &rs)
	case *RenderTexture:
		C.sfRenderTexture_drawConvexShape(target.(*RenderTexture).cptr, this.cptr, &rs)
	}
}
