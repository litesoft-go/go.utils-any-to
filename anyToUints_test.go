package go_utils_any_to

import (
	"fmt"
	check "github.com/litesoft-go/go.utils-check"
	"math"
	"strconv"
	"testing"
)

func TestRangeLimitingToUint64(t *testing.T) {
	tests := []struct {
		min, it, max uint64
		wantErr      bool
	}{
		{5, 6, 5, true},
		{5, 5, 5, false},
		{0, 5, math.MaxUint64, false},
		{0, 0, math.MaxUint64, false},
		{0, math.MaxUint64, math.MaxUint64, false},
	}
	gotUint64, err := RangeLimitingToUint64(nil, "nil", 5, 5)
	if (err == nil) || (gotUint64 != 0) {
		t.Errorf("RangeLimitingToUint64(nil) unexpected result -- int64 = %v error: %v", gotUint64, err)
		return
	}
	for i, tt := range tests {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			got, err := RangeLimitingToUint64(tt.it, "("+name+")", tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("RangeLimitingToUint64() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				fmt.Println(check.TestIndent, "expected error:", name, "--", err)
			} else if got != tt.it {
				t.Errorf("RangeLimitingToUint64(%v) got = %v, want %v", tt.it, got, tt.it)
			}
		})
	}
}

func TestUintVariations(t *testing.T) {
	tests := []struct {
		value uint64
	}{
		{0},
		{1},
		{math.MaxUint8},
		{math.MaxUint16},
		{math.MaxUint32},
		{math.MaxUint64},
	}
	nilUint(t)
	nilU64(t)
	nilU32(t)
	nilU16(t)
	nilU8(t)

	for i, tt := range tests {
		name := fmt.Sprintf("(%v:%v)", i, tt.value)
		t.Run(name, func(t *testing.T) {
			if checkUint(t, tt.value) {
				if checkU64(t, tt.value) {
					if checkU32(t, tt.value) {
						if checkU16(t, tt.value) {
							_ = checkU8(t, tt.value)
						}
					}
				}
			}
		})
	}
}

func nilUint(t *testing.T) {
	i, err := Uint(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Uint' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilU64(t *testing.T) {
	i, err := Uint64(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Uint64' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilU32(t *testing.T) {
	i, err := Uint32(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Uint32' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilU16(t *testing.T) {
	i, err := Uint16(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Uint16' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func nilU8(t *testing.T) {
	i, err := Uint8(nil)
	if (i != 0) || (err == nil) {
		t.Errorf("expected 'Uint8' i==0 and err!=nil, but was i=%v and err=%v", i, err)
	}
}

func checkUint(t *testing.T, value uint64) bool {
	var result uint
	result, err := Uint(value)
	return validateUint(t, math.MaxUint, value, uint64(result), err)
}

func checkU64(t *testing.T, value uint64) bool {
	var result uint64
	result, err := Uint64(value)
	return validateUint(t, math.MaxUint64, value, result, err)
}

func checkU32(t *testing.T, value uint64) bool {
	var result uint32
	result, err := Uint32(value)
	return validateUint(t, math.MaxUint32, value, uint64(result), err)
}

func checkU16(t *testing.T, value uint64) bool {
	var result uint16
	result, err := Uint16(value)
	return validateUint(t, math.MaxUint16, value, uint64(result), err)
}

func checkU8(t *testing.T, value uint64) bool {
	var result uint8
	result, err := Uint8(value)
	return validateUint(t, math.MaxUint8, value, uint64(result), err)
}

func validateUint(t *testing.T, max, value, result uint64, err error) bool {
	if value <= max {
		if (value == result) && (err == nil) {
			return true
		}
		t.Errorf("%v <= %v, but value(%v) != (%v)result and/or err was: %v", value, max, value, result, err)
		return false
	}
	if (result == 0) && (err != nil) {
		return true
	}
	t.Errorf("not (%v <= %v), but result(%v) != 0 and/or err was: %v", value, max, result, err)
	return false
}
