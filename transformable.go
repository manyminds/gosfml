// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, 
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
//			If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/Graphics/Transformable.h> 
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

// Create a new transformable
func NewTransformable() *Transformable {
	transformable := &Transformable{C.sfTransformable_create()}
	runtime.SetFinalizer(transformable, (*Transformable).destroy)

	return transformable
}

// Destroy an existing transformable
func (this *Transformable) destroy() {
	C.sfTransformable_destroy(this.cptr)
	this.cptr = nil
}

// Copy an existing transformable
func (this *Transformable) Copy() *Transformable {
	transformable := &Transformable{C.sfTransformable_copy(this.cptr)}
	runtime.SetFinalizer(transformable, (*Transformable).destroy)
	return transformable
}

// Set the position of a transformable
//
// This function completely overwrites the previous position.
// See Transformable.Move to apply an offset based on the previous position instead.
// The default position of a transformable Transformable object is (0, 0).
//
// 	position: New position
func (this *Transformable) SetPosition(pos Vector2f) {
	C.sfTransformable_setPosition(this.cptr, pos.toC())
}

// Set the scale factors of a transformable
//
// This function completely overwrites the previous scale.
// See Transformable.Scale to add a factor based on the previous scale instead.
// The default scale of a transformable Transformable object is (1, 1).
//
// 	scale: New scale factors
func (this *Transformable) SetScale(scale Vector2f) {
	C.sfTransformable_setScale(this.cptr, scale.toC())
}

// Set the orientation of a transformable
//
// This function completely overwrites the previous rotation.
// See Transformable.Rotate to add an angle based on the previous rotation instead.
// The default rotation of a transformable Transformable object is 0.
//
// 	angle: New rotation, in degrees
func (this *Transformable) SetRotation(rot float32) {
	C.sfTransformable_setRotation(this.cptr, C.float(rot))
}

// Set the local origin of a transformable
//
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a transformable Transformable object is (0, 0).
//
// origin: New origin
func (this *Transformable) SetOrigin(orig Vector2f) {
	C.sfTransformable_setOrigin(this.cptr, orig.toC())
}

// Get the orientation of a transformable
//
// The rotation is always in the range [0, 360].
func (this *Transformable) GetRotation() float32 {
	return float32(C.sfTransformable_getRotation(this.cptr))
}

// Get the position of a transformable
func (this *Transformable) GetPosition() (pos Vector2f) {
	pos.fromC(C.sfTransformable_getPosition(this.cptr))
	return
}

// Get the current scale of a transformable
func (this *Transformable) GetScale() (scale Vector2f) {
	scale.fromC(C.sfTransformable_getScale(this.cptr))
	return
}

// Get the local origin of a transformable
func (this *Transformable) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfTransformable_getOrigin(this.cptr))
	return
}

// Move a transformable by a given offset
//
// This function adds to the current position of the object,
// unlike Transformable.SetPosition which overwrites it.
//
// 	offset: Offset
func (this *Transformable) Move(offset Vector2f) {
	C.sfTransformable_move(this.cptr, offset.toC())
}

// Scale a transformable
//
// This function multiplies the current scale of the object,
// unlike Transformable.SetScale which overwrites it.
//
// 	factors: Scale factors
func (this *Transformable) Scale(factor Vector2f) {
	C.sfTransformable_scale(this.cptr, factor.toC())
}

// Rotate a transformable
//
// This function adds to the current rotation of the object,
// unlike Transformable.SetRotation which overwrites it.
//
// angle: Angle of rotation, in degrees
func (this *Transformable) Rotate(angle float32) {
	C.sfTransformable_rotate(this.cptr, C.float(angle))
}

// Get the combined transform of a transformable
func (this *Transformable) GetTransform() (trans Transform) {
	trans.fromC(C.sfTransformable_getTransform(this.cptr))
	return
}

// Get the inverse of the combined transform of a transformable
func (this *Transformable) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfTransformable_getInverseTransform(this.cptr))
	return
}
