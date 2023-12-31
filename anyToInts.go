package go_utils_any_to

import (
	check "github.com/litesoft-go/go.utils-check"
	"math"
	"math/bits"
	"reflect"
)

func RangeLimitingToInt64(it any, what string, min, max int64) (int64, error) {
	reflectValue, err := ValueOf(it, what, reflect.Value.CanInt)
	if err != nil {
		return 0, err
	}
	return check.LimitIntRange(reflectValue.Int(), // int64 by contract!
		what, min, max)
}

func Int64(it any) (int64, error) {
	return RangeLimitingToInt64(it, "int64", math.MinInt64, math.MaxInt64)
}

func Int32(it any) (int32, error) {
	value, err := RangeLimitingToInt64(it, "int32", math.MinInt32, math.MaxInt32)
	return int32(value), err // on err != nil, value is always 0!
}

func Int16(it any) (int16, error) {
	value, err := RangeLimitingToInt64(it, "int16", math.MinInt16, math.MaxInt16)
	return int16(value), err // on err != nil, value is always 0!
}

func Int8(it any) (int8, error) {
	value, err := RangeLimitingToInt64(it, "int8", math.MinInt8, math.MaxInt8)
	return int8(value), err // on err != nil, value is always 0!
}

func Int(it any) (int, error) {
	var min, max int64
	if bits.UintSize == 64 {
		min, max = math.MinInt64, math.MaxInt64
	} else {
		min, max = math.MinInt32, math.MaxInt32
	}
	value, err := RangeLimitingToInt64(it, "int", min, max)
	return int(value), err // on err != nil, value is always 0!
}
