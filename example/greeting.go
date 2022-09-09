package main

// Import this to export "_malloc" and "_free" to the WebAssembly host.
import (
	"fmt"

	"github.com/tetratelabs/tinymem"
	_ "github.com/tetratelabs/tinymem/exports"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

// greeting gets a greeting for the name.
func greeting(name string) string {
	return fmt.Sprint("Hello, ", name, "!")
}

// _greeting is a WebAssembly export that accepts a string pointer (linear
// memory offset) and returns a pointer/size pair packed into a uint64.
//
// Note: This uses a uint64 instead of two result values for compatibility with
// WebAssembly types.
//
//export greeting
func _greeting(ptr uintptr, size uint32) (ptrSize uint64) {
	g := greeting(tinymem.PtrToString(ptr, size))
	ptr, size = tinymem.StringToPtr(g)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
