package main

import (
	"reflect"
	"testing"
)

func Assert(t *testing.T, want interface{}, got interface{}) {
	switch v := want.(type) {
	case Gear:
		if !reflect.DeepEqual(v, got) {
			t.Errorf("expected %v but got %v", v, got)
		}
	default:
		if want != got {
			t.Errorf("expected %v but got %v", v, got)
		}
	}
}
