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

var _ Drawer = &Sprite{}
var _ Drawer = &CircleShape{}
var _ Drawer = &ConvexShape{}
var _ Drawer = &RectangleShape{}
var _ Drawer = &Text{}
var _ Drawer = &VertexArray{}
