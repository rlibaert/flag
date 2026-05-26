package values

import (
	"flag"
	"fmt"
)

// Stringer declares a [flag.Value] implemented using the parse function and String method of a [fmt.Stringer].
// The actual value type is T.
func Stringer[T fmt.Stringer](parse func(string) (T, error)) flag.Value {
	return Generic(parse, T.String)
}

// StringerVar is like [Stringer] but store the value in p.
func StringerVar[T fmt.Stringer](p *T, parse func(string) (T, error)) flag.Value {
	return GenericVar(p, parse, T.String)
}

// StringerList declares a list-style [flag.Value] implemented using the parse function and String method of a [fmt.Stringer].
// The actual value type is []T.
func StringerList[T fmt.Stringer](parse func(string) (T, error)) flag.Value {
	return GenericList(parse, T.String)
}

// StringerListVar is like [StringerList] but stores the values in p.
func StringerListVar[T fmt.Stringer](p *[]T, parse func(string) (T, error)) flag.Value {
	return GenericListVar(p, parse, T.String)
}

// StringerSlice declares a slice-style [flag.Value] implemented using the parse function and String method of a [fmt.Stringer].
// The input strings are split around sep before parsing.
// The actual value type is []T.
func StringerSlice[T fmt.Stringer](sep string, parse func(string) (T, error)) flag.Value {
	return GenericSlice(sep, parse, T.String)
}

// StringerSliceVar is like [StringerSlice] but stores the values in p.
func StringerSliceVar[T fmt.Stringer](p *[]T, sep string, parse func(string) (T, error)) flag.Value {
	return GenericSliceVar(p, sep, parse, T.String)
}
