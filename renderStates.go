package GoSFML2

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

type RenderState struct {
	BlendMode BlendMode
	Transform Transform
	Texture   Texture
}
