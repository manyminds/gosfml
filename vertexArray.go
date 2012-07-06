/*
	COMPLETE: YES (6.7.2012)
*/

package GoSFML2

// #include <SFML/Graphics.h> 
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

func CreateVertexArray() *VertexArray {
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

func (this *VertexArray) GetBounds() (rect Rectf) {
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
