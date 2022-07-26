package main

import (
	"errors"
	"math"
)

func ValidateGear(g *Gear) (int, error) {
	if g.RefDiameter == 0 {
		err := errors.New(Red + "reference diameter cannot be 0")
		return int(g.RefDiameter), err
	}
	if g.Module == 0 {
		err := errors.New(Red + "module cannot be 0")
		return int(g.Module), err
	}
	if math.Signbit(g.RefDiameter) {
		err := errors.New(Red + "reference diameter cannot be negative")
		return int(g.RefDiameter), err
	}
	if math.Signbit(g.Module) {
		err := errors.New(Red + "module cannot be negative")
		return int(g.Module), err
	}
	return 0, nil
}

func (g *Gear) CountTeeth() (int, error) {
	_, err := ValidateGear(g)
	if err != nil {
		return 0, err
	}

	tc := g.RefDiameter / g.Module
	g.ToothCount = int(tc)
	return g.ToothCount, nil
}

func (g *Gear) GetRefDiameter() (float64, error) {
	_, err := ValidateGear(g)
	if err != nil {
		return 0, err
	}

	t, err := g.CountTeeth()
	if err != nil {
		return 0, err
	}

	rd := g.Module * float64(t)
	return float64(rd), nil
}

func (g *Gear) GetRefPitch() (float64, error) {
	_, err := ValidateGear(g)
	if err == nil {
		return g.Module * Pi, nil
	}
	return 0.0, err
}

func (gl Gears) GetMultiToothCount() (Gears, error) {
	for i, g := range gl {
		_, err := ValidateGear(&g)
		if err != nil {
			return nil, err
		}
		t, err := g.CountTeeth()
		if err != nil {
			return nil, err
		}
		g.ToothCount = t
		gl[i].ToothCount = g.ToothCount
	}
	return gl, nil
}
