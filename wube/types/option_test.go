package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube/types"
)

type OptionSuite struct {
	TypeSuite
}

func (s *OptionSuite) TestMarshal_None() {
	v := types.NewNoneOption[uint32]()
	if s.NoError(v.MarshalWube(s.Encoder)) {
		s.AssertBuffer([]byte{0x00})
	}
}

func (s *OptionSuite) TestMarshal_Some_UInt32() {
	v := types.NewSomeOption[uint32](3)
	if s.NoError(v.MarshalWube(s.Encoder)) {
		s.AssertBuffer([]byte{0x01, 0x03})
	}
}

func (s *OptionSuite) TestUnmarshal_None() {
	var v types.Option[uint32]
	s.SetBuffer([]byte{0x00})
	if s.NoError(v.UnmarshalWube(s.Decoder)) {
		s.True(v.IsNone())
		s.False(v.IsSome())
	}
}

func (s *OptionSuite) TestUnmarshal_Some_UInt32() {
	var v types.Option[uint32]
	s.SetBuffer([]byte{0x01, 0x03})
	if s.NoError(v.UnmarshalWube(s.Decoder)) {
		some, ok := v.Some()
		s.Equal(uint32(3), some)
		s.True(ok)
		s.False(v.IsNone())
	}
}

func TestOption(t *testing.T) {
	suite.Run(t, new(OptionSuite))
}
