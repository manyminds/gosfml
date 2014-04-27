// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

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

type EventType int

const (
	EventTypeClosed                 EventType = C.sfEvtClosed
	EventTypeResized                EventType = C.sfEvtResized
	EventTypeLostFocus              EventType = C.sfEvtLostFocus
	EventTypeGainedFocus            EventType = C.sfEvtGainedFocus
	EventTypeTextEntered            EventType = C.sfEvtTextEntered
	EventTypeKeyPressed             EventType = C.sfEvtKeyPressed
	EventTypeKeyReleased            EventType = C.sfEvtKeyReleased
	EventTypeMouseWheelMoved        EventType = C.sfEvtMouseWheelMoved
	EventTypeMouseButtonPressed     EventType = C.sfEvtMouseButtonPressed
	EventTypeMouseButtonReleased    EventType = C.sfEvtMouseButtonReleased
	EventTypeMouseMoved             EventType = C.sfEvtMouseMoved
	EventTypeMouseEntered           EventType = C.sfEvtMouseEntered
	EventTypeMouseLeft              EventType = C.sfEvtMouseLeft
	EventTypeJoystickButtonPressed  EventType = C.sfEvtJoystickButtonPressed
	EventTypeJoystickButtonReleased EventType = C.sfEvtJoystickButtonReleased
	EventTypeJoystickMoved          EventType = C.sfEvtJoystickMoved
	EventTypeJoystickConnected      EventType = C.sfEvtJoystickConnected
	EventTypeJoystickDisconnected   EventType = C.sfEvtJoystickDisconnected
)

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

type Event interface {
	Type() EventType
}

///////////////////////////////////////////////////////////////
//	EmptyEvents

// The window lost the focus (no data)
type EventLostFocus struct{}

func (EventLostFocus) Type() EventType {
	return EventTypeLostFocus
}

// The window gained the focus (no data)
type EventGainedFocus struct{}

func (EventGainedFocus) Type() EventType {
	return EventTypeGainedFocus
}

// The mouse cursor entered the area of the window (no data)
type EventMouseEntered struct{}

func (EventMouseEntered) Type() EventType {
	return EventTypeMouseEntered
}

// The mouse cursor left the area of the window (no data)
type EventMouseLeft struct{}

func (EventMouseLeft) Type() EventType {
	return EventTypeMouseLeft
}

// The window requested to be closed (no data)
type EventClosed struct{}

func (EventClosed) Type() EventType {
	return EventTypeClosed
}

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

func (EventKeyPressed) Type() EventType {
	return EventTypeKeyPressed
}

func (EventKeyReleased) Type() EventType {
	return EventTypeKeyReleased
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

func (EventResized) Type() EventType {
	return EventTypeResized
}

///////////////////////////////////////////////////////////////
//	TextEvent

type EventTextEntered struct {
	Char rune //< Value of the rune
}

func newTextEventFromC(ev *C.sfTextEvent) EventTextEntered {
	return EventTextEntered{Char: rune(uint32(ev.unicode))}
}

func (EventTextEntered) Type() EventType {
	return EventTypeTextEntered
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

func (EventMouseMoved) Type() EventType {
	return EventTypeMouseMoved
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

func (EventMouseButtonPressed) Type() EventType {
	return EventTypeMouseButtonPressed
}

func (EventMouseButtonReleased) Type() EventType {
	return EventTypeMouseButtonReleased
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

func (EventMouseWheelMoved) Type() EventType {
	return EventTypeMouseWheelMoved
}

///////////////////////////////////////////////////////////////
//	JoystickMoveEvent

type EventJoystickMoved struct {
	JoystickId uint         //< Index of the joystick (in range [0 .. JoystickCount - 1])
	Axis       JoystickAxis //< Axis on which the joystick moved
	Position   float32      //< New position on the axis (in range [-100 .. 100])
}

func newJoystickMoveEventFromC(ev *C.sfJoystickMoveEvent) EventJoystickMoved {
	return EventJoystickMoved{JoystickId: uint(ev.joystickId), Axis: JoystickAxis(ev.axis), Position: float32(ev.position)}
}

func (EventJoystickMoved) Type() EventType {
	return EventTypeJoystickMoved
}

///////////////////////////////////////////////////////////////
//	JoystickButtonEvent

type eventJoystickButton struct {
	JoystickId uint //< Index of the joystick (in range [0 .. JoystickCount - 1])
	Button     uint //< Index of the button that has been pressed (in range [0 .. JoystickButtonCount - 1])
}

func newJoystickButtonEventFromC(ev *C.sfJoystickButtonEvent) eventJoystickButton {
	return eventJoystickButton{JoystickId: uint(ev.joystickId), Button: uint(ev.button)}
}

type EventJoystickButtonPressed eventJoystickButton
type EventJoystickButtonReleased eventJoystickButton

func (EventJoystickButtonPressed) Type() EventType {
	return EventTypeJoystickButtonPressed
}

func (EventJoystickButtonReleased) Type() EventType {
	return EventTypeJoystickButtonReleased
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

func (EventJoystickConnected) Type() EventType {
	return EventTypeJoystickConnected
}

func (EventJoystickDisconnected) Type() EventType {
	return EventTypeJoystickDisconnected
}

///////////////////////////////////////////////////////////////
//standard event handling method used by Window & RenderWindow

func handleEvent(cEvent *C.sfEvent) (ev Event) {
	switch EventType(C.getEventType(cEvent)) {
	case EventTypeResized:
		ev = newSizeEventFromC(C.getSizeEvent(cEvent))
	case EventTypeClosed:
		ev = EventClosed{}
	case EventTypeLostFocus:
		ev = EventLostFocus{}
	case EventTypeGainedFocus:
		ev = EventGainedFocus{}
	case EventTypeTextEntered:
		ev = newTextEventFromC(C.getTextEvent(cEvent))
	case EventTypeKeyReleased:
		ev = (EventKeyReleased)(newKeyEventFromC(C.getKeyEvent(cEvent)))
	case EventTypeKeyPressed:
		ev = (EventKeyPressed)(newKeyEventFromC(C.getKeyEvent(cEvent)))
	case EventTypeMouseWheelMoved:
		ev = newMouseWheelEventFromC(C.getMouseWheelEvent(cEvent))
	case EventTypeMouseButtonReleased:
		ev = (EventMouseButtonReleased)(newMouseButtonEventFromC(C.getMouseButtonEvent(cEvent)))
	case EventTypeMouseButtonPressed:
		ev = (EventMouseButtonPressed)(newMouseButtonEventFromC(C.getMouseButtonEvent(cEvent)))
	case EventTypeMouseMoved:
		ev = newMouseMoveEventFromC(C.getMouseMoveEvent(cEvent))
	case EventTypeMouseLeft:
		ev = EventMouseLeft{}
	case EventTypeMouseEntered:
		ev = EventMouseEntered{}
	case EventTypeJoystickButtonReleased:
		ev = (EventJoystickButtonReleased)(newJoystickButtonEventFromC(C.getJoystickButtonEvent(cEvent)))
	case EventTypeJoystickButtonPressed:
		ev = (EventJoystickButtonPressed)(newJoystickButtonEventFromC(C.getJoystickButtonEvent(cEvent)))
	case EventTypeJoystickMoved:
		ev = newJoystickMoveEventFromC(C.getJoystickMoveEvent(cEvent))
	case EventTypeJoystickDisconnected:
		ev = (EventJoystickDisconnected)(newJoystickConnectEventFromC(C.getJoystickConnectEvent(cEvent)))
	case EventTypeJoystickConnected:
		ev = (EventJoystickConnected)(newJoystickConnectEventFromC(C.getJoystickConnectEvent(cEvent)))
	default:
		panic("Unknown event")
	}
	return
}
