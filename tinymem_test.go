package tinymem_test

import (
	"testing"

	"github.com/tetratelabs/tinymem"
)

func Test_StringRoundTrip(t *testing.T) {
	expected := "foo"
	ptr, size := tinymem.StringToPtr(expected)

	if have := tinymem.PtrToString(ptr, size); have != expected {
		t.Errorf("Unexpected string, have %s, expected %s", have, expected)
	}
}
