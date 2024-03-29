// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

// The best source on android input events is the NDK: include/android/input.h
//
// iOS event handling guide:
// https://developer.apple.com/library/ios/documentation/EventHandling/Conceptual/EventHandlingiPhoneOS

// Touch is a touch event.
type Touch struct {
	Base

	// Sequence is the sequence number. The same number is shared by all events
	// in a sequence. A sequence begins with a single Begin, is followed by
	// zero or more Moves, and ends with a single End. A Sequence
	// distinguishes concurrent sequences but its value is subsequently reused.
	Sequence Sequence
}

// Sequence identifies a sequence of touch events.
type Sequence int64

// todo: this is half-baked -- need real case.  should have position.

func NewTouch(typ Types, seq Sequence) *Touch {
	ev := &Touch{}
	ev.Typ = typ
	ev.SetUnique()
	ev.Sequence = seq
	return ev
}

func (ev *Touch) HasPos() bool {
	return true
}

// todo: what about these higher-level abstractions of touch-like events?

// // MagnifyEvent is used to represent a magnification gesture.
// type MagnifyEvent struct {
// 	GestureEvent
// 	Magnification float64 // the multiplicative scale factor
// }

// func (ev *MagnifyEvent) EventTypes() EventTypes {
// 	return MagnifyEventTypes
// }

// // check for interface implementation
// var _ Event = &MagnifyEvent{}

// ////////////////////////////////////////////

// // RotateEvent is used to represent a rotation gesture.
// type RotateEvent struct {
// 	GestureEvent
// 	Rotation float64 // measured in degrees; positive == clockwise
// }

// func (ev *RotateEvent) EventTypes() EventTypes {
// 	return RotateEventTypes
// }

// // check for interface implementation
// var _ Event = &RotateEvent{}

// // Scroll Event is used to represent a scrolling gesture.
// type ScrollEvent struct {
// 	GestureEvent
// 	Delta image.Point
// }

// func (ev *ScrollEvent) EventTypes() EventTypes {
// 	return ScrollEventTypes
// }

// // check for interface implementation
// var _ Event = &ScrollEvent{}
