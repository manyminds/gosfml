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
)

type EventType int

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

type Event interface {
	GetType() EventType
}

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

func (this *KeyEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	SizeEvent

type SizeEvent struct {
	EventType EventType
	Width     uint
	Height    uint
}

func (this *SizeEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	TextEvent

type TextEvent struct {
	EventType EventType
	Char      uint32
}

func (this *TextEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	MouseMoveEvent

type MouseMoveEvent struct {
	EventType EventType
	X         int
	Y         int
}

func (this *MouseMoveEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	MouseButtonEvent

type MouseButtonEvent struct {
	EventType EventType
	Button    MouseButton
	X         int
	Y         int
}

func (this *MouseButtonEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	MouseWheelEvent

type MouseWheelEvent struct {
	EventType EventType
	Delta     int
	X         int
	Y         int
}

func (this *MouseWheelEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	JoystickMoveEvent

type JoystickMoveEvent struct {
	EventType  EventType
	JoystickId uint
	Axis       JoystickAxis
	position   float32
}

func (this *JoystickMoveEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	JoystickButtonEvent

type JoystickButtonEvent struct {
	EventType  EventType
	JoystickId uint
	Button     uint
}

func (this *JoystickButtonEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	JoystickConnectEvent

type JoystickConnectEvent struct {
	EventType  EventType
	joystickId uint
}

func (this *JoystickConnectEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//	RawEvent

//20 bytes
type RawEvent struct {
	EventType EventType
	data      [16]byte
}

func (this *RawEvent) GetType() EventType {
	return this.EventType
}

///////////////////////////////////////////////////////////////
//standard event handling method used by Window & RenderWindow

func HandleEvent(cEvent *RawEvent) Event {
	eventType := cEvent.GetType()

	switch eventType {
	case Event_Closed:
		return (*RawEvent)(unsafe.Pointer(cEvent))
	case Event_Resized:
		return (*SizeEvent)(unsafe.Pointer(cEvent))
	case Event_TextEntered:
		return (*TextEvent)(unsafe.Pointer(cEvent))
	case Event_KeyPressed:
		return (*KeyEvent)(unsafe.Pointer(cEvent))
	case Event_KeyReleased:
		return (*KeyEvent)(unsafe.Pointer(cEvent))
	case Event_MouseWheelMoved:
		return (*MouseWheelEvent)(unsafe.Pointer(cEvent))
	case Event_MouseButtonPressed:
		fallthrough
	case Event_MouseButtonReleased:
		return (*MouseButtonEvent)(unsafe.Pointer(cEvent))
	case Event_MouseMoved:
		fallthrough
	case Event_MouseEntered:
		fallthrough
	case Event_MouseLeft:
		return (*MouseMoveEvent)(unsafe.Pointer(cEvent))
	case Event_JoystickButtonPressed:
		fallthrough
	case Event_JoystickButtonReleased:
		fallthrough
	case Event_JoystickMoved:
		fallthrough
	case Event_JoystickConnected:
		fallthrough
	case Event_JoystickDisconnected:
		fallthrough
	default:
		return (*RawEvent)(unsafe.Pointer(cEvent))
	}

	//shouldn't get here
	return (*RawEvent)(unsafe.Pointer(cEvent))
}
