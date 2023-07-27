package go_utils_any_to

import (
	"errors"
	"fmt"
	"reflect"
)

// Value -- return the "reflect.Value" and an "error" -- if the
// error is not nil, then the "reflect.Value" state should not be assumed!
// -
func Value(it any) (reflectValue reflect.Value, err error) {
	if it == nil {
		err = errors.New("was nil")
	} else {
		reflectValue = reflect.ValueOf(it)
		if !reflectValue.IsValid() {
			err = errors.New("was IsValid")
		}
	}
	return
}

// ReflectValuePredicate -- is a function "specification" to be used in "ValueOf" below
// to constrain the non-error response to those "reflect.Value"(s) who when
// ReflectValuePredicate is called return true.  The methods in "reflect.Value" that can
// be used as a ReflectValuePredicate are:
// * Bool
// * CanAddr
// * CanSet
// * CanComplex
// * CanFloat
// * CanInt
// * CanInterface
// * CanUint
// * Comparable
// * IsNil
// * IsValid
// * IsZero
// -
type ReflectValuePredicate func(reflect.Value) bool

// ValueOf -- enhances "Value" method above specifically check for appropriate
// condition (normally "type") from the methods listed above (ReflectValuePredicate)
// AND add context "what" to the errors.
// -
func ValueOf(it any, what string, predicate ReflectValuePredicate) (reflectValue reflect.Value, err error) {
	reflectValue, err = Value(it)
	if err != nil {
		err = fmt.Errorf("'%v' expected, but Value('%v') %w", what, it, err)
	}
	can := predicate(reflectValue)
	if !can {
		err = fmt.Errorf("'%v' expected, but was: %v", what, reflectValue.Kind())
	}
	return
}
