// Copyright (c) 2021 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package osevent defines operating system level events, not associated
// with a particular window.
package osevent

//go:generate enumgen

import (
	"fmt"
	"image"

	"goki.dev/gi/v2/oswin"
)

// osevent.Event reports an OS level event
type Event struct {
	oswin.EventBase

	// Action taken on the osevent -- what has changed.  Osevent state fields
	// have current values.
	Action Actions
}

// Actions is the action taken on the osevent by the user.
type Actions int32 //enums:enum

const (
	// OpenFiles means the user indicated that the app should open file(s) stored in Files
	OpenFiles Actions = iota
)

/////////////////////////////
// oswin.Event interface

func (ev *Event) Type() oswin.EventType {
	return oswin.OSEvent
}

func (ev *Event) HasPos() bool {
	return false
}

func (ev *Event) Pos() image.Point {
	return image.Point{}
}

func (ev *Event) OnWinFocus() bool { // os events generally not focus-specific
	return false
}

func (ev *Event) OnFocus() bool {
	return false
}

func (ev *Event) String() string {
	return fmt.Sprintf("Type: %v Action: %v  Time: %v", ev.Type(), ev.Action, ev.Time())
}

// osevent.OpenFilesEvent is for OS open files action to open given files
type OpenFilesEvent struct {
	Event

	// Files are a list of files to open
	Files []string
}

func (ev *OpenFilesEvent) Type() oswin.EventType {
	return oswin.OSOpenFilesEvent
}