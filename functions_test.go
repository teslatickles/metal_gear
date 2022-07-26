package main

import (
	"math"
	"testing"
)

func TestValidateGear(t *testing.T) {
	// bad gear - missing diameter
	g1 := &Gear{
		Module:      0.1,
		RefDiameter: 0,
		ToothCount:  0,
	}

	// bad gear - missing module
	g2 := &Gear{
		Module:      0,
		RefDiameter: 10,
		ToothCount:  0,
	}

	// bad gear - negative module
	g3 := &Gear{
		Module:      -1,
		RefDiameter: 10,
		ToothCount:  0,
	}

	// bad gear - negative diameter
	g4 := &Gear{
		Module:      1,
		RefDiameter: -10,
		ToothCount:  0,
	}

	gl := Gears{*g1, *g2, *g3, *g4}
	for _, g := range gl {
		_, err := ValidateGear(&g)
		if err == nil {
			t.Error("err should not be nil")
		}
	}
}

func BenchmarkValidateGear(b *testing.B) {
	g1 := &Gear{
		Module:      0.3,
		RefDiameter: 13,
		ToothCount:  0,
	}
	g2 := &Gear{
		Module:      0.3,
		RefDiameter: 13,
		ToothCount:  0,
	}
	g3 := &Gear{
		Module:      0.3,
		RefDiameter: 13,
		ToothCount:  0,
	}
	g4 := &Gear{
		Module:      0.3,
		RefDiameter: 13,
		ToothCount:  0,
	}
	gl := Gears{*g1, *g2, *g3, *g4}
	for i := 0; i < b.N; i++ {
		for _, g := range gl {
			_, err := ValidateGear(&g)
			if err != nil {
				b.Errorf("error validating gear %v", err)
			}
		}

	}
}

func TestCountTeeth(t *testing.T) {
	g := Gear{
		Module:      0.1,
		RefDiameter: 10,
		ToothCount:  0,
	}
	f := g.RefDiameter / g.Module
	expected := int(f)
	got, err := g.CountTeeth()
	if err != nil {
		t.Errorf("error %v occurred while attempting CountTeeth on Gear: %v", err, g)
	}
	Assert(t, float64(expected), float64(got))
}

func BenchmarkCountTeeth(b *testing.B) {
	g := Gear{
		Module:      0.3,
		RefDiameter: 13,
		ToothCount:  0,
	}
	for i := 0; i < b.N; i++ {
		g.CountTeeth()
	}
}

func TestGetRefPitch(t *testing.T) {
	expected := math.Pi * 0.3
	g := Gear{
		Module:      0.3,
		RefDiameter: 15,
		ToothCount:  0,
	}
	got, err := g.GetRefPitch()
	if err != nil {
		t.Errorf("error %v occurred while attempting GetRefPitch on Gear: %v", err, g)
	}
	Assert(t, expected, got)
}

func BenchmarkRefPitch(b *testing.B) {
	g := Gear{
		Module:      0.3,
		RefDiameter: 13,
		ToothCount:  0,
	}
	for i := 0; i < b.N; i++ {
		g.GetRefPitch()
	}
}

func TestGetRefDiameter(t *testing.T) {
	g := Gear{
		Module:      0.1,
		RefDiameter: 155,
		ToothCount:  0,
	}
	te, err := g.CountTeeth()
	if err != nil {
		t.Errorf("error %v occurred while attempting CountTeeth on Gear: %v", err, g)
	}
	expected := g.Module * float64(te)

	got, err := g.GetRefDiameter()
	if err != nil {
		panic(err)
	}
	Assert(t, expected, got)
}

func BenchmarkRefDiameter(b *testing.B) {
	g := Gear{
		Module:      0.3,
		RefDiameter: 13,
		ToothCount:  0,
	}
	for i := 0; i < b.N; i++ {
		g.GetRefDiameter()
	}
}

func TestGetMultiToothCount(t *testing.T) {
	g1 := Gear{
		Module:      0.5,
		RefDiameter: 10,
		ToothCount:  0,
	}
	g2 := Gear{
		Module:      0.1,
		RefDiameter: 20,
		ToothCount:  0,
	}
	gl := Gears{g1, g2}
	gtc, err := gl.GetMultiToothCount()
	if err != nil {
		t.Errorf("Error getting tooth count on multiple gears: %v", err)
	}

	Assert(t, gl[0].RefDiameter/g1.Module, float64(gtc[0].ToothCount))
	Assert(t, gl[1].RefDiameter/g2.Module, float64(gtc[1].ToothCount))
}

func BenchmarkGetMultiToothCount(b *testing.B) {
	g1 := Gear{
		Module:      0.5,
		RefDiameter: 14,
		ToothCount:  0,
	}
	g2 := Gear{
		Module:      0.1,
		RefDiameter: 45,
		ToothCount:  0,
	}
	g3 := Gear{
		Module:      1.5,
		RefDiameter: 60,
		ToothCount:  0,
	}
	gl := Gears{g1, g2, g3}
	for i := 0; i < b.N; i++ {
		gl.GetMultiToothCount()
	}
}
