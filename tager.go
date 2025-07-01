package tager

import (
	"fmt"
	"runtime/debug"
)

func TagMe() {
	fmt.Println("Hello")
	buildifo, _ := debug.ReadBuildInfo()
	fmt.Println(buildifo.Deps[0].Version)
}
