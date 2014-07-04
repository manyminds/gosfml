// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

//Sprite, CircleShape, ConvexShape, RectangleShape, Text and VertexArray are Drawers
//A Drawer can be drawn on a RenderTarget
type Drawer interface {
	Draw(target RenderTarget, renderStates RenderStates)
}

/////////////////////////////////////
///		TEST
/////////////////////////////////////

var _ Drawer = (*Sprite)(nil)
var _ Drawer = (*CircleShape)(nil)
var _ Drawer = (*ConvexShape)(nil)
var _ Drawer = (*RectangleShape)(nil)
var _ Drawer = (*Text)(nil)
var _ Drawer = (*VertexArray)(nil)
