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
	PrimitivePoints         = iota ///< List of individual points
	PrimitiveLines                 ///< List of individual lines
	PrimitiveLinesStrip            ///< List of connected lines, a point uses the previous point to form a line
	PrimitiveTriangles             ///< List of individual triangles
	PrimitiveTrianglesStrip        ///< List of connected triangles, a point uses the two previous points to form a triangle
	PrimitiveTrianglesFran         ///< List of connected triangles, a point uses the common center and the previous point to form a triangle
	PrimitiveQuads                 ///< List of individual quads
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

func NewVertexArray() *VertexArray {
	vertexArray := &VertexArray{C.sfVertexArray_create()}
	runtime.SetFinalizer(vertexArray, (*VertexArray).Destroy)
	return vertexArray
}

func (this *VertexArray) Copy() *VertexArray {
	vertexArray := &VertexArray{C.sfVertexArray_copy(this.cptr)}
	runtime.SetFinalizer(vertexArray, (*VertexArray).Destroy)
	return vertexArray
}

func (this *VertexArray) Destroy() {
	C.sfVertexArray_destroy(this.cptr)
	this.cptr = nil
}

func (this *VertexArray) GetVertexCount() uint {
	return uint(C.sfVertexArray_getVertexCount(this.cptr))
}

func (this *VertexArray) GetVertex(index uint) (vert Vertex) {
	vert.fromC(*C.sfVertexArray_getVertex(this.cptr, C.uint(index)))
	return
}

func (this *VertexArray) SetVertex(vertex Vertex, index uint) {
	cVert := C.sfVertexArray_getVertex(this.cptr, C.uint(index))
	cVert.position = vertex.Position.toC()
	cVert.color = vertex.Color.toC()
	cVert.texCoords = vertex.TexCoords.toC()
}

func (this *VertexArray) Clear() {
	C.sfVertexArray_clear(this.cptr)
}

func (this *VertexArray) Resize(vertexCount uint) {
	C.sfVertexArray_resize(this.cptr, C.uint(vertexCount))
}

func (this *VertexArray) Append(vertex Vertex) {
	C.sfVertexArray_append(this.cptr, vertex.toC())
}

func (this *VertexArray) SetPrimitiveType(ptype PrimitiveType) {
	C.sfVertexArray_setPrimitiveType(this.cptr, C.sfPrimitiveType(ptype))
}

func (this *VertexArray) GetPrimitiveType() PrimitiveType {
	return PrimitiveType(C.sfVertexArray_getPrimitiveType(this.cptr))
}

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
