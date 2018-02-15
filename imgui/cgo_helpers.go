// WARNING: This file has automatically been generated on Thu, 15 Feb 2018 23:59:40 EET.
// By https://git.io/c-for-go. DO NOT EDIT.

package imgui

/*
#include "cimgui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocStructImVec2Memory allocates memory for type C.struct_ImVec2 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImVec2Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImVec2Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImVec2Value = unsafe.Sizeof([1]C.struct_ImVec2{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ImVec2) Ref() *C.struct_ImVec2 {
	if x == nil {
		return nil
	}
	return x.ref74e98a33
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ImVec2) Free() {
	if x != nil && x.allocs74e98a33 != nil {
		x.allocs74e98a33.(*cgoAllocMap).Free()
		x.ref74e98a33 = nil
	}
}

// NewImVec2Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewImVec2Ref(ref unsafe.Pointer) *ImVec2 {
	if ref == nil {
		return nil
	}
	obj := new(ImVec2)
	obj.ref74e98a33 = (*C.struct_ImVec2)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ImVec2) PassRef() (*C.struct_ImVec2, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref74e98a33 != nil {
		return x.ref74e98a33, nil
	}
	mem74e98a33 := allocStructImVec2Memory(1)
	ref74e98a33 := (*C.struct_ImVec2)(mem74e98a33)
	allocs74e98a33 := new(cgoAllocMap)
	allocs74e98a33.Add(mem74e98a33)

	var cx_allocs *cgoAllocMap
	ref74e98a33.x, cx_allocs = (C.float)(x.X), cgoAllocsUnknown
	allocs74e98a33.Borrow(cx_allocs)

	var cy_allocs *cgoAllocMap
	ref74e98a33.y, cy_allocs = (C.float)(x.Y), cgoAllocsUnknown
	allocs74e98a33.Borrow(cy_allocs)

	x.ref74e98a33 = ref74e98a33
	x.allocs74e98a33 = allocs74e98a33
	return ref74e98a33, allocs74e98a33

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ImVec2) PassValue() (C.struct_ImVec2, *cgoAllocMap) {
	if x.ref74e98a33 != nil {
		return *x.ref74e98a33, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ImVec2) Deref() {
	if x.ref74e98a33 == nil {
		return
	}
	x.X = (float32)(x.ref74e98a33.x)
	x.Y = (float32)(x.ref74e98a33.y)
}
