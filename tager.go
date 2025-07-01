package tager

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
)

func TagMe() {
	_, file, _, _ := runtime.Caller(0)
	fmt.Println(file)
	// Получаем имя текущего пакета через рефлексию
	pkgPath := reflect.TypeOf(TagMe).PkgPath()
	fmt.Println(pkgPath)

	buildifo, _ := debug.ReadBuildInfo()
	fmt.Println(buildifo.Deps[0].Version)
}
