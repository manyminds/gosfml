package GoSFML2

// #include <SFML/Graphics.h> 
import "C"

import (
	"runtime"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type View struct {
	cptr *C.sfView
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateView() *View {
	view := &View{C.sfView_create()}
	runtime.SetFinalizer(view, (*View).Destroy)
	return view
}

func CreateViewFromRect(rect *Rectf) *View {
	view := &View{C.sfView_createFromRect(rect.toC())}
	runtime.SetFinalizer(view, (*View).Destroy)
	return view
}

func (this *View) Destroy() {
	C.sfView_destroy(this.cptr)
	this.cptr = nil
}

func (this *View) Copy() *View {
	view := &View{C.sfView_copy(this.cptr)}
	runtime.SetFinalizer(view, (*View).Destroy)
	return view
}

func (this *View) SetCenter(center *Vector2f) {
	C.sfView_setCenter(this.cptr, center.toC())
}

func (this *View) SetSize(size *Vector2f) {
	C.sfView_setSize(this.cptr, size.toC())
}

func (this *View) SetRotation(rotation float32) {
	C.sfView_setRotation(this.cptr, C.float(rotation))
}

func (this *View) SetViewport(viewport *Rectf) {
	C.sfView_setViewport(this.cptr, viewport.toC())
}

func (this *View) Reset(rect *Rectf) {
	C.sfView_reset(this.cptr, rect.toC())
}

func (this *View) GetCenter() (center Vector2f) {
	center.fromC(C.sfView_getCenter(this.cptr))
	return
}

func (this *View) GetSize() (size Vector2f) {
	size.fromC(C.sfView_getSize(this.cptr))
	return
}

func (this *View) GetRotation() float32 {
	return float32(C.sfView_getRotation(this.cptr))
}

func (this *View) GetViewport() (rect Rectf) {
	rect.fromC(C.sfView_getViewport(this.cptr))
	return
}

func (this *View) Move(offset *Vector2f) {
	C.sfView_move(this.cptr, offset.toC())
}

func (this *View) Rotate(angle float32) {
	C.sfView_rotate(this.cptr, C.float(angle))
}

func (this *View) Zoom(factor float32) {
	C.sfView_zoom(this.cptr, C.float(factor))
}
