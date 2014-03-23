/*
#############################################
#	GOSFML2
#	Events
#############################################
*/

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"runtime"
	"strconv"
	"time"
	"unicode"
)

func init() {
	runtime.LockOSThread()
}

/////////////////////////////////////
///		LOGGER
/////////////////////////////////////

type Logger []*sf.Text

func (logger Logger) PushBack(msg string) {
	for i := 0; i < len(logger)-1; i++ {
		oldMsg := logger[i+1].GetString()
		logger[i].SetString(oldMsg)
	}
	logger[len(logger)-1].SetString(msg)
}

/////////////////////////////////////
///		MAIN
/////////////////////////////////////

func main() {

	ticker := time.NewTicker(time.Second / 30)

	renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Events (GoSFML2)", sf.StyleDefault, sf.DefaultContextSettings())

	//load font
	font, _ := sf.NewFontFromFile("resources/Vera.ttf")

	text, _ := sf.NewText(font)
	text.SetColor(sf.ColorBlack())
	text.SetPosition(sf.Vector2f{80, 100})
	text.SetString("Move your mouse and press some keys")

	//logger
	const NumberOfItems = 20
	logger := make(Logger, NumberOfItems)
	for i := 0; i < NumberOfItems; i++ {
		logger[i], _ = sf.NewText(font)
		logger[i].SetColor(sf.ColorBlack())
		logger[i].SetPosition(sf.Vector2f{100, 150 + float32(i)*20})
		logger[i].SetCharacterSize(12)
	}

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			//poll events
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch ev := event.(type) {
				case sf.EventKeyPressed:
					logger.PushBack("Key pressed: " + strconv.Itoa(int(ev.Code)) + " Shift: " + strconv.Itoa(int(ev.Shift)) + " Control: " +
						strconv.Itoa(int(ev.Control)) + " Alt: " + strconv.Itoa(int(ev.Alt)) + " System: " + strconv.Itoa(int(ev.System)))

					//exit on ESC
					if ev.Code == sf.KeyEscape {
						renderWindow.Close()
					}
				case sf.EventKeyReleased:
					logger.PushBack("Key released: " + strconv.Itoa(int(ev.Code)) + " Shift: " + strconv.Itoa(int(ev.Shift)) + " Control: " +
						strconv.Itoa(int(ev.Control)) + " Alt: " + strconv.Itoa(int(ev.Alt)) + " System: " + strconv.Itoa(int(ev.System)))
				case sf.EventGainedFocus:
					logger.PushBack("Gained Focus")
				case sf.EventLostFocus:
					logger.PushBack("Lost Focus")
				case sf.EventResized:
					logger.PushBack("Resized width: " + strconv.Itoa(int(ev.Width)) + " height: " + strconv.Itoa(int(ev.Height)))
				case sf.EventTextEntered:
					if unicode.IsPrint(ev.Char) {
						logger.PushBack("Text entered: " + string(ev.Char))
					}
				case sf.EventMouseButtonPressed:
					logger.PushBack("Mouse pressed: " + strconv.Itoa(int(ev.Button)) + " [X: " + strconv.Itoa(ev.X) + " Y: " + strconv.Itoa(ev.Y) + "]")
				case sf.EventMouseLeft:
					logger.PushBack("Mouse left")
				case sf.EventMouseEntered:
					logger.PushBack("Mouse entered")
				case sf.EventMouseWheelMoved:
					logger.PushBack("Mouse wheel moved: " + strconv.Itoa(int(ev.Delta)))
				case sf.EventMouseMoved:
					logger.PushBack("Mouse moved: [X: " + strconv.Itoa(ev.X) + " Y: " + strconv.Itoa(ev.Y) + "]")
				case sf.EventClosed:
					renderWindow.Close()
				case sf.EventJoystickConnected:
					logger.PushBack("Joystick connected")
				case sf.EventJoystickDisconnected:
					logger.PushBack("Joystick disconnected")
				case sf.EventJoystickButtonPressed:
					logger.PushBack("Joystick Button pressed: " + strconv.Itoa(int(ev.Button)))
				case sf.EventJoystickButtonReleased:
					logger.PushBack("Joystick Button released: " + strconv.Itoa(int(ev.Button)))
				case sf.EventJoystickMoved:
					logger.PushBack("Joystick moved: [Axis" + strconv.Itoa(int(ev.Axis)) + "Value: " + strconv.Itoa(int(sf.JoystickGetAxisPosition(ev.JoystickId, ev.Axis))) + "]")
				}
			}
		}

		renderWindow.Clear(sf.ColorWhite())
		renderWindow.Draw(text, sf.DefaultRenderStates())
		for _, t := range logger {
			renderWindow.Draw(t, sf.DefaultRenderStates())
		}
		renderWindow.Display()
	}
}
