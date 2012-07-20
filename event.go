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
// int getSizeEvent() { return sizeof(sfEvent); }
import "C"
import "unsafe"

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
	EventType EventType
	Code      KeyCode
	Alt       int
	Control   int
	Shift     int
	System    int
}

///////////////////////////////////////////////////////////////
//	SizeEvent

type SizeEvent struct {
	EventType EventType
	Width     uint
	Height    uint
}

///////////////////////////////////////////////////////////////
//	TextEvent

type TextEvent struct {
	EventType EventType
	Char      uint32
}

///////////////////////////////////////////////////////////////
//	MouseMoveEvent

type MouseMoveEvent struct {
	EventType EventType
	X         int
	Y         int
}

///////////////////////////////////////////////////////////////
//	MouseButtonEvent

type MouseButtonEvent struct {
	EventType EventType
	Button    MouseButton
	X         int
	Y         int
}

///////////////////////////////////////////////////////////////
//	MouseWheelEvent

type MouseWheelEvent struct {
	EventType EventType
	Delta     int
	X         int
	Y         int
}

///////////////////////////////////////////////////////////////
//	JoystickMoveEvent

type JoystickMoveEvent struct {
	EventType  EventType
	JoystickId uint
	Axis       JoystickAxis
	position   float32
}

///////////////////////////////////////////////////////////////
//	JoystickButtonEvent

type JoystickButtonEvent struct {
	EventType  EventType
	JoystickId uint
	Button     uint
}

///////////////////////////////////////////////////////////////
//	JoystickConnectEvent

type JoystickConnectEvent struct {
	EventType  EventType
	joystickId uint
}

///////////////////////////////////////////////////////////////
//	RawEvent

type RawEvent struct {
	EventType EventType
	data      [20]byte
}

///////////////////////////////////////////////////////////////
//standard event handling method used by Window & RenderWindow

func HandleEvent(cEvent *RawEvent) (ev Event, evt EventType) {
	evt = cEvent.EventType

	switch evt {
	case Event_Closed:
	case Event_Resized:
		ev = (*SizeEvent)(unsafe.Pointer(cEvent))
	case Event_TextEntered:
		ev = (*TextEvent)(unsafe.Pointer(cEvent))
	case Event_KeyReleased, Event_KeyPressed:
		ev = (*KeyEvent)(unsafe.Pointer(cEvent))
	case Event_MouseWheelMoved:
		ev = (*MouseWheelEvent)(unsafe.Pointer(cEvent))
	case Event_MouseButtonReleased, Event_MouseButtonPressed:
		ev = (*MouseButtonEvent)(unsafe.Pointer(cEvent))
	case Event_MouseLeft, Event_MouseEntered, Event_MouseMoved:
		ev = (*MouseMoveEvent)(unsafe.Pointer(cEvent))
	case Event_JoystickButtonReleased, Event_JoystickButtonPressed:
		ev = (*JoystickButtonEvent)(unsafe.Pointer(cEvent))
	case Event_JoystickMoved:
		ev = (*JoystickMoveEvent)(unsafe.Pointer(cEvent))
	case Event_JoystickDisconnected, Event_JoystickConnected:
		ev = (*JoystickConnectEvent)(unsafe.Pointer(cEvent))
	default:
	}
	return
}

/////////////////////////////////////
///		Testing
/////////////////////////////////////

func sizeofEvent() int {
	return int(C.getSizeEvent())
}
