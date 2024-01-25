package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube/types"
)

type FlagsSuite struct {
	TypeSuite
}

func (s *FlagsSuite) TestSetAt() {
	flags := types.NewFlags(14)
	flags.SetAt(0)
	flags.SetAt(1)
	flags.SetAt(7)
	flags.SetAt(8)
	flags.SetAt(10)
	s.Equal([]byte{0b11000001, 0b10100000}, flags.Bytes())
}

func (s *FlagsSuite) TestMarshalWube() {
	flags := types.NewFlags(14)
	flags.SetAt(0)
	flags.SetAt(1)
	flags.SetAt(7)
	flags.SetAt(8)
	flags.SetAt(10)
	if s.NoError(flags.MarshalWube(s.Encoder)) {
		s.AssertBuffer([]byte{0b11000001, 0b10100000})
	}
}

func (s *FlagsSuite) TestUnmarshalWube() {
	flags := types.NewFlags(14)
	s.SetBuffer([]byte{0b11000001, 0b10100000})
	if s.NoError(flags.UnmarshalWube(s.Decoder)) {
		s.True(flags.IsSetAt(0))
		s.True(flags.IsSetAt(1))
		s.False(flags.IsSetAt(2))
		s.False(flags.IsSetAt(3))
		s.False(flags.IsSetAt(4))
		s.False(flags.IsSetAt(5))
		s.False(flags.IsSetAt(6))
		s.True(flags.IsSetAt(7))
		s.True(flags.IsSetAt(8))
		s.False(flags.IsSetAt(9))
		s.True(flags.IsSetAt(10))
		s.False(flags.IsSetAt(11))
		s.False(flags.IsSetAt(12))
		s.False(flags.IsSetAt(13))
		s.False(flags.IsSetAt(14))
	}
}

func TestFlags(t *testing.T) {
	suite.Run(t, new(FlagsSuite))
}
