# wrpc-bindgen-go

This utility can be used to generate Go wRPC stubs from WIT packages.

## Usage

1. Generate a JSON representation using [wasm-tools](https://crates.io/crates/wasm-tools).
2. Go to the relevant package directory and run `wasm-tools component wit -j . > wit.json`
3. Use `wrpc-bindgen-go` to generate stubs from the JSON file.
