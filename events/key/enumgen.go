// Code generated by "enumgen"; DO NOT EDIT.

package key

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

var _ModifiersValues = []Modifiers{0, 1, 2, 3}

// ModifiersN is the highest valid value
// for type Modifiers, plus one.
const ModifiersN Modifiers = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ModifiersNoOp() {
	var x [1]struct{}
	_ = x[Shift-(0)]
	_ = x[Control-(1)]
	_ = x[Alt-(2)]
	_ = x[Meta-(3)]
}

var _ModifiersNameToValueMap = map[string]Modifiers{
	`Shift`:   0,
	`shift`:   0,
	`Control`: 1,
	`control`: 1,
	`Alt`:     2,
	`alt`:     2,
	`Meta`:    3,
	`meta`:    3,
}

var _ModifiersDescMap = map[Modifiers]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
}

var _ModifiersMap = map[Modifiers]string{
	0: `Shift`,
	1: `Control`,
	2: `Alt`,
	3: `Meta`,
}

// String returns the string representation
// of this Modifiers value.
func (i Modifiers) String() string {
	str := ""
	for _, ie := range _ModifiersValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this Modifiers value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Modifiers) BitIndexString() string {
	if str, ok := _ModifiersMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Modifiers value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Modifiers) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the Modifiers value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *Modifiers) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _ModifiersNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _ModifiersNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type Modifiers")
		}
	}
	return nil
}

// Int64 returns the Modifiers value as an int64.
func (i Modifiers) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Modifiers value from an int64.
func (i *Modifiers) SetInt64(in int64) {
	*i = Modifiers(in)
}

// Desc returns the description of the Modifiers value.
func (i Modifiers) Desc() string {
	if str, ok := _ModifiersDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ModifiersValues returns all possible values
// for the type Modifiers.
func ModifiersValues() []Modifiers {
	return _ModifiersValues
}

// Values returns all possible values
// for the type Modifiers.
func (i Modifiers) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ModifiersValues))
	for i, d := range _ModifiersValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Modifiers.
func (i Modifiers) IsValid() bool {
	_, ok := _ModifiersMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i Modifiers) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *Modifiers) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Modifiers) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Modifiers) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
