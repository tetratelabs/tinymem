//go:build tinygo.wasm || wasi
// +build tinygo.wasm wasi

package tinymem

import "reflect"

// sliceHeader returns a pointer to a reflect.SliceHeader in TinyGo, which
// requires Len and Cap as uinptr even if they are int fields in Go.
//
// See https://github.com/tinygo-org/tinygo/issues/1364
func sliceHeader(ptr uintptr, size uint32) *reflect.SliceHeader {
	return &reflect.SliceHeader{Data: ptr, Len: uintptr(size), Cap: uintptr(size)}
}
