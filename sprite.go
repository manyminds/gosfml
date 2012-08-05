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

// #include <SFML/Graphics.h> 
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

func NewSprite(tex *Texture) *Sprite {
	shape := &Sprite{C.sfSprite_create(), nil}
	runtime.SetFinalizer(shape, (*Sprite).Destroy)

	//set texture
	shape.SetTexture(tex, true)

	return shape
}

func (this *Sprite) Copy() *Sprite {
	sprite := &Sprite{C.sfSprite_copy(this.cptr), this.texture}
	runtime.SetFinalizer(sprite, (*Sprite).Destroy)
	return sprite
}

func (this *Sprite) Destroy() {
	C.sfSprite_destroy(this.cptr)
	this.cptr = nil
}

func (this *Sprite) SetPosition(pos Vector2f) {
	C.sfSprite_setPosition(this.cptr, pos.toC())
}

func (this *Sprite) SetScale(scale Vector2f) {
	C.sfSprite_setScale(this.cptr, scale.toC())
}

func (this *Sprite) SetOrigin(orig Vector2f) {
	C.sfSprite_setOrigin(this.cptr, orig.toC())
}

func (this *Sprite) SetRotation(rot float32) {
	C.sfSprite_setRotation(this.cptr, C.float(rot))
}

func (this *Sprite) Move(offset Vector2f) {
	C.sfSprite_move(this.cptr, offset.toC())
}

func (this *Sprite) Scale(factor Vector2f) {
	C.sfSprite_scale(this.cptr, factor.toC())
}

func (this *Sprite) Rotate(angle float32) {
	C.sfSprite_rotate(this.cptr, C.float(angle))
}

func (this *Sprite) GetRotation() float32 {
	return float32(C.sfSprite_getRotation(this.cptr))
}

func (this *Sprite) GetPosition() (pos Vector2f) {
	pos.fromC(C.sfSprite_getPosition(this.cptr))
	return
}

func (this *Sprite) GetScale() (scale Vector2f) {
	scale.fromC(C.sfSprite_getScale(this.cptr))
	return
}

func (this *Sprite) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfSprite_getOrigin(this.cptr))
	return
}

func (this *Sprite) SetTexture(texture *Texture, resetRect bool) {
	C.sfSprite_setTexture(this.cptr, texture.toCPtr(), goBool2C(resetRect))
	this.texture = texture
}

func (this *Sprite) SetTextureRect(rect Recti) {
	C.sfSprite_setTextureRect(this.cptr, rect.toC())
}

func (this *Sprite) GetTexture() *Texture {
	return this.texture
}

func (this *Sprite) GetTextureRect() (rect Recti) {
	rect.fromC(C.sfSprite_getTextureRect(this.cptr))
	return
}

func (this *Sprite) GetColor() (color Color) {
	color.fromC(C.sfSprite_getColor(this.cptr))
	return
}

func (this *Sprite) SetColor(color Color) {
	C.sfSprite_setColor(this.cptr, color.toC())
}

func (this *Sprite) GetTransform() (trans Transform) {
	trans.fromC(C.sfSprite_getTransform(this.cptr))
	return
}

func (this *Sprite) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfSprite_getInverseTransform(this.cptr))
	return
}

func (this *Sprite) GetLocalBounds() (rect Rectf) {
	rect.fromC(C.sfSprite_getLocalBounds(this.cptr))
	return
}

func (this *Sprite) GetGlobalBounds() (rect Rectf) {
	rect.fromC(C.sfSprite_getGlobalBounds(this.cptr))
	return
}

func (this *Sprite) Draw(target RenderTarget, renderStates *RenderStates) {
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawSprite(target.(*RenderWindow).cptr, this.cptr, renderStates.toCPtr())
	case *RenderTexture:
		C.sfRenderTexture_drawSprite(target.(*RenderTexture).cptr, this.cptr, renderStates.toCPtr())
	}
}
