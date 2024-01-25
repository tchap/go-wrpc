# Wube encoding and decoding

Based on https://github.com/wasmCloud/wasmCloud/issues/1128

The functionality present in this package is pretty raw and mostly suitable to be used from
generated code, although nothing really prevents people from encoding and decoding the more complex
types manually.

## Status

I am playing with Wube and how to implement it in Go, so I am constantly reworking things.

Encoding:

- [x] Numeric types
- [x] Characters
- [x] Strings
- [x] Booleans
- [x] Enums - Only helpers for code generation available as `Encoder.WriteEnum`.
- [x] Flags - Only helpers for code generation available as `types.Flags`.
- [x] Lists
- [x] Tuples - Implemented as `types.{V1, V2, V3, V4, V5}` (implements `wube.Marshaler`).
- [x] Options - Implemented as `types.Option` (implements `wube.Marshaler`).
- [x] Results - Implemented as `types.Result` (implements `wube.Marshaler`).
- [x] Futures - Implemented as `types.Future` (implements `wube.Marshaler`).
- [x] Records

Decoding:

- [x] Numeric types
- [x] Characters
- [x] Strings
- [x] Booleans
- [x] Enums - Only helpers for code generation available as `Decoder.ReadEnum`.
- [x] Flags - Only helpers for code generation available as `types.Flags`.
- [x] Lists
- [x] Tuples - Implemented as `types.{V1, V2, V3, V4, V5}` (implements `wube.Unmarshaler`).
- [x] Options - Implemented as `types.Option` (implements `wube.Unarshaler`).
- [x] Results - Implemented as `types.Result` (implements `wube.Unarshaler`).
- [x] Futures - Implemented as `types.Future` (implements `wube.Unmarshaler`)
- [x] Records

Types to be handled by code generation using the logic implemented in this package:

- Enums - We need to know max discriminant value when decoding,
  which is not really possible when decoding into a zero value of any type.
- Variants - These must be inherently handled by code generation in much the same way
  oneof is implemented in Protocol Buffers for Go.

### Streams, resources and functions

These are not really encoded or decoded directly as objects on the language level,
they are concepts implemented on the wRPC level.
