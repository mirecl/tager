package tager

import (
	"fmt"
	"runtime/debug"
)

func TagMe() {
	buildifo, _ := debug.ReadBuildInfo()
	fmt.Println(buildifo.Deps[0].Version)
}
