package types

import "github.com/tchap/go-wrpc/wube"

type Future[T any] struct {
	ready      bool
	readyValue T
}

func NewPendingFuture[T any]() Future[T] {
	return Future[T]{}
}

func NewReadyFuture[T any](value T) Future[T] {
	return Future[T]{
		ready:      true,
		readyValue: value,
	}
}

func (fut Future[T]) IsPending() bool {
	return !fut.ready
}

func (fut Future[T]) IsReady() bool {
	return fut.ready
}

func (fut Future[T]) Ready() (T, bool) {
	return fut.readyValue, fut.ready
}

func (fut Future[T]) MarshalWube(enc wube.Encoder) error {
	if !fut.ready {
		return enc.WriteByte(0)
	}

	if err := enc.WriteByte(1); err != nil {
		return err
	}
	return enc.Encode(fut.readyValue)
}

func (fut *Future[T]) UnmarshalWube(dec wube.Decoder) error {
	d, err := dec.ReadUInt8()
	if err != nil {
		return err
	}

	if d == 0 {
		return nil
	}
	err = dec.Decode(&fut.readyValue)
	fut.ready = err == nil
	return err
}
