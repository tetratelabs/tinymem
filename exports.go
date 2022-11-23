package tinymem

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
	// TODO: how should we handle zero?
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
	// intentionally corrupt first byte.
	if b := alivePointers[ptr]; b != nil && len(b) > 0 {
		b[0] = '?'
	}
	delete(alivePointers, ptr)
}
