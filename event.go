package GoSFML2

// #include <SFML/Window.h>
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

type SizeEvent struct {
	EventType EventType
	Width     uint
	Height    uint
}

func (this *SizeEvent) GetType() EventType {
	return this.EventType
}

type TextEvent struct {
	EventType EventType
	Char      uint32
}

func (this *TextEvent) GetType() EventType {
	return this.EventType
}

type MouseMoveEvent struct {
	EventType EventType
	X         int
	Y         int
}

func (this *MouseMoveEvent) GetType() EventType {
	return this.EventType
}

type MouseButtonEvent struct {
	EventType EventType
	Button    MouseButton
	X         int
	Y         int
}

func (this *MouseButtonEvent) GetType() EventType {
	return this.EventType
}

type MouseWheelEvent struct {
	EventType EventType
	Delta     int
	X         int
	Y         int
}

func (this *MouseWheelEvent) GetType() EventType {
	return this.EventType
}

type JoystickMoveEvent struct {
	EventType  EventType
	JoystickId uint
	Axis       JoystickAxis
	position   float32
}

func (this *JoystickMoveEvent) GetType() EventType {
	return this.EventType
}

type JoystickButtonEvent struct {
	EventType  EventType
	JoystickId uint
	Button     uint
}

func (this *JoystickButtonEvent) GetType() EventType {
	return this.EventType
}

type JoystickConnectEvent struct {
	EventType  EventType
	joystickId uint
}

func (this *JoystickConnectEvent) GetType() EventType {
	return this.EventType
}

//20 bytes
type RawEvent struct {
	EventType EventType
	data      [16]byte
}

func (this *RawEvent) GetType() EventType {
	return this.EventType
}
