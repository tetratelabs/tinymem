package tinymem

import (
	"testing"
	"unsafe"
)

func Test_mallocRoundTrip(t *testing.T) {
	t.Run("free when not yet allocated", func(t *testing.T) {
		free(123) // doesn't panic
	})

	size := uint32(3)

	ptr := malloc(size)
	t.Run("malloc", func(t *testing.T) {
		buf := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size)
		buf[0] = 'f'
		buf[1] = 'o'
		buf[2] = 'o'

		expected := "foo"
		if have := PtrToString(ptr, size); expected != have {
			t.Errorf("Unexpected string, have %s, expected %s", have, expected)
		}
	})

	t.Run("free", func(t *testing.T) {
		free(ptr)

		if _, ok := alivePointers[ptr]; ok {
			t.Errorf("Unexpected to still have a string")
		}
	})
}
