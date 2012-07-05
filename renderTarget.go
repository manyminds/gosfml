package GoSFML2

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

//implemented by RenderTexture and RenderWindow
type RenderTarget interface {
	Clear(Color)
	Display()
	SetView(*View)
	GetView() *View
	GetDefaultView() *View
	GetViewport(view *View) Recti
	ConvertCoords(Vector2i, *View) Vector2f
	PushGLStates()
	PopGLStates()
	ResetGLStates()
	GetSize() Vector2u
	Draw(Drawable, *RenderStates)
}
