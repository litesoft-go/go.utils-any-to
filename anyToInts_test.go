package go_utils_any_to

import (
	"fmt"
	check "github.com/litesoft-go/go.utils-check"
	"math"
	"strconv"
	"testing"
)

func TestRangeLimitingToInt64(t *testing.T) {
	tests := []struct {
		min, it, max int64
		wantErr      bool
	}{
		{5, 6, 5, true},
		{5, 5, 5, false},
		{math.MinInt64, 5, math.MaxInt64, false},
		{math.MinInt64, math.MinInt64, math.MaxInt64, false},
		{math.MinInt64, math.MaxInt64, math.MaxInt64, false},
	}
	gotInt64, err := RangeLimitingToInt64(nil, "nil", 5, 5)
	if (err == nil) || (gotInt64 != 0) {
		t.Errorf("RangeLimitingToInt64(nil) unexpected result -- int64 = %v error: %v", gotInt64, err)
		return
	}
	for i, tt := range tests {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			got, err := RangeLimitingToInt64(tt.it, "("+name+")", tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("RangeLimitingToInt64() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				fmt.Println(check.TestIndent, "expected error:", name, "--", err)
			} else if got != tt.it {
				t.Errorf("RangeLimitingToInt64(%v) got = %v, want %v", tt.it, got, tt.it)
			}
		})
	}
}

func TestIntVariations(t *testing.T) {
	tests := []struct {
		value int64
	}{
		{math.MinInt64},
		{math.MinInt32},
		{math.MinInt16},
		{math.MinInt8},
		{-1},
		{0},
		{1},
		{math.MaxInt8},
		{math.MaxInt16},
		{math.MaxInt32},
		{math.MaxInt64},
	}
	nilInt(t)
	nilI64(t)
	nilI32(t)
	nilI16(t)
	nilI8(t)

	for i, tt := range tests {
		name := fmt.Sprintf("(%v:%v)", i, tt.value)
		t.Run(name, func(t *testing.T) {
			if checkInt(t, tt.value) {
				if checkI64(t, tt.value) {
					if checkI32(t, tt.value) {
						if checkI16(t, tt.value) {
							_ = checkI8(t, tt.value)
						}
					}
				}
			}
		})
	}
}

func nilInt(t *testing.T) {
	i, err := Int(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Int' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilI64(t *testing.T) {
	i, err := Int64(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Int64' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilI32(t *testing.T) {
	i, err := Int32(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Int32' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilI16(t *testing.T) {
	i, err := Int16(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Int16' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilI8(t *testing.T) {
	i, err := Int8(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Int8' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func checkInt(t *testing.T, value int64) bool {
	var result int
	result, err := Int(value)
	return validateInt(t, math.MinInt, math.MaxInt, value, int64(result), err)
}

func checkI64(t *testing.T, value int64) bool {
	var result int64
	result, err := Int64(value)
	return validateInt(t, math.MinInt64, math.MaxInt64, value, result, err)
}

func checkI32(t *testing.T, value int64) bool {
	var result int32
	result, err := Int32(value)
	return validateInt(t, math.MinInt32, math.MaxInt32, value, int64(result), err)
}

func checkI16(t *testing.T, value int64) bool {
	var result int16
	result, err := Int16(value)
	return validateInt(t, math.MinInt16, math.MaxInt16, value, int64(result), err)
}

func checkI8(t *testing.T, value int64) bool {
	var result int8
	result, err := Int8(value)
	return validateInt(t, math.MinInt8, math.MaxInt8, value, int64(result), err)
}

func validateInt(t *testing.T, min, max, value, result int64, err error) bool {
	if (min <= value) && (value <= max) {
		if (value == result) && (err == nil) {
			return true
		}
		t.Errorf("%v <= %v <= %v, but value(%v) != (%v)result and/or err was: %v", min, value, max, value, result, err)
		return false
	}
	if (result == 0) && (err != nil) {
		return true
	}
	t.Errorf("not (%v <= %v <= %v), but result(%v) != 0 and/or err was: %v", min, value, max, result, err)
	return false
}
