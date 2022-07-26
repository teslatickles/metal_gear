package main

import (
	"errors"
	"fmt"
	"strconv"
)

func MakeGear(mod float64, rd float64) (Gear, error) {
	if mod == 0 {
		return Gear{}, errors.New(Red + "module cannot be 0")
	}
	if rd == 0 {
		return Gear{}, errors.New(Red + "reference diameter cannot be 0")
	}
	return Gear{
		Module:      mod,
		RefDiameter: rd,
		ToothCount:  0, //init
	}, nil
}

func SetArgsToGears(args []string) (Gears, error) {
	var mod float64
	var gl Gears

	if len(args) < 2 {
		return nil, errors.New(Red + "must pass at least two args: [0]: Module, [1n]: Reference Diameters (space-delimited)")
	}
	for i, a := range args {
		if i == 0 {
			m, err := ToFloat(a)
			if err != nil {
				return nil, err
			}
			mod = m
		} else {
			rd, err := ToFloat(a)
			if err != nil {
				return nil, err
			}
			g, err := MakeGear(mod, rd)
			if err != nil {
				return nil, err
			}
			gl = append(gl, g)
		}
	}
	if len(gl) == 0 {
		return nil, errors.New(Red + "gear List (gl) is empty")
	}
	return gl, nil
}

func PrintGears(gl Gears) {
	for _, g := range gl {
		fmt.Printf(Green+"⚙️  %+v\n", g)
	}
}

func ToFloat(n interface{}) (float64, error) {
	switch v := n.(type) {
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0.0, err
		}
		return f, nil
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0.0, nil
	}
}
