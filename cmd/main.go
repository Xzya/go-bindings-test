package main

import (
	"fmt"
	"github.com/Xzya/go-bindings-test/imgui"
)

func main() {
	v := new(imgui.ImVec2)
	imgui.Foo(v)
	fmt.Printf("%#v\n", v)
}
