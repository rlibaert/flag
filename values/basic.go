package values

import (
	"flag"
	"fmt"
)

type basic = interface {
	bool |
		complex64 | complex128 |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		string | []byte
}

func parseBasic[T basic](s string) (T, error) {
	var v T
	var err error

	switch p := any(&v).(type) {
	case *string:
		*p = s

	case *[]byte:
		*p = []byte(s)

	default:
		_, err = fmt.Sscan(s, p)
	}

	return v, err
}

func formatBasic[T basic](v T) string {
	return fmt.Sprint(v)
}

// Basic declares a [flag.Value] for Go [basic] types.
// The actual value type is T.
func Basic[T basic]() flag.Value {
	return Generic[T](parseBasic, formatBasic)
}

// BasicVar is like [Basic] but stores the value in p.
func BasicVar[T basic](p *T) flag.Value {
	return GenericVar(p, parseBasic, formatBasic)
}

// BasicList declares a list-style [flag.Value] for Go [basic] types.
// The actual value type is []T.
func BasicList[T basic]() flag.Value {
	return GenericList[T](parseBasic, formatBasic)
}

// BasicListVar is like [BasicList] but stores the values in p.
func BasicListVar[T basic](p *[]T) flag.Value {
	return GenericListVar(p, parseBasic, formatBasic)
}

// BasicSlice declares a slice-style [flag.Value] for Go [basic] types.
// The input strings are split around sep before parsing.
// The actual value type is []T.
func BasicSlice[T basic](sep string) flag.Value {
	return GenericSlice[T](sep, parseBasic, formatBasic)
}

// BasicSliceVar is like [BasicSlice] but stores the values in p.
func BasicSliceVar[T basic](p *[]T, sep string) flag.Value {
	return GenericSliceVar(p, sep, parseBasic, formatBasic)
}
