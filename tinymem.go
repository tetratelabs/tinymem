package tinymem

import "unsafe"

// PtrToString returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func PtrToString(ptr uintptr, size uint32) string {
	// Get a slice view of the underlying bytes in the stream. We use
	// SliceHeader, not StringHeader as it allows us to fix the capacity to
	// what was allocated.
	return *(*string)(unsafe.Pointer(sliceHeader(ptr, size)))
}

// StringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
func StringToPtr(s string) (uintptr, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return unsafePtr, uint32(len(buf))
}
