// WARNING: This file has automatically been generated on Thu, 15 Feb 2018 23:59:40 EET.
// By https://git.io/c-for-go. DO NOT EDIT.

package imgui

/*
#include "cimgui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// Foo function as declared in imgui/cimgui.h:20
func Foo(out *ImVec2) {
	cout, _ := out.PassRef()
	C.foo(cout)
}
