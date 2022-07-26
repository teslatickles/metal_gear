package main

import (
	"os"
)

func main() {
	gl, err := SetArgsToGears(os.Args[1:])
	if err != nil {
		panic(err)
	}

	gr, err := gl.GetMultiToothCount()
	if err != nil {
		panic(err)
	}
	PrintGears(gr)
}
