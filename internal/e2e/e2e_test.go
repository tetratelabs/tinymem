//go:build !tinygo.wasm && !wasi

package e2e

import (
	"context"
	_ "embed"
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero/api"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

// testCtx is an arbitrary, non-default context. Non-nil also prevents linter errors.
var testCtx = context.WithValue(context.Background(), struct{}{}, "arbitrary")

var guestWasm map[string][]byte

const (
	guestWasmExample = "example"
)

// TestMain ensures we can read the test wasm prior to running e2e tests.
func TestMain(m *testing.M) {
	wasms := []string{guestWasmExample}
	guestWasm = make(map[string][]byte, len(wasms))

	if exampleWasm, err := os.ReadFile(path.Join("..", "..", "example", "greeting.wasm")); err != nil {
		log.Panicln(err)
	} else {
		guestWasm[guestWasmExample] = exampleWasm
	}

	os.Exit(m.Run())
}

func Test_EndToEnd(t *testing.T) {
	type testCase struct {
		name  string
		guest []byte
		test  func(t *testing.T, guest api.Module)
	}

	tests := []testCase{
		{
			name:  "Example",
			guest: guestWasm[guestWasmExample],
			test: func(t *testing.T, guest api.Module) {

				// Get references to WebAssembly functions we'll use in this example.
				greeting := guest.ExportedFunction("greeting")
				malloc := guest.ExportedFunction("_malloc")
				free := guest.ExportedFunction("_free")

				name := "foo"
				nameSize := uint64(len(name))

				// Instead of an arbitrary memory offset, use TinyGo's allocator. Notice
				// there is nothing string-specific in this allocation function. The same
				// function could be used to pass binary serialized data to Wasm.
				results, err := malloc.Call(testCtx, nameSize)
				require.NoError(t, err)

				namePtr := results[0]
				// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
				// So, we have to free it when finished
				defer free.Call(testCtx, namePtr)

				// The pointer is a linear memory offset, which is where we write the name.
				ok := guest.Memory().Write(uint32(namePtr), []byte(name))
				require.True(t, ok, "out of memory writing %s", name)

				// Finally, we get the greeting message "greet" printed. This shows how to
				// read-back something allocated by TinyGo.
				ptrSize, err := greeting.Call(testCtx, namePtr, nameSize)
				require.NoError(t, err)

				// Note: This pointer is still owned by TinyGo, so don't try to free it!
				greetingPtr := uint32(ptrSize[0] >> 32)
				greetingSize := uint32(ptrSize[0])

				// The pointer is a linear memory offset, which is where we write the name.
				bytes, ok := guest.Memory().Read(greetingPtr, greetingSize)
				require.True(t, ok, "out of memory reading greeting(%d, %d)", greetingPtr, greetingSize)

				require.Equal(t, "Hello, foo!", string(bytes))
			},
		},
	}

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(testCtx)
	defer r.Close(testCtx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements system I/O such as console output and
	// is required for `tinygo build -target=wasi`
	if _, err := wasi_snapshot_preview1.Instantiate(testCtx, r); err != nil {
		t.Errorf("Error instantiating WASI - %v", err)
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			g, err := r.InstantiateModuleFromBinary(testCtx, tc.guest)
			if err != nil {
				t.Errorf("Error instantiating TinyGo guest - %v", err)
			}
			defer g.Close(testCtx)

			tc.test(t, g)
		})
	}
}
