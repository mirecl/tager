package tager

import (
	"fmt"
	"runtime/debug"
)

func TagMe() {
	fmt.Println(debug.ReadBuildInfo())
	buildifo, _ := debug.ReadBuildInfo()
	fmt.Println(buildifo.Deps)
	fmt.Println(buildifo.Deps[0].Version)
}
