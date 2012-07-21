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

package GoSFML2

// #include <SFML/Window.h>
// int getEventType(sfEvent* ev) { return ev->type; }
// struct sfSizeEvent* getSizeEvent(sfEvent* ev) { return &ev->size; }
// struct sfKeyEvent* getKeyEvent(sfEvent* ev) { return &ev->key; }
// struct sfTextEvent* getTextEvent(sfEvent* ev) { return &ev->text; }
// struct sfMouseMoveEvent* getMouseMoveEvent(sfEvent* ev) { return &ev->mouseMove; }
// struct sfMouseButtonEvent* getMouseButtonEvent(sfEvent* ev) { return &ev->mouseButton; }
// struct sfMouseWheelEvent* getMouseWheelEvent(sfEvent* ev) { return &ev->mouseWheel; }
// struct sfJoystickMoveEvent* getJoystickMoveEvent(sfEvent* ev) { return &ev->joystickMove; }
// struct sfJoystickButtonEvent* getJoystickButtonEvent(sfEvent* ev) { return &ev->joystickButton; }
// struct sfJoystickConnectEvent* getJoystickConnectEvent(sfEvent* ev) { return &ev->joystickConnect; }
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Event_Closed = iota
	Event_Resized
	Event_LostFocus
	Event_GainedFocus
	Event_TextEntered
	Event_KeyPressed
	Event_KeyReleased
	Event_MouseWheelMoved
	Event_MouseButtonPressed
	Event_MouseButtonReleased
	Event_MouseMoved
	Event_MouseEntered
	Event_MouseLeft
	Event_JoystickButtonPressed
	Event_JoystickButtonReleased
	Event_JoystickMoved
	Event_JoystickConnected
	Event_JoystickDisconnected
	Event_Error
)

type EventType int

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

type Event interface{}

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type KeyEvent struct {
	Code    KeyCode
	Alt     int
	Control int
	Shift   int
	System  int
}

func newKeyEventFromC(ev *C.struct_sfKeyEvent) *KeyEvent {
	return &KeyEvent{Code: KeyCode(ev.code), Alt: int(ev.alt), Control: int(ev.control), Shift: int(ev.shift), System: int(ev.system)}
}

///////////////////////////////////////////////////////////////
//	SizeEvent

type SizeEvent struct {
	Width  uint
	Height uint
}

func newSizeEventFromC(ev *C.struct_sfSizeEvent) *SizeEvent {
	return &SizeEvent{Width: uint(ev.width), Height: uint(ev.height)}
}

///////////////////////////////////////////////////////////////
//	TextEvent

type TextEvent struct {
	Char rune
}

func newTextEventFromC(ev *C.struct_sfTextEvent) *TextEvent {
	return &TextEvent{Char: rune(ev.unicode)}
}

///////////////////////////////////////////////////////////////
//	MouseMoveEvent

type MouseMoveEvent struct {
	EventType EventType
	X         int
	Y         int
}

func newMouseMoveEventFromC(ev *C.struct_sfMouseMoveEvent) *MouseMoveEvent {
	return &MouseMoveEvent{X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	MouseButtonEvent

type MouseButtonEvent struct {
	Button MouseButton
	X      int
	Y      int
}

func newMouseButtonEventFromC(ev *C.struct_sfMouseButtonEvent) *MouseButtonEvent {
	return &MouseButtonEvent{Button: MouseButton(ev.button), X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	MouseWheelEvent

type MouseWheelEvent struct {
	Delta int
	X     int
	Y     int
}

func newMouseWheelEventFromC(ev *C.struct_sfMouseWheelEvent) *MouseWheelEvent {
	return &MouseWheelEvent{Delta: int(ev.delta), X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	JoystickMoveEvent

type JoystickMoveEvent struct {
	JoystickId uint
	Axis       JoystickAxis
	position   float32
}

func newJoystickMoveEventFromC(ev *C.struct_sfJoystickMoveEvent) *JoystickMoveEvent {
	return &JoystickMoveEvent{JoystickId: uint(ev.joystickId), Axis: JoystickAxis(ev.axis), position: float32(ev.position)}
}

///////////////////////////////////////////////////////////////
//	JoystickButtonEvent

type JoystickButtonEvent struct {
	JoystickId uint
	Button     uint
}

func newJoystickButtonEventFromC(ev *C.struct_sfJoystickButtonEvent) *JoystickButtonEvent {
	return &JoystickButtonEvent{JoystickId: uint(ev.joystickId), Button: uint(ev.button)}
}

///////////////////////////////////////////////////////////////
//	JoystickConnectEvent

type JoystickConnectEvent struct {
	JoystickId uint
}

func newJoystickConnectEventFromC(ev *C.struct_sfJoystickConnectEvent) *JoystickConnectEvent {
	return &JoystickConnectEvent{JoystickId: uint(ev.joystickId)}
}

///////////////////////////////////////////////////////////////
//standard event handling method used by Window & RenderWindow

func handleEvent(cEvent *C.sfEvent) (ev Event, evt EventType) {
	evt = EventType(C.getEventType(cEvent))

	switch evt {
	case Event_Closed:
	case Event_Resized:
		ev = newSizeEventFromC(C.getSizeEvent(cEvent))
	case Event_TextEntered:
		ev = newTextEventFromC(C.getTextEvent(cEvent))
	case Event_KeyReleased, Event_KeyPressed:
		ev = newKeyEventFromC(C.getKeyEvent(cEvent))
	case Event_MouseWheelMoved:
		ev = newMouseWheelEventFromC(C.getMouseWheelEvent(cEvent))
	case Event_MouseButtonReleased, Event_MouseButtonPressed:
		ev = newMouseButtonEventFromC(C.getMouseButtonEvent(cEvent))
	case Event_MouseLeft, Event_MouseEntered, Event_MouseMoved:
		ev = newMouseMoveEventFromC(C.getMouseMoveEvent(cEvent))
	case Event_JoystickButtonReleased, Event_JoystickButtonPressed:
		ev = newJoystickButtonEventFromC(C.getJoystickButtonEvent(cEvent))
	case Event_JoystickMoved:
		ev = newJoystickMoveEventFromC(C.getJoystickMoveEvent(cEvent))
	case Event_JoystickDisconnected, Event_JoystickConnected:
		ev = newJoystickConnectEventFromC(C.getJoystickConnectEvent(cEvent))
	}
	return
}
