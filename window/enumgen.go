// Code generated by "enumgen"; DO NOT EDIT.

package window

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _ActionsValues = []Actions{0, 1, 2, 3, 4, 5, 6, 7, 8}

// ActionsN is the highest valid value
// for type Actions, plus one.
const ActionsN Actions = 9

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ActionsNoOp() {
	var x [1]struct{}
	_ = x[Close-(0)]
	_ = x[Minimize-(1)]
	_ = x[Resize-(2)]
	_ = x[Move-(3)]
	_ = x[Focus-(4)]
	_ = x[DeFocus-(5)]
	_ = x[Paint-(6)]
	_ = x[Show-(7)]
	_ = x[ScreenUpdate-(8)]
}

var _ActionsNameToValueMap = map[string]Actions{
	`Close`:        0,
	`close`:        0,
	`Minimize`:     1,
	`minimize`:     1,
	`Resize`:       2,
	`resize`:       2,
	`Move`:         3,
	`move`:         3,
	`Focus`:        4,
	`focus`:        4,
	`DeFocus`:      5,
	`defocus`:      5,
	`Paint`:        6,
	`paint`:        6,
	`Show`:         7,
	`show`:         7,
	`ScreenUpdate`: 8,
	`screenupdate`: 8,
}

var _ActionsDescMap = map[Actions]string{
	0: `Close means that the window is about to close, but has not yet closed.`,
	1: `Minimize means that the window has been iconified / miniaturized / is no longer visible.`,
	2: `Resize means that the window was resized, including changes in DPI associated with moving to a new screen. Position may have also changed too. Requires a redraw.`,
	3: `Move means that the window was moved but NOT resized or changed in any other way -- does not require a redraw, but anything tracking positions will want to update.`,
	4: `Focus indicates that the window has been activated for receiving user input.`,
	5: `DeFocus indicates that the window is no longer activated for receiving input.`,
	6: `Paint indicates a request to repaint the window. This is sent right after the window is opened.`,
	7: `Show is for the WindowShow event -- sent by the system 500 msec after the window has opened, to trigger one-time actions such as configuring the main menu after the window has opened.`,
	8: `ScreenUpdate occurs when any of the screen information is updated This event is sent to the first window on the list of active windows and it should then perform any necessary updating`,
}

var _ActionsMap = map[Actions]string{
	0: `Close`,
	1: `Minimize`,
	2: `Resize`,
	3: `Move`,
	4: `Focus`,
	5: `DeFocus`,
	6: `Paint`,
	7: `Show`,
	8: `ScreenUpdate`,
}

// String returns the string representation
// of this Actions value.
func (i Actions) String() string {
	if str, ok := _ActionsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Actions value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Actions) SetString(s string) error {
	if val, ok := _ActionsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ActionsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Actions")
}

// Int64 returns the Actions value as an int64.
func (i Actions) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Actions value from an int64.
func (i *Actions) SetInt64(in int64) {
	*i = Actions(in)
}

// Desc returns the description of the Actions value.
func (i Actions) Desc() string {
	if str, ok := _ActionsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ActionsValues returns all possible values
// for the type Actions.
func ActionsValues() []Actions {
	return _ActionsValues
}

// Values returns all possible values
// for the type Actions.
func (i Actions) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ActionsValues))
	for i, d := range _ActionsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Actions.
func (i Actions) IsValid() bool {
	_, ok := _ActionsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Actions) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Actions) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}