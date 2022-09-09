## TinyGo allocation example

This example shows how to pass strings in and out of a Wasm function defined
in TinyGo, built via

```bash
$ tinygo build -o greeting.wasm -scheduler=none -target=wasi greeting.go
```

Notably, this exports functions the host can use to pass a string parameter via
ptr, len pairs. In WebAssembly, the exported signatures look like this:

```webassembly
(func (export "_malloc") (param $size i32) (result (;$ptr;) i32))
(func (export "_free") (param $ptr i32))
```