package main

import (
	"reflect"
	"testing"
)

func TestMakeGear(t *testing.T) {
	got, err := MakeGear(0.15, 15.0)
	if err != nil {
		t.Errorf("error making gear with MakeGear: %v", err)
	}

	want := Gear{
		Module:      0.15,
		RefDiameter: 15,
		ToothCount:  0,
	}
	Assert(t, want, got)
}

func BenchmarkMakeGear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeGear(0.25, 255.0)
	}
}

func TestSetArgsToGears(t *testing.T) {
	args1 := []string{"0.3", "155", "150", "125", "90", "55", "15"}
	args2 := []string{"0.1", "125", "90", "55", "65"}
	args3 := []string{"1.5", "109", "155"}
	argsTestSet := [][]string{args1, args2, args3}

	want := []Gears{
		{
			Gear{
				Module:      0.3,
				RefDiameter: 155,
				ToothCount:  0,
			},
			Gear{
				Module:      0.3,
				RefDiameter: 150,
				ToothCount:  0,
			},
			Gear{
				Module:      0.3,
				RefDiameter: 125,
				ToothCount:  0,
			},
			Gear{
				Module:      0.3,
				RefDiameter: 90,
				ToothCount:  0,
			},
			Gear{
				Module:      0.3,
				RefDiameter: 55,
				ToothCount:  0,
			},
			Gear{
				Module:      0.3,
				RefDiameter: 15,
				ToothCount:  0,
			},
		},
		{
			Gear{
				Module:      0.1,
				RefDiameter: 125,
				ToothCount:  0,
			},
			Gear{
				Module:      0.1,
				RefDiameter: 90,
				ToothCount:  0,
			},
			Gear{
				Module:      0.1,
				RefDiameter: 55,
				ToothCount:  0,
			},
			Gear{
				Module:      0.1,
				RefDiameter: 65,
				ToothCount:  0,
			},
		},
		{
			Gear{
				Module:      1.5,
				RefDiameter: 109,
				ToothCount:  0,
			},
			Gear{
				Module:      1.5,
				RefDiameter: 155,
				ToothCount:  0,
			},
		},
	}

	for i, args := range argsTestSet {
		if len(args) < 2 {
			t.Errorf("need at least two arguments passed, you passed: %v", args)
		}
		got, err := SetArgsToGears(args)
		if err != nil {
			t.Errorf("error during SetArgsToGears(args), args: %v", err)
		}
		if !reflect.DeepEqual(want[i], got) {
			t.Errorf("wanted %v want but got %v", want[i], got)
		}
	}

}

func BenchmarkSetArgsToGears(b *testing.B) {
	args1 := []string{"0.3", "155", "150", "125", "90", "55", "15"}
	args2 := []string{"0.1", "125", "90", "55", "65"}
	args3 := []string{"1.5", "109", "155"}
	argsTestSet := [][]string{args1, args2, args3}

	for i := 0; i < b.N; i++ {
		for _, args := range argsTestSet {
			SetArgsToGears(args)
		}
	}

}

func TestToFloat(t *testing.T) {
	var got []float64
	var members []interface{} = []interface{}{"666", 84567, 3.14, nil}
	for _, m := range members {
		fm, err := ToFloat(m)
		if err != nil {
			t.Errorf("Error parsing member value to float %v", err)
		}
		got = append(got, fm)
	}

	want := []float64{666.0, 84567.0, 3.14, 0.0}
	for i, w := range want {
		if w != got[i] {
			t.Errorf("Wanted %v but got %v", want, got)
		}
	}
}

func BenchmarkToFloat(b *testing.B) {
	members := []interface{}{10, 145, "123", "513.2", 3.1415, nil}
	for i := 0; i < b.N; i++ {
		for _, e := range members {
			ToFloat(e)
		}
	}
}
