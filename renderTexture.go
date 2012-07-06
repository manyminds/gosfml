package GoSFML2

/*
 #include <SFML/Graphics.h>
 #include <stdlib.h>
*/
import "C"

import (
	"runtime"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type RenderTexture struct {
	cptr *C.sfRenderTexture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateRenderTexture(width, height uint, depthbuffer bool) *RenderTexture {
	renderTexture := &RenderTexture{C.sfRenderTexture_create(C.uint(width), C.uint(height), goBool2C(depthbuffer))}
	runtime.SetFinalizer(renderTexture, (*RenderTexture).Destroy)
	return renderTexture
}

func (this *RenderTexture) Destroy() {
	C.sfRenderTexture_destroy(this.cptr)
	this.cptr = nil
}

func (this *RenderTexture) GetSize() (size Vector2u) {
	size.fromC(C.sfRenderTexture_getSize(this.cptr))
	return
}

func (this *RenderTexture) SetActive(active bool) {
	C.sfRenderTexture_setActive(this.cptr, goBool2C(active))
}

func (this *RenderTexture) Display() {
	C.sfRenderTexture_display(this.cptr)
}

func (this *RenderTexture) Clear(color Color) {
	C.sfRenderTexture_clear(this.cptr, color.toC())
}

func (this *RenderTexture) SetView(view *View) {
	C.sfRenderTexture_setView(this.cptr, view.cptr)
}

func (this *RenderTexture) GetView() *View {
	return &View{C.sfRenderTexture_getView(this.cptr)}
}

func (this *RenderTexture) GetDefaultView() *View {
	return &View{C.sfRenderTexture_getDefaultView(this.cptr)}
}

func (this *RenderTexture) GetViewport(view *View) (viewport Recti) {
	viewport.fromC(C.sfRenderTexture_getViewport(this.cptr, view.cptr))
	return
}

func (this *RenderTexture) ConvertCoords(pos Vector2i, view *View) (coord Vector2f) {
	coord.fromC(C.sfRenderTexture_convertCoords(this.cptr, pos.toC(), view.cptr))
	return
}

func (this *RenderTexture) Draw(drawable Drawable, renderStates *RenderStates) {
	switch drawable.(type) {
	case *CircleShape:
		C.sfRenderTexture_drawCircleShape(this.cptr, drawable.(*CircleShape).cptr, renderStates.toCPtr())
	case *RectangleShape:
		C.sfRenderTexture_drawCircleShape(this.cptr, drawable.(*RectangleShape).cptr, renderStates.toCPtr())
	case *Sprite:
		C.sfRenderTexture_drawSprite(this.cptr, drawable.(*Sprite).cptr, renderStates.toCPtr())
	case *Text:
		C.sfRenderTexture_drawText(this.cptr, drawable.(*Text).cptr, renderStates.toCPtr())
	case *ConvexShape:
		C.sfRenderTexture_drawConvexShape(this.cptr, drawable.(*ConvexShape).cptr, renderStates.toCPtr())
	case *VertexArray:
		C.sfRenderTexture_drawVertexArray(this.cptr, drawable.(*VertexArray).cptr, renderStates.toCPtr())
	default:
		//invalid shape
	}
}

func (this *RenderTexture) PushGLStates() {
	C.sfRenderTexture_pushGLStates(this.cptr)
}

func (this *RenderTexture) PopGLStates() {
	C.sfRenderTexture_popGLStates(this.cptr)
}

func (this *RenderTexture) ResetGLStates() {
	C.sfRenderTexture_resetGLStates(this.cptr)
}

func (this *RenderTexture) GetTexture() *Texture {
	return &Texture{C.sfRenderTexture_getTexture(this.cptr)}
}

func (this *RenderTexture) SetSmooth(smooth bool) {
	C.sfRenderTexture_setSmooth(this.cptr, goBool2C(smooth))
}

func (this *RenderTexture) IsSmooth() bool {
	return sfBool2Go(C.sfRenderTexture_isSmooth(this.cptr))
}
