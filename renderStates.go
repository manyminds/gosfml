package GoSFML2

/*
 #include <SFML/Graphics.h>
*/
import "C"
import "unsafe"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Blend_Alpha    = iota ///< Pixel = Src * a + Dest * (1 - a)
	Blend_Add             ///< Pixel = Src + Dest
	Blend_Multiply        ///< Pixel = Src * Dest
	Blend_None            ///< No blending
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type BlendMode int

type RenderStates struct {
	blendMode BlendMode
	transform Transform
	texture   *C.sfTexture
	shader    *C.sfShader
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func NewRenderStates(blendMode BlendMode, transform Transform, texture Texture, shader Shader) (rt RenderStates) {
	rt.blendMode = blendMode
	rt.transform = transform
	rt.shader = shader.cptr
	rt.texture = texture.cptr
	return
}

// shader can be nil
func (this *RenderStates) SetShader(shader *Shader) {
	if shader == nil {
		this.shader = nil
	} else {
		this.shader = shader.cptr
	}
}

// texture can be nil
func (this *RenderStates) SetTexture(texture *Texture) {
	if texture == nil {
		this.texture = nil
	} else {
		this.texture = texture.cptr
	}
}

func (this *RenderStates) SetTramsform(transform Transform) {
	this.transform = transform
}

func (this *RenderStates) SetBlendMode(blendMode BlendMode) {
	this.blendMode = blendMode
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *RenderStates) toCPtr() *C.sfRenderStates {
	return (*C.sfRenderStates)(unsafe.Pointer(this))
}
