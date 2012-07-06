package GoSFML2

/*
 #include <SFML/Graphics.h> 
 #include <stdlib.h>
 #include <stddef.h>
*/
import "C"
import "runtime"
import "unsafe"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Text_Regular    = 0      ///< Regular characters, no style
	Text_Bold       = 1 << 0 ///< Characters are bold
	Text_Italic     = 1 << 1 ///< Characters are in italic
	Text_Underlined = 1 << 2 ///< Characters are underlined
)

type TextStyle uint32

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Text struct {
	cptr *C.sfFont
	font *Font
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateText() *Text {
	text := &Text{C.sfText_create(), nil}
	runtime.SetFinalizer(text, (*Text).Destroy)
	return text
}

func (this *Text) Destroy() {
	C.sfText_destroy(this.cptr)
	this.cptr = nil
}

func (this *Text) Copy() *Text {
	return &Text{C.sfText_copy(this.cptr), this.font}
}

func (this *Text) SetPosition(pos Vector2f) {
	C.sfText_setPosition(this.cptr, pos.toC())
}

func (this *Text) SetScale(scale Vector2f) {
	C.sfText_setScale(this.cptr, scale.toC())
}

func (this *Text) SetOrigin(orig Vector2f) {
	C.sfText_setOrigin(this.cptr, orig.toC())
}

func (this *Text) SetRotation(rot float32) {
	C.sfText_setRotation(this.cptr, C.float(rot))
}

func (this *Text) Move(offset Vector2f) {
	C.sfText_move(this.cptr, offset.toC())
}

func (this *Text) Scale(factor Vector2f) {
	C.sfText_scale(this.cptr, factor.toC())
}

func (this *Text) Rotate(angle float32) {
	C.sfText_rotate(this.cptr, C.float(angle))
}

func (this *Text) GetRotation() float32 {
	return float32(C.sfText_getRotation(this.cptr))
}

func (this *Text) GetPosition() (pos Vector2f) {
	pos.fromC(C.sfText_getPosition(this.cptr))
	return
}

func (this *Text) GetScale() (scale Vector2f) {
	scale.fromC(C.sfText_getScale(this.cptr))
	return
}

func (this *Text) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfText_getOrigin(this.cptr))
	return
}

func (this *Text) GetTransform() (trans Transform) {
	trans.fromC(C.sfText_getTransform(this.cptr))
	return
}

func (this *Text) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfText_getInverseTransform(this.cptr))
	return
}

func (this *Text) SetString(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.sfText_setString(this.cptr, cText)
}

func (this *Text) SetFont(font *Font) {
	C.sfText_setFont(this.cptr, font.cptr)
	this.font = font
}

func (this *Text) SetCharacterSize(size uint) {
	C.sfText_setCharacterSize(this.cptr, C.uint(size))
}

func (this *Text) SetStyle(style TextStyle) {
	C.sfText_setStyle(this.cptr, C.sfUint32(style))
}

func (this *Text) SetColor(color Color) {
	C.sfText_setColor(this.cptr, color.toC())
}

func (this *Text) GetString() string {
	cstr := C.sfText_getString(this.cptr)
	return C.GoString(cstr)
}

func (this *Text) GetFont() *Font {
	return this.font
}

func (this *Text) GetCharacterSize() uint {
	return uint(C.sfText_getCharacterSize(this.cptr))
}

func (this *Text) GetStyle() TextStyle {
	return TextStyle(C.sfText_getStyle(this.cptr))
}

func (this *Text) GetColor() (color Color) {
	color.fromC(C.sfText_getColor(this.cptr))
	return
}

func (this *Text) FintCharacterPos(index uint) (pos Vector2f) {
	pos.fromC(C.sfText_findCharacterPos(this.cptr, C.size_t(index)))
	return
}

func (this *Text) GetLocalBounds() (rect Rectf) {
	rect.fromC(C.sfSprite_getLocalBounds(this.cptr))
	return
}

func (this *Text) GetGlobalBounds() (rect Rectf) {
	rect.fromC(C.sfSprite_getGlobalBounds(this.cptr))
	return
}

func (this *Text) Draw(target RenderTarget, renderStates *RenderStates) {
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawText(target.(*RenderWindow).cptr, this.cptr, renderStates.toCPtr())
	case *RenderTexture:
		C.sfRenderWindow_drawText(target.(*RenderTexture).cptr, this.cptr, renderStates.toCPtr())
	}
}
