// Code generated by "enumgen"; DO NOT EDIT.

package mouse

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _ButtonsValues = []Buttons{0, 1, 2, 3}

// ButtonsN is the highest valid value
// for type Buttons, plus one.
const ButtonsN Buttons = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ButtonsNoOp() {
	var x [1]struct{}
	_ = x[NoButton-(0)]
	_ = x[Left-(1)]
	_ = x[Middle-(2)]
	_ = x[Right-(3)]
}

var _ButtonsNameToValueMap = map[string]Buttons{
	`NoButton`: 0,
	`nobutton`: 0,
	`Left`:     1,
	`left`:     1,
	`Middle`:   2,
	`middle`:   2,
	`Right`:    3,
	`right`:    3,
}

var _ButtonsDescMap = map[Buttons]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
}

var _ButtonsMap = map[Buttons]string{
	0: `NoButton`,
	1: `Left`,
	2: `Middle`,
	3: `Right`,
}

// String returns the string representation
// of this Buttons value.
func (i Buttons) String() string {
	if str, ok := _ButtonsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Buttons value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Buttons) SetString(s string) error {
	if val, ok := _ButtonsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ButtonsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Buttons")
}

// Int64 returns the Buttons value as an int64.
func (i Buttons) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Buttons value from an int64.
func (i *Buttons) SetInt64(in int64) {
	*i = Buttons(in)
}

// Desc returns the description of the Buttons value.
func (i Buttons) Desc() string {
	if str, ok := _ButtonsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ButtonsValues returns all possible values
// for the type Buttons.
func ButtonsValues() []Buttons {
	return _ButtonsValues
}

// Values returns all possible values
// for the type Buttons.
func (i Buttons) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ButtonsValues))
	for i, d := range _ButtonsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Buttons.
func (i Buttons) IsValid() bool {
	_, ok := _ButtonsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Buttons) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Buttons) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _ActionsValues = []Actions{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// ActionsN is the highest valid value
// for type Actions, plus one.
const ActionsN Actions = 10

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ActionsNoOp() {
	var x [1]struct{}
	_ = x[NoAction-(0)]
	_ = x[Press-(1)]
	_ = x[Release-(2)]
	_ = x[DoubleClick-(3)]
	_ = x[Move-(4)]
	_ = x[Drag-(5)]
	_ = x[Scroll-(6)]
	_ = x[Enter-(7)]
	_ = x[Exit-(8)]
	_ = x[Hover-(9)]
}

var _ActionsNameToValueMap = map[string]Actions{
	`NoAction`:    0,
	`noaction`:    0,
	`Press`:       1,
	`press`:       1,
	`Release`:     2,
	`release`:     2,
	`DoubleClick`: 3,
	`doubleclick`: 3,
	`Move`:        4,
	`move`:        4,
	`Drag`:        5,
	`drag`:        5,
	`Scroll`:      6,
	`scroll`:      6,
	`Enter`:       7,
	`enter`:       7,
	`Exit`:        8,
	`exit`:        8,
	`Hover`:       9,
	`hover`:       9,
}

var _ActionsDescMap = map[Actions]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
	5: ``,
	6: ``,
	7: ``,
	8: ``,
	9: ``,
}

var _ActionsMap = map[Actions]string{
	0: `NoAction`,
	1: `Press`,
	2: `Release`,
	3: `DoubleClick`,
	4: `Move`,
	5: `Drag`,
	6: `Scroll`,
	7: `Enter`,
	8: `Exit`,
	9: `Hover`,
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

var _SelectModesValues = []SelectModes{0, 1, 2, 3, 4, 5, 6}

// SelectModesN is the highest valid value
// for type SelectModes, plus one.
const SelectModesN SelectModes = 7

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _SelectModesNoOp() {
	var x [1]struct{}
	_ = x[SelectOne-(0)]
	_ = x[ExtendContinuous-(1)]
	_ = x[ExtendOne-(2)]
	_ = x[NoSelect-(3)]
	_ = x[Unselect-(4)]
	_ = x[SelectQuiet-(5)]
	_ = x[UnselectQuiet-(6)]
}

var _SelectModesNameToValueMap = map[string]SelectModes{
	`SelectOne`:        0,
	`selectone`:        0,
	`ExtendContinuous`: 1,
	`extendcontinuous`: 1,
	`ExtendOne`:        2,
	`extendone`:        2,
	`NoSelect`:         3,
	`noselect`:         3,
	`Unselect`:         4,
	`unselect`:         4,
	`SelectQuiet`:      5,
	`selectquiet`:      5,
	`UnselectQuiet`:    6,
	`unselectquiet`:    6,
}

var _SelectModesDescMap = map[SelectModes]string{
	0: `SelectOne selects a single item, and is the default when no modifier key is pressed`,
	1: `ExtendContinuous, activated by Shift key, extends the selection to select a continuous region of selected items, with no gaps`,
	2: `ExtendOne, activated by Control or Meta / Command, extends the selection by adding the one additional item just clicked on, creating a potentially discontinuous set of selected items`,
	3: `NoSelect means do not update selection -- this is used programmatically and not available via modifier key`,
	4: `Unselect means unselect items -- this is used programmatically and not available via modifier key -- typically ExtendOne will unselect if already selected`,
	5: `SelectQuiet means select without doing other updates or signals -- for bulk updates with a final update at the end -- used programmatically`,
	6: `UnselectQuiet means unselect without doing other updates or signals -- for bulk updates with a final update at the end -- used programmatically`,
}

var _SelectModesMap = map[SelectModes]string{
	0: `SelectOne`,
	1: `ExtendContinuous`,
	2: `ExtendOne`,
	3: `NoSelect`,
	4: `Unselect`,
	5: `SelectQuiet`,
	6: `UnselectQuiet`,
}

// String returns the string representation
// of this SelectModes value.
func (i SelectModes) String() string {
	if str, ok := _SelectModesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the SelectModes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *SelectModes) SetString(s string) error {
	if val, ok := _SelectModesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _SelectModesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type SelectModes")
}

// Int64 returns the SelectModes value as an int64.
func (i SelectModes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the SelectModes value from an int64.
func (i *SelectModes) SetInt64(in int64) {
	*i = SelectModes(in)
}

// Desc returns the description of the SelectModes value.
func (i SelectModes) Desc() string {
	if str, ok := _SelectModesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// SelectModesValues returns all possible values
// for the type SelectModes.
func SelectModesValues() []SelectModes {
	return _SelectModesValues
}

// Values returns all possible values
// for the type SelectModes.
func (i SelectModes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_SelectModesValues))
	for i, d := range _SelectModesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type SelectModes.
func (i SelectModes) IsValid() bool {
	_, ok := _SelectModesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SelectModes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SelectModes) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}