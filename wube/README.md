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
- [x] Variants - Only helpers for code generation available as `Encoder.WriteVariant`.
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

Enums and variants are best handled using code generation, see `examples/http` for a possible
code to be generated. Notably there is no `Decoder.ReadVariant` since we need to know the object
to unmarshal based on the discriminant read. This could be handled by receiving the mapping of
discriminants to the given objects to unmarshal into, but there is little value in this anyway.

### Streams, resources and functions

These are not really encoded or decoded directly as objects on the language level,
they are concepts implemented on the wRPC level.
