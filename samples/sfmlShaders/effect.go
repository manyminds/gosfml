/*
#############################################
#	GOSFML2
#	SFML Example: Shaders
#	Ported from C++ to Go
#############################################
*/

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
	"math/rand"
)

type Effect interface {
	GetName() string
	Load()
	Update(time, x, y float32)
	Draw(target sf.RenderTarget)
}

//////////////////////////////////////////
//		SHADERS
//////////////////////////////////////////

//////////////////////////////////////////
//		WaveBlur

type WaveBlur struct {
	shader *sf.Shader
	text   *sf.Text
}

func (this *WaveBlur) GetName() string {
	return "wave + blur"
}

func (this *WaveBlur) Load() {
	font, _ := sf.NewFontFromFile("resources/sansation.ttf")
	this.text, _ = sf.NewText(font)
	this.text.SetString("Praesent suscipit augue in velit pulvinar hendrerit varius purus aliquam.\n" +
		"Mauris mi odio, bibendum quis fringilla a, laoreet vel orci. Proin vitae vulputate tortor.\n" +
		"Praesent cursus ultrices justo, ut feugiat ante vehicula quis.\n" +
		"Donec fringilla scelerisque mauris et viverra.\n" +
		"Maecenas adipiscing ornare scelerisque. Nullam at libero elit.\n" +
		"Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.\n" +
		"Nullam leo urna, tincidunt id semper eget, ultricies sed mi.\n" +
		"Morbi mauris massa, commodo id dignissim vel, lobortis et elit.\n" +
		"Fusce vel libero sed neque scelerisque venenatis.\n" +
		"Integer mattis tincidunt quam vitae iaculis.\n" +
		"Vivamus fringilla sem non velit venenatis fermentum.\n" +
		"Vivamus varius tincidunt nisi id vehicula.\n" +
		"Integer ullamcorper, enim vitae euismod rutrum, massa nisl semper ipsum,\n" +
		"vestibulum sodales sem ante in massa.\n" +
		"Vestibulum in augue non felis convallis viverra.\n" +
		"Mauris ultricies dolor sed massa convallis sed aliquet augue fringilla.\n" +
		"Duis erat eros, porta in accumsan in, blandit quis sem.\n" +
		"In hac habitasse platea dictumst. Etiam fringilla est id odio dapibus sit amet semper dui laoreet.\n")

	this.text.SetCharacterSize(22)
	this.text.SetPosition(sf.Vector2f{30, 20})

	// Load the shader
	this.shader, _ = sf.NewShaderFromFile("resources/wave.vert", "resources/blur.frag")
}

func (this *WaveBlur) Update(time, x, y float32) {
	this.shader.SetFloatParameter("wave_phase", time)
	this.shader.SetFloatParameter("wave_amplitude", x*40, y*40)
	this.shader.SetFloatParameter("blur_radius", (x+y)*0.008)
}

func (this *WaveBlur) Draw(target sf.RenderTarget) {
	target.Draw(this.text, sf.RenderStates{BlendMode: sf.BlendAlpha, Transform: sf.TransformIdentity(), Shader: this.shader, Texture: nil})
}

//////////////////////////////////////////
//		Pixelate

type Pixelate struct {
	shader *sf.Shader
	sprite *sf.Sprite
}

func (this *Pixelate) GetName() string {
	return "pixelate"
}

func (this *Pixelate) Load() {
	// Load the texture and initialize the sprite
	texture, _ := sf.NewTextureFromFile("resources/background.jpg", nil)
	this.sprite, _ = sf.NewSprite(texture)

	// Load the shader
	this.shader, _ = sf.NewShaderFromFile("", "resources/pixelate.frag")
}

func (this *Pixelate) Update(time, x, y float32) {
	this.shader.SetFloatParameter("pixel_threshold", (x+y)/30)
}

func (this *Pixelate) Draw(target sf.RenderTarget) {
	target.Draw(this.sprite, sf.RenderStates{BlendMode: sf.BlendAlpha, Transform: sf.TransformIdentity(), Shader: this.shader, Texture: nil})
}

//////////////////////////////////////////
//		Storm

type StormBlink struct {
	shader   *sf.Shader
	vertices *sf.VertexArray
}

func (this *StormBlink) GetName() string {
	return "storm + blink"
}

func (this *StormBlink) Load() {
	// Create the points
	this.vertices, _ = sf.NewVertexArray()
	for i := 0; i < 40000; i++ {
		x := rand.Float32() * 800
		y := rand.Float32() * 600

		r := byte(rand.Float32() * 255)
		g := byte(rand.Float32() * 255)
		b := byte(rand.Float32() * 255)

		vertex := sf.Vertex{Position: sf.Vector2f{x, y}, Color: sf.Color{r, g, b, 255}, TexCoords: sf.Vector2f{0, 0}}

		this.vertices.Append(vertex)
	}

	// Load the shader
	this.shader, _ = sf.NewShaderFromFile("resources/storm.vert", "resources/blink.frag")
}

func (this *StormBlink) Update(time, x, y float32) {
	radius := float32(200 + math.Cos(float64(time))*150)

	this.shader.SetFloatParameter("storm_position", x*800, y*600)
	this.shader.SetFloatParameter("storm_inner_radius", radius/3)
	this.shader.SetFloatParameter("storm_total_radius", radius)
	this.shader.SetFloatParameter("blink_alpha", 0.5+float32(math.Cos(float64(time*3)))*0.25)
}

func (this *StormBlink) Draw(target sf.RenderTarget) {
	target.Draw(this.vertices, sf.RenderStates{BlendMode: sf.BlendAlpha, Transform: sf.TransformIdentity(), Shader: this.shader, Texture: nil})
}

//////////////////////////////////////////
//		Edge

type Edge struct {
	shader            *sf.Shader
	surface           *sf.RenderTexture
	surfaceSprite     *sf.Sprite
	backgroundTexture *sf.Texture
	entityTexture     *sf.Texture
	backgroundSprite  *sf.Sprite
	entities          []*sf.Sprite
}

func (this *Edge) GetName() string {
	return "edge post-effect"
}

func (this *Edge) Load() {
	// Create the off-screen surface
	this.surface, _ = sf.NewRenderTexture(800, 600, false)
	this.surface.SetSmooth(false)

	// This sprite is used to render the off-screen surface
	this.surfaceSprite, _ = sf.NewSprite(this.surface.GetTexture())

	// Load the textures
	this.backgroundTexture, _ = sf.NewTextureFromFile("resources/sfml.png", nil)
	this.backgroundTexture.SetSmooth(true)
	this.entityTexture, _ = sf.NewTextureFromFile("resources/devices.png", nil)
	this.entityTexture.SetSmooth(true)

	// Initialize the background sprite
	this.backgroundSprite, _ = sf.NewSprite(this.backgroundTexture)
	this.backgroundSprite.SetPosition(sf.Vector2f{135, 100})

	// Load the moving entities
	for i := 0; i < 6; i++ {
		entity, _ := sf.NewSprite(nil)
		entity.SetTexture(this.entityTexture, false)
		entity.SetTextureRect(sf.IntRect{96 * i, 0, 96, 96})
		this.entities = append(this.entities, entity)
	}

	// Load the shader
	this.shader, _ = sf.NewShaderFromFile("", "resources/edge.frag")
	this.shader.SetCurrentTextureParameter("texture")
}

func (this *Edge) Update(time, x, y float32) {
	this.shader.SetFloatParameter("edge_threshold", 1-(x+y)/2)

	// Update the position of the moving entities
	for i := 0; i < len(this.entities); i++ {
		x := float32(math.Cos(float64(0.25*(time*float32(i)+float32(len(this.entities)-i))))*300 + 350)
		y := float32(math.Sin(float64(0.25*(time*float32(i)+float32(len(this.entities)-i))))*200 + 250)
		position := sf.Vector2f{x, y}
		this.entities[i].SetPosition(position)
	}

	// Render the updated scene to the off-screen surface
	this.surface.Clear(sf.ColorWhite())
	this.surface.Draw(this.backgroundSprite, sf.DefaultRenderStates())
	for i := 0; i < len(this.entities); i++ {
		this.surface.Draw(this.entities[i], sf.DefaultRenderStates())
	}
	this.surface.Display()
}

func (this *Edge) Draw(target sf.RenderTarget) {
	target.Draw(this.surfaceSprite, sf.RenderStates{BlendMode: sf.BlendAlpha, Transform: sf.TransformIdentity(), Shader: this.shader, Texture: nil})
}
