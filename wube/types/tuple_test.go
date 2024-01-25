package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube/types"
)

type TupleSuite struct {
	TypeSuite
}

func (s *TupleSuite) TestV1_Bool() {
	s.Run("Marshal", func() {
		v := types.T1[bool]{V1: true}
		if s.NoError(s.Encoder.Encode(&v)) {
			s.AssertBuffer([]byte{0x01})
		}
	})

	s.Run("Unmarshal", func() {
		var v types.T1[bool]
		s.SetBuffer([]byte{0x01})
		if s.NoError(s.Decoder.Decode(&v)) {
			s.True(v.V1)
		}
	})
}

func (s *TupleSuite) TestV1_Bool_Pointer() {
	s.Run("Marshal", func() {
		boolValue := true
		v := types.T1[*bool]{V1: &boolValue}
		if s.NoError(s.Encoder.Encode(&v)) {
			s.AssertBuffer([]byte{0x01})
		}
	})

	s.Run("Unmarshal", func() {
		var v types.T1[*bool]
		s.SetBuffer([]byte{0x01})
		if s.NoError(s.Decoder.Decode(&v)) {
			s.True(*v.V1)
		}
	})
}

func TestTuple(t *testing.T) {
	suite.Run(t, new(TupleSuite))
}
