package go_utils_any_to

import (
	"fmt"
	check "github.com/litesoft-go/go.utils-check"
	"reflect"
	"testing"
)

type TestInterface interface {
	Active() bool
}

func TestValue(t *testing.T) {
	tests := []struct {
		name    string
		arg     any
		wantErr bool
	}{
		{"nil", nil, true},
		{"empty", "", false},
		{"int", int(67), false},
		{"funcType", reflect.Value.CanUint, false},
		{"interface.func", TestInterface.Active, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reflectValue, err := Value(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				fmt.Println(check.TestIndent, "expected error:", tt.name, "--", err)
			} else {
				fmt.Printf("%v reflectValue: '%v'\n", check.TestIndent, reflectValue)
				//if !reflect.DeepEqual(reflectValue, tt.wantReflectValue)
				//t.Errorf("Value() reflectValue = %v, want %v", reflectValue, tt.wantReflectValue)
			}
		})
	}
}

func TestValueOf(t *testing.T) {
	tests := []struct {
		name      string
		it        any
		predicate ReflectValuePredicate
		wantErr   bool
	}{
		{"nil", nil, reflect.Value.IsValid, true},
		{"empty", "", reflect.Value.IsZero, false},
		{"int(76)", int(76), reflect.Value.CanUint, true},
		{"int(67)", int(67), reflect.Value.CanInt, false},
		{"uint(67)", uint(67), reflect.Value.CanUint, false},
		{"funcType", reflect.Value.CanUint, reflect.Value.IsValid, false},
		{"interface.func", TestInterface.Active, reflect.Value.IsValid, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reflectValue, err := ValueOf(tt.it, tt.name, tt.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValueOf() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				fmt.Println(check.TestIndent, "expected error:", tt.name, "--", err)
			} else {
				fmt.Printf("%v reflectValue: '%v'\n", check.TestIndent, reflectValue)
				//if !reflect.DeepEqual(reflectValue, tt.wantReflectValue)
				//t.Errorf("Value() reflectValue = %v, want %v", reflectValue, tt.wantReflectValue)
			}
		})
	}
}
