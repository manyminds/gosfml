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

// #include <stdlib.h>
// #include <SFML/Graphics/RenderTexture.h> 
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

func NewRenderTexture(width, height uint, depthbuffer bool) *RenderTexture {
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

func (this *RenderTexture) MapPixelToCoords(pos Vector2i, view *View) (coords Vector2f) {
	coords.fromC(C.sfRenderTexture_mapPixelToCoords(this.cptr, pos.toC(), view.cptr))
	return
}

func (this *RenderTexture) MapCoordsToPixel(pos Vector2f, view *View) (coords Vector2i) {
	coords.fromC(C.sfRenderTexture_mapCoordsToPixel(this.cptr, pos.toC(), view.toCPtr()))
	return
}

func (this *RenderTexture) Draw(drawer Drawer, renderStates *RenderStates) {
	drawer.Draw(this, renderStates)
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
