package types

import (
	"errors"

	"github.com/tchap/go-wrpc/wube"
)

var ErrFlagIndexOverflow = errors.New("flag index overflow")

type Flags struct {
	count uint
	bs    []byte
}

func NewFlags(count uint) Flags {
	n := count / 8
	if count%8 > 0 {
		n++
	}
	return Flags{
		count: count,
		bs:    make([]byte, n),
	}
}

// Count returns the flag count these flags are expected to hold.
func (flags Flags) Count() uint {
	return flags.count
}

// Bytes returns the underlying byte slice.
func (flags Flags) Bytes() []byte {
	return flags.bs
}

// IsSetAt will return false when index is larger than the flags capacity.
func (flags Flags) IsSetAt(index uint) bool {
	if index >= flags.count {
		return false
	}

	return flags.bs[index/8]&(byte(0b10000000)>>(index%8)) > 0
}

// SetAt sets the given flag index starting with zero.
// An error is returned when the flag is out of expected bounds.
func (flags Flags) SetAt(index uint) error {
	if index >= flags.count {
		return ErrFlagIndexOverflow
	}

	flags.bs[index/8] |= byte(0b10000000) >> (index % 8)
	return nil
}

// MustSetAt is much like SetAt, it just panics on error.
func (flags Flags) MustSetAt(index uint) {
	if err := flags.SetAt(index); err != nil {
		panic(err)
	}
}

func (flags Flags) MarshalWube(enc wube.Encoder) error {
	_, err := enc.Write(flags.bs)
	return err
}

func (flags *Flags) UnmarshalWube(dec wube.Decoder) error {
	_, err := dec.Read(flags.bs)
	return err
}
