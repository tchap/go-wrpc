package types

import (
	"errors"
	"fmt"

	"github.com/tchap/go-wrpc/wube"
)

type Result[Ok, Err any] struct {
	okValue  Ok
	okSet    bool
	errValue Err
	errSet   bool
}

func NewOkResult[Ok, Err any](value Ok) Result[Ok, Err] {
	return Result[Ok, Err]{
		okValue: value,
		okSet:   true,
	}
}

func NewErrResult[Ok, Err any](value Err) Result[Ok, Err] {
	return Result[Ok, Err]{
		errValue: value,
		errSet:   true,
	}
}

func (res Result[Ok, Err]) IsOk() bool {
	return res.okSet
}

func (res Result[Ok, Err]) Ok() (Ok, bool) {
	return res.okValue, res.okSet
}

func (res Result[Ok, Err]) IsErr() bool {
	return res.errSet
}

func (res Result[Ok, Err]) Err() (Err, bool) {
	return res.errValue, res.errSet
}

func (res Result[Ok, Err]) MarshalWube(enc wube.Encoder) error {
	if res.okSet {
		return res.marshal(enc, 0, res.okValue)
	}

	if res.errSet {
		return res.marshal(enc, 1, res.errValue)
	}

	return errors.New("failed to encode result: cannot encode a zero result")
}

func (res Result[Ok, Err]) marshal(enc wube.Encoder, d byte, v any) error {
	if err := enc.WriteByte(d); err != nil {
		return err
	}
	return enc.Encode(v)
}

func (res *Result[Ok, Err]) UnmarshalWube(dec wube.Decoder) error {
	d, err := dec.ReadUInt8()
	if err != nil {
		return err
	}

	switch d {
	case 0:
		if err := dec.Decode(&res.okValue); err != nil {
			return err
		}
		res.okSet = true
		res.errSet = false
		return nil

	case 1:
		if err := dec.Decode(&res.errValue); err != nil {
			return err
		}
		res.okSet = false
		res.errSet = true
		return nil

	default:
		return fmt.Errorf("failed to decode result: discriminant out of range: %d", d)
	}
}
