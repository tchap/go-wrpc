package wube

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/go-delve/delve/pkg/dwarf/leb128"
)

var encoding = binary.LittleEndian

func Marshal(v any) ([]byte, error) {
	// Compute the resulting size.
	sizeBuffer := NewSizeBuffer()
	if err := NewEncoder(sizeBuffer).Encode(v); err != nil {
		return nil, err
	}

	// Decode into a pre-allocated buffer.
	bs := make([]byte, sizeBuffer.Size())
	if err := NewEncoder(bytes.NewBuffer(bs)).Encode(v); err != nil {
		return nil, err
	}
	return bs, nil
}

type Encoder struct {
	b WriteBuffer
}

func NewEncoder(b WriteBuffer) Encoder {
	return Encoder{b: b}
}

// Encode encodes the given value based on type.
//
// Note that Go's rune is int32 while Wube distates a rune to be uint32.
// Which means Encoder.WriteChar must be used explicitly or rune case to uint32.
func (enc Encoder) Encode(v any) error {
	if v, ok := v.(Marshaler); ok {
		return v.MarshalWube(enc)
	}

	vv := reflect.Indirect(reflect.ValueOf(v))
	switch vv.Kind() {
	case reflect.Uint8:
		return enc.WriteUInt8(uint8(vv.Uint()))

	case reflect.Uint16:
		return enc.WriteUInt16(uint16(vv.Uint()))

	case reflect.Uint32:
		return enc.WriteUInt32(uint32(vv.Uint()))

	case reflect.Uint64:
		return enc.WriteUInt64(vv.Uint())

	case reflect.Int8:
		return enc.WriteInt8(int8(vv.Int()))

	case reflect.Int16:
		return enc.WriteInt16(int16(vv.Int()))

	case reflect.Int32:
		return enc.WriteInt32(int32(vv.Int()))

	case reflect.Int64:
		return enc.WriteInt64(vv.Int())

	case reflect.Float32:
		return enc.WriteFloat32(float32(vv.Float()))

	case reflect.Float64:
		return enc.WriteFloat64(vv.Float())

	case reflect.Bool:
		return enc.WriteBool(vv.Bool())

	case reflect.String:
		return enc.WriteString(vv.String())

	case reflect.Struct:
		return enc.writeRecord(vv)

	case reflect.Slice:
		return enc.writeList(vv)

	default:
		return &ErrUnsupportedType{v: v}
	}
}

func (enc Encoder) writeRecord(v reflect.Value) error {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// Skip unexported fields.
		if !t.Field(i).IsExported() {
			continue
		}

		if err := enc.Encode(v.Field(i).Interface()); err != nil {
			return err
		}
	}
	return nil
}

func (enc Encoder) writeList(v reflect.Value) error {
	if err := enc.WriteLen(uint64(v.Len())); err != nil {
		return err
	}

	for i := 0; i < v.Len(); i++ {
		if err := enc.Encode(v.Index(i).Interface()); err != nil {
			return err
		}
	}
	return nil
}

func (enc Encoder) writeUnsignedNumber(v uint64) error {
	b := newLEB128Writer(enc.b)
	leb128.EncodeUnsigned(&b, v)
	return b.Err()
}

func (enc Encoder) writeSignedNumber(v int64) error {
	b := newLEB128Writer(enc.b)
	leb128.EncodeSigned(&b, v)
	return b.Err()
}

func (enc Encoder) Write(bs []byte) (int, error) {
	return enc.b.Write(bs)
}

func (enc Encoder) WriteByte(v byte) error {
	return enc.b.WriteByte(v)
}

// WriteUInt8 just calls WriteByte.
func (enc Encoder) WriteUInt8(v uint8) error {
	return enc.WriteByte(v)
}

func (enc Encoder) WriteUInt16(v uint16) error {
	return enc.writeUnsignedNumber(uint64(v))
}

func (enc Encoder) WriteUInt32(v uint32) error {
	return enc.writeUnsignedNumber(uint64(v))
}

func (enc Encoder) WriteUInt64(v uint64) error {
	return enc.writeUnsignedNumber(uint64(v))
}

// WriteInt8 just calls WriteByte.
func (enc Encoder) WriteInt8(v int8) error {
	return enc.WriteByte(byte(v))
}

func (enc Encoder) WriteInt16(v int16) error {
	return enc.writeSignedNumber(int64(v))
}

func (enc Encoder) WriteInt32(v int32) error {
	return enc.writeSignedNumber(int64(v))
}

func (enc Encoder) WriteInt64(v int64) error {
	return enc.writeSignedNumber(int64(v))
}

func (enc Encoder) WriteFloat32(v float32) error {
	return binary.Write(enc.b, encoding, v)
}

func (enc Encoder) WriteFloat64(v float64) error {
	return binary.Write(enc.b, encoding, v)
}

func (enc Encoder) WriteChar(v rune) error {
	return enc.WriteUInt32(uint32(v))
}

func (enc Encoder) WriteBool(v bool) error {
	if v {
		return enc.b.WriteByte(1)
	}
	return enc.b.WriteByte(0)
}

func (enc Encoder) WriteString(v string) error {
	// Strings are length-prefixed UTF8-encoded byte arrays.
	if err := enc.WriteLen(uint64(len(v))); err != nil {
		return err
	}

	_, err := io.WriteString(enc.b, v)
	return err
}

func (enc Encoder) WriteEnum(d uint64) error {
	return enc.WriteUInt64(d)
}

func (enc Encoder) WriteVariant(d uint64, v any) error {
	if err := enc.WriteEnum(d); err != nil {
		return err
	}
	return enc.Encode(v)
}

func (enc Encoder) WriteLen(n uint64) error {
	return enc.WriteUInt64(n)
}
