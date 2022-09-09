// Package exports contains WebAssembly function exports for memory allocation.
// It is not necessary to import these otherwise.
//
// Ex.
//
//	import _ github.com/tetratelabs/tinymem/exports
package exports

import "unsafe"

var alivePointers = map[uintptr][]byte{}

// malloc is a WebAssembly export named "_malloc" that allocates a pointer
// (linear memory offset) that can be used for the given size in bytes.
//
// Note: This is an ownership transfer, which means the caller must call free
// when finished.
//
//export _malloc
func malloc(size uint32) uintptr {
	buf := make([]byte, size)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	alivePointers[unsafePtr] = buf
	return unsafePtr
}

// free is a WebAssembly export named "_free" that deallocates a pointer
// allocated by malloc.
//
//export _free
func free(ptr uintptr) {
	delete(alivePointers, ptr)
}
