[![Build](https://github.com/tetratelabs/tinymem/workflows/build/badge.svg)](https://github.com/tetratelabs/tinymem)
[![Go Report Card](https://goreportcard.com/badge/github.com/tetratelabs/tinymem)](https://goreportcard.com/report/github.com/tetratelabs/tinymem)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

# TinyMem

TinyMem is a collection of WebAssembly compatible memory utilities for TinyGo.
This allows you to implement tested best practices, without copy/paste!

Check out our [example](example), which lets the WebAssembly host safely use
TinyGo's memory allocator, in order to pass strings which aren't natively
supported in WebAssembly.