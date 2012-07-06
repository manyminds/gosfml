/*
	COMPLETE: YES (5.7.2012)
*/

package GoSFML2

// #include <SFML/Graphics.h> 
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

func CreateConvexShape() *ConvexShape {
	shape := &ConvexShape{C.sfConvexShape_create(), nil}
	runtime.SetFinalizer(shape, (*CircleShape).Destroy)
	return shape
}

func (this *ConvexShape) Copy() *ConvexShape {
	shape := &ConvexShape{C.sfConvexShape_copy(this.cptr), this.texture}
	runtime.SetFinalizer(shape, (*CircleShape).Destroy)
	return shape
}

func (this *ConvexShape) Destroy() {
	C.sfConvexShape_destroy(this.cptr)
	this.cptr = nil
}

func (this *ConvexShape) SetPosition(pos Vector2f) {
	C.sfConvexShape_setPosition(this.cptr, pos.toC())
}

func (this *ConvexShape) SetScale(scale Vector2f) {
	C.sfConvexShape_setScale(this.cptr, scale.toC())
}

func (this *ConvexShape) SetOrigin(orig Vector2f) {
	C.sfConvexShape_setOrigin(this.cptr, orig.toC())
}

func (this *ConvexShape) SetRotation(rot float32) {
	C.sfConvexShape_setRotation(this.cptr, C.float(rot))
}

func (this *ConvexShape) GetRotation() float32 {
	return float32(C.sfConvexShape_getRotation(this.cptr))
}

func (this *ConvexShape) GetPosition() (position Vector2f) {
	position.fromC(C.sfConvexShape_getPosition(this.cptr))
	return
}

func (this *ConvexShape) GetScale() (scale Vector2f) {
	scale.fromC(C.sfConvexShape_getScale(this.cptr))
	return
}

func (this *ConvexShape) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfConvexShape_getOrigin(this.cptr))
	return
}

func (this *ConvexShape) Move(offset Vector2f) {
	C.sfConvexShape_move(this.cptr, offset.toC())
}

func (this *ConvexShape) Scale(factor Vector2f) {
	C.sfConvexShape_scale(this.cptr, factor.toC())
}

func (this *ConvexShape) Rotate(angle float32) {
	C.sfConvexShape_rotate(this.cptr, C.float(angle))
}

func (this *ConvexShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfConvexShape_setTexture(this.cptr, texture.cptr, goBool2C(resetRect))
	this.texture = texture
}

func (this *ConvexShape) SetTextureRect(rect Recti) {
	C.sfConvexShape_setTextureRect(this.cptr, rect.toC())
}

func (this *ConvexShape) SetFillColor(color Color) {
	C.sfConvexShape_setFillColor(this.cptr, color.toC())
}

func (this *ConvexShape) SetOutlineColor(color Color) {
	C.sfConvexShape_setOutlineColor(this.cptr, color.toC())
}

func (this *ConvexShape) SetOutlineThickness(thickness float32) {
	C.sfConvexShape_setOutlineThickness(this.cptr, C.float(thickness))
}

func (this *ConvexShape) GetTexture() *Texture {
	return this.texture
}

func (this *ConvexShape) GetTransform() (transform Transform) {
	transform.fromC(C.sfConvexShape_getTransform(this.cptr))
	return
}

func (this *ConvexShape) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfConvexShape_getInverseTransform(this.cptr))
	return
}

func (this *ConvexShape) GetTextureRect() (rect Recti) {
	rect.fromC(C.sfConvexShape_getTextureRect(this.cptr))
	return
}

func (this *ConvexShape) GetFillColor() (color Color) {
	color.fromC(C.sfConvexShape_getFillColor(this.cptr))
	return
}

func (this *ConvexShape) GetOutlineColor() (color Color) {
	color.fromC(C.sfConvexShape_getOutlineColor(this.cptr))
	return
}

func (this *ConvexShape) GetOutlineThickness() float32 {
	return float32(C.sfConvexShape_getOutlineThickness(this.cptr))
}

func (this *ConvexShape) GetPointCount() uint {
	return uint(C.sfConvexShape_getPointCount(this.cptr))
}

func (this *ConvexShape) GetPoint(index uint) (point Vector2f) {
	point.fromC(C.sfConvexShape_getPoint(this.cptr, C.uint(index)))
	return
}

func (this *ConvexShape) SetPointCount(count uint) {
	C.sfConvexShape_setPointCount(this.cptr, C.uint(count))
}

func (this *ConvexShape) SetPoint(index uint, point Vector2f) {
	C.sfConvexShape_setPoint(this.cptr, C.uint(index), point.toC())
}

func (this *ConvexShape) GetLocalBounds() (rect Rectf) {
	rect.fromC(C.sfConvexShape_getLocalBounds(this.cptr))
	return
}

func (this *ConvexShape) GetGlobalBounds() (rect Rectf) {
	rect.fromC(C.sfConvexShape_getGlobalBounds(this.cptr))
	return
}

func (this *ConvexShape) Draw(target RenderTarget, renderStates *RenderStates) {
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawConvexShape(target.(*RenderWindow).cptr, this.cptr, renderStates.toCPtr())
	case *RenderTexture:
		C.sfRenderWindow_drawConvexShape(target.(*RenderTexture).cptr, this.cptr, renderStates.toCPtr())
	}
}
