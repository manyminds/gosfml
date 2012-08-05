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

type Transformable struct {
	cptr *C.sfTransformable
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewTransformable(tex *Texture) *Transformable {
	transformable := &Transformable{C.sfTransformable_create()}
	runtime.SetFinalizer(transformable, (*Transformable).Destroy)

	return transformable
}

func (this *Transformable) Destroy() {
	C.sfTransformable_destroy(this.cptr)
	this.cptr = nil
}

func (this *Transformable) Copy() *Transformable {
	transformable := &Transformable{C.sfTransformable_copy(this.cptr)}
	runtime.SetFinalizer(transformable, (*Transformable).Destroy)
	return transformable
}

func (this *Transformable) SetPosition(pos Vector2f) {
	C.sfTransformable_setPosition(this.cptr, pos.toC())
}

func (this *Transformable) SetScale(scale Vector2f) {
	C.sfTransformable_setScale(this.cptr, scale.toC())
}

func (this *Transformable) SetRotation(rot float32) {
	C.sfTransformable_setRotation(this.cptr, C.float(rot))
}

func (this *Transformable) SetOrigin(orig Vector2f) {
	C.sfTransformable_setOrigin(this.cptr, orig.toC())
}

func (this *Transformable) GetRotation() float32 {
	return float32(C.sfTransformable_getRotation(this.cptr))
}

func (this *Transformable) GetPosition() (pos Vector2f) {
	pos.fromC(C.sfTransformable_getPosition(this.cptr))
	return
}

func (this *Transformable) GetScale() (scale Vector2f) {
	scale.fromC(C.sfTransformable_getScale(this.cptr))
	return
}

func (this *Transformable) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfTransformable_getOrigin(this.cptr))
	return
}

func (this *Transformable) Move(offset Vector2f) {
	C.sfTransformable_move(this.cptr, offset.toC())
}

func (this *Transformable) Scale(factor Vector2f) {
	C.sfTransformable_scale(this.cptr, factor.toC())
}

func (this *Transformable) Rotate(angle float32) {
	C.sfTransformable_rotate(this.cptr, C.float(angle))
}

func (this *Transformable) GetTransform() (trans Transform) {
	trans.fromC(C.sfTransformable_getTransform(this.cptr))
	return
}

func (this *Transformable) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfTransformable_getInverseTransform(this.cptr))
	return
}
