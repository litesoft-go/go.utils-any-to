package go_utils_any_to

import (
	"errors"
	"reflect"
)

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
