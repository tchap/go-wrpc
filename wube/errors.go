package wube

import (
	"fmt"
)

type ErrUnsupportedType struct {
	v interface{}
}

func (err *ErrUnsupportedType) Value() interface{} {
	return err.v
}

func (err *ErrUnsupportedType) Error() string {
	return fmt.Sprintf("wube: unsupported type: %T", err.v)
}

type ErrMaxListSizeExceeded struct {
	n int
}

func (err *ErrMaxListSizeExceeded) Error() string {
	return fmt.Sprintf("wube: max list size exceeded: %d", err.n)
}
