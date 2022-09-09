//go:build !tinygo.wasm && !wasi

package internal

import "reflect"

// SliceHeader returns a pointer to a reflect.SliceHeader when not TinyGo.
func SliceHeader(ptr uintptr, size uint32) *reflect.SliceHeader {
	return &reflect.SliceHeader{Data: ptr, Len: int(size), Cap: int(size)}
}
