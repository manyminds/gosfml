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

func NewCircleShape(radius float32) *CircleShape {
	shape := &CircleShape{C.sfCircleShape_create(), nil}
	shape.SetRadius(radius)
	runtime.SetFinalizer(shape, (*CircleShape).Destroy)
	return shape
}

func (this *CircleShape) Copy() *CircleShape {
	shape := &CircleShape{C.sfCircleShape_copy(this.cptr), this.texture}
	runtime.SetFinalizer(shape, (*CircleShape).Destroy)
	return shape
}

func (this *CircleShape) Destroy() {
	C.sfCircleShape_destroy(this.cptr)
	this.cptr = nil
}

func (this *CircleShape) SetPosition(pos Vector2f) {
	C.sfCircleShape_setPosition(this.cptr, pos.toC())
}

func (this *CircleShape) SetScale(scale Vector2f) {
	C.sfCircleShape_setScale(this.cptr, scale.toC())
}

func (this *CircleShape) SetOrigin(orig Vector2f) {
	C.sfCircleShape_setOrigin(this.cptr, orig.toC())
}

func (this *CircleShape) SetRotation(rot float32) {
	C.sfCircleShape_setRotation(this.cptr, C.float(rot))
}

func (this *CircleShape) GetRotation() float32 {
	return float32(C.sfCircleShape_getRotation(this.cptr))
}

func (this *CircleShape) GetPosition() (position Vector2f) {
	position.fromC(C.sfCircleShape_getPosition(this.cptr))
	return
}

func (this *CircleShape) GetScale() (scale Vector2f) {
	scale.fromC(C.sfCircleShape_getScale(this.cptr))
	return
}

func (this *CircleShape) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfCircleShape_getOrigin(this.cptr))
	return
}

func (this *CircleShape) Move(offset Vector2f) {
	C.sfCircleShape_move(this.cptr, offset.toC())
}

func (this *CircleShape) Scale(factor Vector2f) {
	C.sfCircleShape_scale(this.cptr, factor.toC())
}

func (this *CircleShape) Rotate(angle float32) {
	C.sfCircleShape_rotate(this.cptr, C.float(angle))
}

func (this *CircleShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfCircleShape_setTexture(this.cptr, texture.toCPtr(), goBool2C(resetRect))
	this.texture = texture
}

func (this *CircleShape) SetTextureRect(rect IntRect) {
	C.sfCircleShape_setTextureRect(this.cptr, rect.toC())
}

func (this *CircleShape) SetFillColor(color Color) {
	C.sfCircleShape_setFillColor(this.cptr, color.toC())
}

func (this *CircleShape) SetOutlineColor(color Color) {
	C.sfCircleShape_setOutlineColor(this.cptr, color.toC())
}

func (this *CircleShape) SetOutlineThickness(thickness float32) {
	C.sfCircleShape_setOutlineThickness(this.cptr, C.float(thickness))
}

func (this *CircleShape) GetTexture() *Texture {
	return this.texture
}

func (this *CircleShape) GetTransform() (transform Transform) {
	transform.fromC(C.sfCircleShape_getTransform(this.cptr))
	return
}

func (this *CircleShape) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfCircleShape_getInverseTransform(this.cptr))
	return
}

func (this *CircleShape) GetTextureRect() (rect IntRect) {
	rect.fromC(C.sfCircleShape_getTextureRect(this.cptr))
	return
}

func (this *CircleShape) GetFillColor() (color Color) {
	color.fromC(C.sfCircleShape_getFillColor(this.cptr))
	return
}

func (this *CircleShape) GetOutlineColor() (color Color) {
	color.fromC(C.sfCircleShape_getOutlineColor(this.cptr))
	return
}

func (this *CircleShape) GetOutlineThickness() float32 {
	return float32(C.sfCircleShape_getOutlineThickness(this.cptr))
}

func (this *CircleShape) GetPointCount() uint {
	return uint(C.sfCircleShape_getPointCount(this.cptr))
}

func (this *CircleShape) GetPoint(index uint) (point Vector2f) {
	point.fromC(C.sfCircleShape_getPoint(this.cptr, C.uint(index)))
	return
}

func (this *CircleShape) SetRadius(radius float32) {
	C.sfCircleShape_setRadius(this.cptr, C.float(radius))
}

func (this *CircleShape) GetRadius() float32 {
	return float32(C.sfCircleShape_getRadius(this.cptr))
}

func (this *CircleShape) SetPointCount(count uint) {
	C.sfCircleShape_setPointCount(this.cptr, C.uint(count))
}

func (this *CircleShape) GetLocalBounds() (rect FloatRect) {
	rect.fromC(C.sfCircleShape_getLocalBounds(this.cptr))
	return
}

func (this *CircleShape) GetGlobalBounds() (rect FloatRect) {
	rect.fromC(C.sfCircleShape_getGlobalBounds(this.cptr))
	return
}

func (this *CircleShape) Draw(target RenderTarget, renderStates *RenderStates) {
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawCircleShape(target.(*RenderWindow).cptr, this.cptr, renderStates.toCPtr())
	case *RenderTexture:
		C.sfRenderTexture_drawCircleShape(target.(*RenderTexture).cptr, this.cptr, renderStates.toCPtr())
	}
}
