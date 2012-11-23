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

// #include <SFML/Graphics/RectangleShape.h> 
// #include <SFML/Graphics/RenderWindow.h> 
// #include <SFML/Graphics/RenderTexture.h> 
import "C"
import "runtime"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type RectangleShape struct {
	cptr    *C.sfRectangleShape
	texture *Texture //to prevent the GC from deleting the texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewRectangleShape() *RectangleShape {
	shape := &RectangleShape{C.sfRectangleShape_create(), nil}
	runtime.SetFinalizer(shape, (*RectangleShape).Destroy)
	return shape
}

func (this *RectangleShape) Copy() *RectangleShape {
	shape := &RectangleShape{C.sfRectangleShape_copy(this.cptr), this.texture}
	runtime.SetFinalizer(shape, (*RectangleShape).Destroy)
	return shape
}

func (this *RectangleShape) Destroy() {
	C.sfRectangleShape_destroy(this.cptr)
	this.cptr = nil
}

func (this *RectangleShape) SetPosition(pos Vector2f) {
	C.sfRectangleShape_setPosition(this.cptr, pos.toC())
}

func (this *RectangleShape) SetScale(scale Vector2f) {
	C.sfRectangleShape_setScale(this.cptr, scale.toC())
}

func (this *RectangleShape) SetOrigin(orig Vector2f) {
	C.sfRectangleShape_setOrigin(this.cptr, orig.toC())
}

func (this *RectangleShape) SetRotation(rot float32) {
	C.sfRectangleShape_setRotation(this.cptr, C.float(rot))
}

func (this *RectangleShape) GetRotation() float32 {
	return float32(C.sfRectangleShape_getRotation(this.cptr))
}

func (this *RectangleShape) GetPosition() (position Vector2f) {
	position.fromC(C.sfRectangleShape_getPosition(this.cptr))
	return
}

func (this *RectangleShape) GetScale() (scale Vector2f) {
	scale.fromC(C.sfRectangleShape_getScale(this.cptr))
	return
}

func (this *RectangleShape) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfRectangleShape_getOrigin(this.cptr))
	return
}

func (this *RectangleShape) Move(offset Vector2f) {
	C.sfRectangleShape_move(this.cptr, offset.toC())
}

func (this *RectangleShape) Scale(factor Vector2f) {
	C.sfRectangleShape_scale(this.cptr, factor.toC())
}

func (this *RectangleShape) Rotate(angle float32) {
	C.sfRectangleShape_rotate(this.cptr, C.float(angle))
}

func (this *RectangleShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfRectangleShape_setTexture(this.cptr, texture.cptr, goBool2C(resetRect))
	this.texture = texture
}

func (this *RectangleShape) SetTextureRect(rect IntRect) {
	C.sfRectangleShape_setTextureRect(this.cptr, rect.toC())
}

func (this *RectangleShape) GetTexture() *Texture {
	return this.texture
}

func (this *RectangleShape) GetTextureRect() (rect IntRect) {
	rect.fromC(C.sfRectangleShape_getTextureRect(this.cptr))
	return
}

func (this *RectangleShape) SetFillColor(color Color) {
	C.sfRectangleShape_setFillColor(this.cptr, color.toC())
}

func (this *RectangleShape) SetOutlineColor(color Color) {
	C.sfRectangleShape_setOutlineColor(this.cptr, color.toC())
}

func (this *RectangleShape) SetOutlineThickness(thickness float32) {
	C.sfRectangleShape_setOutlineThickness(this.cptr, C.float(thickness))
}

func (this *RectangleShape) SetSize(size Vector2f) {
	C.sfRectangleShape_setSize(this.cptr, size.toC())
}

func (this *RectangleShape) GetSize() (size Vector2f) {
	size.fromC(C.sfRectangleShape_getSize(this.cptr))
	return
}

func (this *RectangleShape) GetTransform() (transform Transform) {
	transform.fromC(C.sfRectangleShape_getTransform(this.cptr))
	return
}

func (this *RectangleShape) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfRectangleShape_getInverseTransform(this.cptr))
	return
}

func (this *RectangleShape) GetFillColor() (color Color) {
	color.fromC(C.sfRectangleShape_getFillColor(this.cptr))
	return
}

func (this *RectangleShape) GetOutlineColor() (color Color) {
	color.fromC(C.sfRectangleShape_getOutlineColor(this.cptr))
	return
}

func (this *RectangleShape) GetOutlineThickness() float32 {
	return float32(C.sfRectangleShape_getOutlineThickness(this.cptr))
}

func (this *RectangleShape) GetPointCount() uint {
	return uint(C.sfRectangleShape_getPointCount(this.cptr))
}

func (this *RectangleShape) GetPoint(index uint) (point Vector2f) {
	point.fromC(C.sfRectangleShape_getPoint(this.cptr, C.uint(index)))
	return
}

func (this *RectangleShape) GetLocalBounds() (rect FloatRect) {
	rect.fromC(C.sfRectangleShape_getLocalBounds(this.cptr))
	return
}

func (this *RectangleShape) GetGlobalBounds() (rect FloatRect) {
	rect.fromC(C.sfRectangleShape_getGlobalBounds(this.cptr))
	return
}

func (this *RectangleShape) Draw(target RenderTarget, renderStates *RenderStates) {
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawRectangleShape(target.(*RenderWindow).cptr, this.cptr, renderStates.toCPtr())
	case *RenderTexture:
		C.sfRenderTexture_drawRectangleShape(target.(*RenderTexture).cptr, this.cptr, renderStates.toCPtr())
	}
}
