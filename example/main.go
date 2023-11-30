package main

import (
	"os"

	"github.com/robertlestak/genericflag"
)

func main() {
	gf := genericflag.NewFlagSet("build")
	if err := gf.Parse(os.Args[1:]); err != nil {
		panic(err)
	}
	// print all the flags
	for k, v := range gf.Flags {
		for _, val := range v {
			println(k, val)
		}
	}
}
