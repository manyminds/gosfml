// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

//RenderTexture and RenderWindow are RenderTargets
type RenderTarget interface {
	Clear(Color)
	Display()
	SetView(*View)
	GetView() *View
	GetDefaultView() *View
	GetViewport(view *View) IntRect
	MapPixelToCoords(Vector2i, *View) Vector2f
	MapCoordsToPixel(Vector2f, *View) Vector2i
	PushGLStates()
	PopGLStates()
	ResetGLStates()
	GetSize() Vector2u
	Draw(Drawer, RenderStates)
	DrawPrimitives([]Vertex, PrimitiveType, RenderStates)
}

/////////////////////////////////////
///		TEST
/////////////////////////////////////

var _ RenderTarget = &RenderTexture{}
var _ RenderTarget = &RenderWindow{}
