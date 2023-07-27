package go_utils_any_to

import (
	check "github.com/litesoft-go/go.utils-check"
	"math"
	"math/bits"
	"reflect"
)

func RangeLimitingToUint64(it any, what string, min, max uint64) (uint64, error) {
	reflectValue, err := ValueOf(it, what, reflect.Value.CanUint)
	if err != nil {
		return 0, err
	}
	return check.LimitUintRange(reflectValue.Uint(), // uint64 by contract!
		what, min, max)
}

func Uint64(it any) (uint64, error) {
	return RangeLimitingToUint64(it, "uint64", 0, math.MaxUint64)
}

func Uint32(it any) (uint32, error) {
	value, err := RangeLimitingToUint64(it, "uint32", 0, math.MaxUint32)
	return uint32(value), err // on err != nil, value is always 0!
}

func Uint16(it any) (uint16, error) {
	value, err := RangeLimitingToUint64(it, "uint16", 0, math.MaxUint16)
	return uint16(value), err // on err != nil, value is always 0!
}

func Uint8(it any) (uint8, error) {
	value, err := RangeLimitingToUint64(it, "uint8", 0, math.MaxUint8)
	return uint8(value), err // on err != nil, value is always 0!
}

func Uint(it any) (uint, error) {
	var max uint64
	if bits.UintSize == 64 {
		max = math.MaxUint64
	} else {
		max = math.MaxUint32
	}
	value, err := RangeLimitingToUint64(it, "uint", 0, max)
	return uint(value), err // on err != nil, value is always 0!
}
