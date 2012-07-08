/*
	COMPLETE: YES (4.7.2012)
*/

package GoSFML2

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
	if tex != nil {
		shape.SetTexture(tex, true)
	}

	return shape
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
	C.sfSprite_setTexture(this.cptr, texture.cptr, goBool2C(resetRect))
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