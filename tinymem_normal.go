//go:build !tinygo.wasm && !wasi
// +build !tinygo.wasm,!wasi

package tinymem

import "reflect"

// sliceHeader returns a pointer to a reflect.SliceHeader when not TinyGo.
func sliceHeader(ptr uintptr, size uint32) *reflect.SliceHeader {
	return &reflect.SliceHeader{Data: ptr, Len: int(size), Cap: int(size)}
}
