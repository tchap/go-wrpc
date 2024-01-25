package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube/types"
)

type FutureSuite struct {
	TypeSuite
}

func (s *FutureSuite) TestMarshal_Pending() {
	v := types.NewPendingFuture[uint32]()
	if s.NoError(v.MarshalWube(s.Encoder)) {
		s.AssertBuffer([]byte{0x00})
	}
}

func (s *FutureSuite) TestMarshal_Ready_UInt32() {
	v := types.NewReadyFuture[uint32](3)
	if s.NoError(v.MarshalWube(s.Encoder)) {
		s.AssertBuffer([]byte{0x01, 0x03})
	}
}

func (s *FutureSuite) TestUnmarshal_Pending() {
	var v types.Future[uint32]
	s.SetBuffer([]byte{0x00})
	if s.NoError(v.UnmarshalWube(s.Decoder)) {
		s.True(v.IsPending())
		s.False(v.IsReady())
	}
}

func (s *FutureSuite) TestUnmarshal_Ready_UInt32() {
	var v types.Future[uint32]
	s.SetBuffer([]byte{0x01, 0x03})
	if s.NoError(v.UnmarshalWube(s.Decoder)) {
		ready, ok := v.Ready()
		s.Equal(uint32(3), ready)
		s.True(ok)
		s.False(v.IsPending())
	}
}

func TestFuture(t *testing.T) {
	suite.Run(t, new(FutureSuite))
}
