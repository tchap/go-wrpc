package types

import "github.com/tchap/go-wrpc/wube"

type Option[T any] struct {
	isSet bool
	value T
}

func NewNoneOption[T any]() Option[T] {
	return Option[T]{}
}

func NewSomeOption[T any](value T) Option[T] {
	return Option[T]{
		isSet: true,
		value: value,
	}
}

func (o Option[T]) IsNone() bool {
	return !o.isSet
}

func (o Option[T]) IsSome() bool {
	return o.isSet
}

func (o Option[T]) Some() (T, bool) {
	return o.value, o.isSet
}

func (o Option[T]) MarshalWube(enc wube.Encoder) error {
	if !o.isSet {
		return enc.WriteByte(0)
	}

	if err := enc.WriteByte(1); err != nil {
		return err
	}
	return enc.Encode(o.value)
}

func (o *Option[T]) UnmarshalWube(dec wube.Decoder) error {
	d, err := dec.ReadUInt8()
	if err != nil {
		return err
	}

	if d == 0 {
		return nil
	}
	err = dec.Decode(&o.value)
	o.isSet = err == nil
	return err
}
