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

// #include <SFML/Graphics/VertexArray.h> 
// #include <SFML/Graphics/RenderWindow.h> 
// #include <SFML/Graphics/RenderTexture.h> 
import "C"
import "runtime"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Primitive_Points         = iota ///< List of individual points
	Primitive_Lines                 ///< List of individual lines
	Primitive_LinesStrip            ///< List of connected lines, a point uses the previous point to form a line
	Primitive_Triangles             ///< List of individual triangles
	Primitive_TrianglesStrip        ///< List of connected triangles, a point uses the two previous points to form a triangle
	Primitive_TrianglesFran         ///< List of connected triangles, a point uses the common center and the previous point to form a triangle
	Primitive_Quads                 ///< List of individual quads
)

type PrimitiveType int

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type VertexArray struct {
	cptr *C.sfVertexArray
}

type Vertex struct {
	Position  Vector2f
	Color     Color
	TexCoords Vector2f
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new vertex array
func NewVertexArray() *VertexArray {
	vertexArray := &VertexArray{C.sfVertexArray_create()}
	runtime.SetFinalizer(vertexArray, (*VertexArray).Destroy)
	return vertexArray
}

// Copy an existing vertex array
func (this *VertexArray) Copy() *VertexArray {
	vertexArray := &VertexArray{C.sfVertexArray_copy(this.cptr)}
	runtime.SetFinalizer(vertexArray, (*VertexArray).Destroy)
	return vertexArray
}

// Destroy an existing vertex array
func (this *VertexArray) Destroy() {
	C.sfVertexArray_destroy(this.cptr)
	this.cptr = nil
}

// Return the vertex count of a vertex array
func (this *VertexArray) GetVertexCount() uint {
	return uint(C.sfVertexArray_getVertexCount(this.cptr))
}

// Get access to a vertex by its index
//
// This function doesn't check index, it must be in range
// [0, vertex count - 1]. The behaviour is undefined
// otherwise.
func (this *VertexArray) GetVertex(index uint) (vert Vertex) {
	vert.fromC(*C.sfVertexArray_getVertex(this.cptr, C.uint(index)))
	return
}

// Sets a vertex by its index
//
// This function doesn't check index, it must be in range
// [0, vertex count - 1]. The behaviour is undefined
// otherwise.
func (this *VertexArray) SetVertex(vertex Vertex, index uint) {
	cVert := C.sfVertexArray_getVertex(this.cptr, C.uint(index))
	cVert.position = vertex.Position.toC()
	cVert.color = vertex.Color.toC()
	cVert.texCoords = vertex.TexCoords.toC()
}

// Clear a vertex array
//
// This function removes all the vertices from the array.
// It doesn't deallocate the corresponding memory, so that
// adding new vertices after clearing doesn't involve
// reallocating all the memory.
func (this *VertexArray) Clear() {
	C.sfVertexArray_clear(this.cptr)
}

// Resize the vertex array
//
// If vertexCount is greater than the current size, the previous
// vertices are kept and new (default-constructed) vertices are
// added.
// If vertexCount is less than the current size, existing vertices
// are removed from the array.
//
// vertexCount: New size of the array (number of vertices)
func (this *VertexArray) Resize(vertexCount uint) {
	C.sfVertexArray_resize(this.cptr, C.uint(vertexCount))
}

// Add a vertex to a vertex array array
//
// vertex: Vertex to add
func (this *VertexArray) Append(vertex Vertex) {
	C.sfVertexArray_append(this.cptr, vertex.toC())
}

// Set the type of primitives of a vertex array
//
// This function defines how the vertices must be interpreted
// when it's time to draw them:
// As points
// As lines
// As triangles
// As quads
// The default primitive type is Points.
//
// type: Type of primitive
func (this *VertexArray) SetPrimitiveType(ptype PrimitiveType) {
	C.sfVertexArray_setPrimitiveType(this.cptr, C.sfPrimitiveType(ptype))
}

// Get the type of primitives drawn by a vertex array
func (this *VertexArray) GetPrimitiveType() PrimitiveType {
	return PrimitiveType(C.sfVertexArray_getPrimitiveType(this.cptr))
}

// Compute the bounding rectangle of a vertex array
//
// This function returns the axis-aligned rectangle that
// contains all the vertices of the array
func (this *VertexArray) GetBounds() (rect FloatRect) {
	rect.fromC(C.sfVertexArray_getBounds(this.cptr))
	return
}

func (this *VertexArray) Draw(target RenderTarget, renderStates *RenderStates) {
	switch target.(type) {
	case *RenderWindow:
		C.sfRenderWindow_drawVertexArray(target.(*RenderWindow).cptr, this.cptr, renderStates.toCPtr())
	case *RenderTexture:
		C.sfRenderTexture_drawVertexArray(target.(*RenderTexture).cptr, this.cptr, renderStates.toCPtr())
	}
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Vertex) fromC(vertex C.sfVertex) {
	this.Position.fromC(vertex.position)
	this.Color.fromC(vertex.color)
	this.TexCoords.fromC(vertex.texCoords)
}

func (this *Vertex) toC() C.sfVertex {
	return C.sfVertex{position: this.Position.toC(), color: this.Color.toC(), texCoords: this.TexCoords.toC()}
}
