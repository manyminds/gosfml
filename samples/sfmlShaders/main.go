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
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	ticker := time.NewTicker(time.Second / 60)

	renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Shaders (GoSFML2)", sf.StyleDefault, sf.DefaultContextSettings())

	// Create the effects
	effects := [...]Effect{&WaveBlur{}, &Pixelate{}, &StormBlink{}, &Edge{}}

	// Initialize them
	for _, eff := range effects {
		eff.Load()
	}

	current := 0

	// Create the messages background
	textBackgroundTexture, _ := sf.NewTextureFromFile("resources/text-background.png", nil)
	textBackground, _ := sf.NewSprite(textBackgroundTexture)
	textBackground.SetPosition(sf.Vector2f{0, 520})
	textBackground.SetColor(sf.Color{255, 255, 255, 200})

	// Load font
	font, _ := sf.NewFontFromFile("resources/sansation.ttf")

	// Create the description text
	description, _ := sf.NewText(font)
	description.SetString("Current effect: " + effects[current].GetName())
	description.SetCharacterSize(20)
	description.SetColor(sf.Color{80, 80, 80, 255})
	description.SetPosition(sf.Vector2f{10, 530})

	// Create the instructions text
	instructions, _ := sf.NewText(font)
	instructions.SetString("Press left and right arrows to change the current shader")
	instructions.SetCharacterSize(20)
	instructions.SetColor(sf.Color{80, 80, 80, 255})
	instructions.SetPosition(sf.Vector2f{280, 555})

	var timeAccu time.Duration = 0
	var mousePos sf.Vector2i

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			//poll events
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch ev := event.(type) {
				case sf.EventKeyReleased:
					switch ev.Code {
					case sf.KeyEscape:
						renderWindow.Close()
					case sf.KeyLeft:
						if current > 0 {
							current--
						} else {
							current = len(effects) - 1
						}
						description.SetString("Current effect: " + effects[current].GetName())
					case sf.KeyRight:
						if current < len(effects)-1 {
							current++
						} else {
							current = 0
						}
						description.SetString("Current effect: " + effects[current].GetName())
					}
				case sf.EventMouseMoved:
					mousePos = sf.Vector2i{ev.X, ev.Y}
				case sf.EventClosed:
					renderWindow.Close()
				}
			}

			timeAccu += time.Second / 60

			// Clear the window
			renderWindow.Clear(sf.Color{255, 128, 0, 255})

			// render effect
			convPos := renderWindow.MapPixelToCoords(mousePos, nil)
			x := convPos.X / float32(renderWindow.GetSize().X)
			y := convPos.Y / float32(renderWindow.GetSize().Y)

			effects[current].Update(float32(timeAccu.Seconds()), x, y)

			effects[current].Draw(renderWindow)

			renderWindow.Draw(textBackground, sf.DefaultRenderStates())
			renderWindow.Draw(instructions, sf.DefaultRenderStates())
			renderWindow.Draw(description, sf.DefaultRenderStates())

			// Display things on screen
			renderWindow.Display()
		}
	}
}
