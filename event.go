// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications,
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software.
//			If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/Window.h>
// int getEventType(sfEvent* ev) { return ev->type; }
// sfSizeEvent* getSizeEvent(sfEvent* ev) { return &ev->size; }
// sfKeyEvent* getKeyEvent(sfEvent* ev) { return &ev->key; }
// sfTextEvent* getTextEvent(sfEvent* ev) { return &ev->text; }
// sfMouseMoveEvent* getMouseMoveEvent(sfEvent* ev) { return &ev->mouseMove; }
// sfMouseButtonEvent* getMouseButtonEvent(sfEvent* ev) { return &ev->mouseButton; }
// sfMouseWheelEvent* getMouseWheelEvent(sfEvent* ev) { return &ev->mouseWheel; }
// sfJoystickMoveEvent* getJoystickMoveEvent(sfEvent* ev) { return &ev->joystickMove; }
// sfJoystickButtonEvent* getJoystickButtonEvent(sfEvent* ev) { return &ev->joystickButton; }
// sfJoystickConnectEvent* getJoystickConnectEvent(sfEvent* ev) { return &ev->joystickConnect; }
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	eventClosed = iota
	eventResized
	eventLostFocus
	eventGainedFocus
	eventTextEntered
	eventKeyPressed
	eventKeyReleased
	eventMouseWheelMoved
	eventMouseButtonPressed
	eventMouseButtonReleased
	eventMouseMoved
	eventMouseEntered
	eventMouseLeft
	eventJoystickButtonPressed
	eventJoystickButtonReleased
	eventJoystickMoved
	eventJoystickConnected
	eventJoystickDisconnected
	eventNone
)

type eventType int

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

type Event interface{}

///////////////////////////////////////////////////////////////
//	EmptyEvents

// The window lost the focus (no data)
type EventLostFocus struct{}

// The window gained the focus (no data)
type EventGainedFocus struct{}

// The mouse cursor entered the area of the window (no data)
type EventMouseEntered struct{}

// The mouse cursor left the area of the window (no data)
type EventMouseLeft struct{}

// The window requested to be closed (no data)
type EventClosed struct{}

///////////////////////////////////////////////////////////////
//	KeyEvent

type eventKey struct {
	Code    KeyCode //< Code of the key that has been pressed
	Alt     int     //< Is the Alt key pressed?
	Control int     //< Is the Control key pressed?
	Shift   int     //< Is the Shift key pressed?
	System  int     //< Is the System key pressed?
}

type EventKeyPressed eventKey
type EventKeyReleased eventKey

func newKeyEventFromC(ev *C.sfKeyEvent) eventKey {
	return eventKey{Code: KeyCode(ev.code), Alt: int(ev.alt), Control: int(ev.control), Shift: int(ev.shift), System: int(ev.system)}
}

///////////////////////////////////////////////////////////////
//	SizeEvent

type EventResized struct {
	Width  uint //< New width, in pixels
	Height uint //< New height, in pixels
}

func newSizeEventFromC(ev *C.sfSizeEvent) EventResized {
	return EventResized{Width: uint(ev.width), Height: uint(ev.height)}
}

///////////////////////////////////////////////////////////////
//	TextEvent

type EventTextEntered struct {
	Char rune //< Value of the rune
}

func newTextEventFromC(ev *C.sfTextEvent) EventTextEntered {
	return EventTextEntered{Char: rune(ev.unicode)}
}

///////////////////////////////////////////////////////////////
//	MouseMoveEvent

type EventMouseMoved struct {
	X int //< X position of the mouse pointer, relative to the left of the owner window
	Y int //< Y position of the mouse pointer, relative to the top of the owner window
}

func newMouseMoveEventFromC(ev *C.sfMouseMoveEvent) EventMouseMoved {
	return EventMouseMoved{X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	MouseButtonEvent

type eventMouseButton struct {
	Button MouseButton //< Code of the button that has been pressed
	X      int         //< X position of the mouse pointer, relative to the left of the owner window
	Y      int         //< Y position of the mouse pointer, relative to the top of the owner window
}

type EventMouseButtonPressed eventMouseButton
type EventMouseButtonReleased eventMouseButton

func newMouseButtonEventFromC(ev *C.sfMouseButtonEvent) eventMouseButton {
	return eventMouseButton{Button: MouseButton(ev.button), X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	MouseWheelEvent

type EventMouseWheelMoved struct {
	Delta int //< Number of ticks the wheel has moved (positive is up, negative is down)
	X     int //< X position of the mouse pointer, relative to the left of the owner window
	Y     int //< Y position of the mouse pointer, relative to the top of the owner window
}

func newMouseWheelEventFromC(ev *C.sfMouseWheelEvent) EventMouseWheelMoved {
	return EventMouseWheelMoved{Delta: int(ev.delta), X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	JoystickMoveEvent

type EventJoystickMoved struct {
	JoystickId uint         //< Index of the joystick (in range [0 .. JoystickCount - 1])
	Axis       JoystickAxis //< Axis on which the joystick moved
	position   float32      //< New position on the axis (in range [-100 .. 100])
}

func newJoystickMoveEventFromC(ev *C.sfJoystickMoveEvent) EventJoystickMoved {
	return EventJoystickMoved{JoystickId: uint(ev.joystickId), Axis: JoystickAxis(ev.axis), position: float32(ev.position)}
}

///////////////////////////////////////////////////////////////
//	JoystickButtonEvent

type eventJoystickButton struct {
	JoystickId uint //< Index of the joystick (in range [0 .. JoystickCount - 1])
	Button     uint //< Index of the button that has been pressed (in range [0 .. JoystickButtonCount - 1])
}

type EventJoystickButtonPressed eventJoystickButton
type EventJoystickButtonReleased eventJoystickButton

func newJoystickButtonEventFromC(ev *C.sfJoystickButtonEvent) eventJoystickButton {
	return eventJoystickButton{JoystickId: uint(ev.joystickId), Button: uint(ev.button)}
}

///////////////////////////////////////////////////////////////
//	JoystickConnectEvent

type eventJoystickConnection struct {
	JoystickId uint //< Index of the joystick (in range [0 .. JoystickCount - 1])
}

type EventJoystickConnected eventJoystickConnection
type EventJoystickDisconnected eventJoystickConnection

func newJoystickConnectEventFromC(ev *C.sfJoystickConnectEvent) eventJoystickConnection {
	return eventJoystickConnection{JoystickId: uint(ev.joystickId)}
}

///////////////////////////////////////////////////////////////
//standard event handling method used by Window & RenderWindow

func handleEvent(cEvent *C.sfEvent) (ev Event) {
	switch eventType(C.getEventType(cEvent)) {
	case eventResized:
		ev = newSizeEventFromC(C.getSizeEvent(cEvent))
	case eventClosed:
		ev = EventClosed{}
	case eventLostFocus:
		ev = EventLostFocus{}
	case eventGainedFocus:
		ev = EventGainedFocus{}
	case eventTextEntered:
		ev = newTextEventFromC(C.getTextEvent(cEvent))
	case eventKeyReleased:
		ev = (EventKeyReleased)(newKeyEventFromC(C.getKeyEvent(cEvent)))
	case eventKeyPressed:
		ev = (EventKeyPressed)(newKeyEventFromC(C.getKeyEvent(cEvent)))
	case eventMouseWheelMoved:
		ev = newMouseWheelEventFromC(C.getMouseWheelEvent(cEvent))
	case eventMouseButtonReleased:
		ev = (EventMouseButtonReleased)(newMouseButtonEventFromC(C.getMouseButtonEvent(cEvent)))
	case eventMouseButtonPressed:
		ev = (EventMouseButtonPressed)(newMouseButtonEventFromC(C.getMouseButtonEvent(cEvent)))
	case eventMouseMoved:
		ev = newMouseMoveEventFromC(C.getMouseMoveEvent(cEvent))
	case eventMouseLeft:
		ev = EventMouseLeft{}
	case eventMouseEntered:
		ev = EventMouseEntered{}
	case eventJoystickButtonReleased:
		ev = (EventJoystickButtonReleased)(newJoystickButtonEventFromC(C.getJoystickButtonEvent(cEvent)))
	case eventJoystickButtonPressed:
		ev = (EventJoystickButtonPressed)(newJoystickButtonEventFromC(C.getJoystickButtonEvent(cEvent)))
	case eventJoystickMoved:
		ev = newJoystickMoveEventFromC(C.getJoystickMoveEvent(cEvent))
	case eventJoystickDisconnected:
		ev = (EventJoystickDisconnected)(newJoystickConnectEventFromC(C.getJoystickConnectEvent(cEvent)))
	case eventJoystickConnected:
		ev = (EventJoystickConnected)(newJoystickConnectEventFromC(C.getJoystickConnectEvent(cEvent)))
	}
	return
}
