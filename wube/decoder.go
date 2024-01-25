package wube

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/go-delve/delve/pkg/dwarf/leb128"
)

const DefaultMaxListSize = 10000

func Unmarshal(data []byte, v any) error {
	return NewDecoder(bytes.NewBuffer(data)).Decode(v)
}

type Decoder struct {
	b ReadBuffer

	MaxListSize int
}

func NewDecoder(b ReadBuffer) Decoder {
	return Decoder{
		b:           b,
		MaxListSize: DefaultMaxListSize,
	}
}

func (dec Decoder) Decode(v any) error {
	// Try the Unmarshaler interface.
	if v, ok := v.(Unmarshaler); ok {
		return v.UnmarshalWube(dec)
	}

	vv := reflect.ValueOf(v)
	vt := vv.Type()

	// When not a Marshaler, a pointer is always expected.
	if vt.Kind() != reflect.Pointer {
		return &ErrUnsupportedType{v: v}
	}

	// When using generics, *T fields can get passed to Decode as **T.
	// We need to dereference in that case.
	if deref := vt.Elem(); deref.Kind() == reflect.Pointer {
		vv = vv.Elem()
		vt = deref
	}

	// In case we are decoding into a nil pointer,
	// try to allocate a new object to decode into.
	if vv.IsNil() && vv.CanSet() {
		vv.Set(reflect.New(vt.Elem()))
	}

	// Point to the actual values.
	vv = reflect.Indirect(vv)
	vt = vt.Elem()

	switch vt.Kind() {
	case reflect.Uint8:
		d, err := dec.ReadUInt8()
		if err != nil {
			return err
		}

		vv.SetUint(uint64(d))
		return nil

	case reflect.Uint16:
		d, err := dec.ReadUInt16()
		if err != nil {
			return err
		}

		vv.SetUint(uint64(d))
		return nil

	case reflect.Uint32:
		d, err := dec.ReadUInt32()
		if err != nil {
			return err
		}

		vv.SetUint(uint64(d))
		return nil

	case reflect.Uint64:
		d, err := dec.ReadUInt64()
		if err != nil {
			return err
		}

		vv.SetUint(uint64(d))
		return nil

	case reflect.Int8:
		d, err := dec.ReadInt8()
		if err != nil {
			return err
		}

		vv.SetInt(int64(d))
		return nil

	case reflect.Int16:
		d, err := dec.ReadInt16()
		if err != nil {
			return err
		}

		vv.SetInt(int64(d))
		return nil

	case reflect.Int32:
		d, err := dec.ReadInt32()
		if err != nil {
			return err
		}

		vv.SetInt(int64(d))
		return nil

	case reflect.Int64:
		d, err := dec.ReadInt64()
		if err != nil {
			return err
		}

		vv.SetInt(int64(d))
		return nil

	case reflect.Float32:
		d, err := dec.ReadFloat32()
		if err != nil {
			return err
		}

		vv.SetFloat(float64(d))
		return nil

	case reflect.Float64:
		d, err := dec.ReadFloat64()
		if err != nil {
			return err
		}

		vv.SetFloat(d)
		return nil

	case reflect.Bool:
		d, err := dec.ReadBool()
		if err != nil {
			return err
		}

		vv.SetBool(d)
		return nil

	case reflect.String:
		d, err := dec.ReadString()
		if err != nil {
			return err
		}

		vv.SetString(d)
		return nil

	case reflect.Struct:
		return dec.readStruct(vt, vv)

	case reflect.Slice:
		return dec.readSlice(vt, vv)

	default:
		return &ErrUnsupportedType{v: v}
	}
}

func (dec Decoder) readStruct(t reflect.Type, v reflect.Value) error {
	for i := 0; i < t.NumField(); i++ {
		// Skip unexported fields.
		if !t.Field(i).IsExported() {
			continue
		}

		if err := dec.Decode(v.Field(i).Addr().Interface()); err != nil {
			return err
		}
	}
	return nil
}

func (dec Decoder) readSlice(t reflect.Type, v reflect.Value) error {
	// Read list length.
	n, err := dec.ReadLen()
	if err != nil {
		return err
	}
	if n > dec.MaxListSize {
		return &ErrMaxListSizeExceeded{n}
	}

	// Allocate the given slice.
	xs := reflect.MakeSlice(t, n, n)

	// Read items.
	for i := 0; i < n; i++ {
		if err := dec.Decode(xs.Index(i).Addr().Interface()); err != nil {
			return err
		}
	}
	v.Set(xs)
	return nil
}

func (dec Decoder) readUnsignedNumber() (uint64, error) {
	b := newLEB128Reader(dec.b)
	v, _ := leb128.DecodeUnsigned(&b)
	return v, b.Err()
}

func (dec Decoder) readSignedNumber() (int64, error) {
	b := newLEB128Reader(dec.b)
	v, _ := leb128.DecodeSigned(&b)
	return v, b.Err()
}

func (dec Decoder) Read(bs []byte) (int, error) {
	return dec.b.Read(bs)
}

func (dec Decoder) ReadByte() (byte, error) {
	return dec.b.ReadByte()
}

// ReadUInt8 is just a proxy for ReadByte.
func (dec Decoder) ReadUInt8() (byte, error) {
	return dec.ReadByte()
}

func (dec Decoder) ReadUInt16() (uint16, error) {
	v, err := dec.readUnsignedNumber()
	return uint16(v), err
}

func (dec Decoder) ReadUInt32() (uint32, error) {
	v, err := dec.readUnsignedNumber()
	return uint32(v), err
}

func (dec Decoder) ReadUInt64() (uint64, error) {
	return dec.readUnsignedNumber()
}

// ReadInt8 is just a proxy for ReadByte.
func (dec Decoder) ReadInt8() (byte, error) {
	return dec.ReadByte()
}

func (dec Decoder) ReadInt16() (int16, error) {
	v, err := dec.readSignedNumber()
	return int16(v), err
}

func (dec Decoder) ReadInt32() (int32, error) {
	v, err := dec.readSignedNumber()
	return int32(v), err
}

func (dec Decoder) ReadInt64() (int64, error) {
	return dec.readSignedNumber()
}

func (dec Decoder) ReadFloat32() (float32, error) {
	var v float32
	err := binary.Read(dec.b, encoding, &v)
	return v, err
}

func (dec Decoder) ReadFloat64() (float64, error) {
	var v float64
	err := binary.Read(dec.b, encoding, &v)
	return v, err
}

func (dec Decoder) ReadBool() (bool, error) {
	d, err := dec.b.ReadByte()
	return d > 0, err
}

func (dec Decoder) ReadString() (string, error) {
	// Strings are length-prefixed UTF8-encoded byte arrays.
	n, err := dec.ReadLen()
	if err != nil {
		return "", err
	}
	if n > dec.MaxListSize {
		return "", &ErrMaxListSizeExceeded{n}
	}

	bs := make([]byte, n)
	if _, err := io.ReadFull(dec.b, bs); err != nil {
		return "", err
	}
	return string(bs), nil
}

func (dec Decoder) ReadLen() (int, error) {
	n, err := dec.ReadUInt32()
	return int(n), err
}

func (dec Decoder) ReadEnum(dMax uint64) (uint64, error) {
	switch {
	case dMax <= maxUInt8:
		d, err := dec.ReadUInt8()
		return uint64(d), err

	case dMax <= maxUInt16:
		d, err := dec.ReadUInt16()
		return uint64(d), err

	case dMax <= maxUInt32:
		d, err := dec.ReadUInt32()
		return uint64(d), err

	default:
		return dec.ReadUInt64()
	}
}
