// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Vertex.h>
// #include <SFML/Graphics/RenderWindow.h>
// #include <SFML/Graphics/RenderTexture.h>
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	PrimitivePoints         PrimitiveType = C.sfPoints         ///< List of individual points
	PrimitiveLines          PrimitiveType = C.sfLines          ///< List of individual lines
	PrimitiveLinesStrip     PrimitiveType = C.sfLinesStrip     ///< List of connected lines, a point uses the previous point to form a line
	PrimitiveTriangles      PrimitiveType = C.sfTriangles      ///< List of individual triangles
	PrimitiveTrianglesStrip PrimitiveType = C.sfTrianglesStrip ///< List of connected triangles, a point uses the two previous points to form a triangle
	PrimitiveTrianglesFan   PrimitiveType = C.sfTrianglesFan   ///< List of connected triangles, a point uses the common center and the previous point to form a triangle
	PrimitiveQuads          PrimitiveType = C.sfQuads          ///< List of individual quads
)

type PrimitiveType int

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type VertexArray struct {
	Vertices      []Vertex
	PrimitiveType PrimitiveType
}

type Vertex struct {
	//Created by cgo -godefs - DO NOT EDIT
	Position  Vector2f
	Color     Color
	TexCoords Vector2f
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new vertex array
func NewVertexArray() (*VertexArray, error) {
	vertexArray := &VertexArray{}
	return vertexArray, nil
}

// Copy an existing vertex array
func (this *VertexArray) Copy() *VertexArray {
	vertexArray := &VertexArray{Vertices: make([]Vertex, len(this.Vertices)), PrimitiveType: this.PrimitiveType}
	copy(vertexArray.Vertices, this.Vertices)

	return vertexArray
}

// Return the vertex count of a vertex array
func (this *VertexArray) GetVertexCount() int {
	return len(this.Vertices)
}

// Clear a vertex array
//
// This function removes all the vertices from the array.
// It doesn't deallocate the corresponding memory, so that
// adding new vertices after clearing doesn't involve
// reallocating all the memory.
func (this *VertexArray) Clear() {
	this.Vertices = this.Vertices[:0]
}

// Resize the vertex array
//
// If vertexCount is greater than the current size, the previous
// vertices are kept and new (default-constructed i.e. Vertex.Color is set to white)
// vertices are added.
// If vertexCount is less than the current size, existing vertices
// are removed from the array.
//
// 	vertexCount: New size of the array (number of vertices)
func (this *VertexArray) Resize(vertexCount int) {
	if vertexCount > len(this.Vertices) {
		i := len(this.Vertices)

		//copy current vertices
		vertices := make([]Vertex, vertexCount)
		copy(vertices, this.Vertices)
		this.Vertices = vertices

		//newly added vertices are white (matches SFML behaviour)
		for ; i < len(this.Vertices); i++ {
			this.Vertices[i].Color = ColorWhite()
		}
	} else {
		//resize
		this.Vertices = this.Vertices[:vertexCount]
	}
}

// Add a vertex to a vertex array array
//
// 	vertex: Vertex to add
//
// Note: You may want to specify the vertex color - otherwise the newly added vertex will be invisible
//
//	example: vertexArray.Append(Vertex{Position: Vector2f{}, Color: ColorWhite})
func (this *VertexArray) Append(vertex Vertex) {
	this.Vertices = append(this.Vertices, vertex)
}

// Compute the bounding rectangle of a vertex array
//
// This function returns the axis-aligned rectangle that
// contains all the vertices of the array
func (this *VertexArray) GetBounds() FloatRect {
	if len(this.Vertices) > 0 {
		left := this.Vertices[0].Position.X
		top := this.Vertices[0].Position.Y
		right := this.Vertices[0].Position.X
		bottom := this.Vertices[0].Position.Y

		for i := 1; i < len(this.Vertices); i++ {
			pos := this.Vertices[i].Position

			if pos.X < left {
				left = pos.X
			} else if pos.X > right {
				right = pos.X
			}

			if pos.Y < top {
				top = pos.Y
			} else if pos.Y > bottom {
				bottom = pos.Y
			}
		}

		return FloatRect{left, top, right - left, bottom - top}
	}

	return FloatRect{}
}

// Draws a VertexArray on a render target
//
// Note: Use RenderTarget.DrawPrimitives to draw only a subset of its vertices
func (this *VertexArray) Draw(target RenderTarget, renderStates RenderStates) {
	if target != nil {
		target.DrawPrimitives(this.Vertices, this.PrimitiveType, renderStates)
	}
}
