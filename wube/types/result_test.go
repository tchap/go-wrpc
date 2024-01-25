package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube/types"
)

type ResultSuite struct {
	TypeSuite
}

func (s *ResultSuite) TestMarshal_Ok_UInt32() {
	v := types.NewOkResult[uint32, string](3)
	if s.NoError(v.MarshalWube(s.Encoder)) {
		s.AssertBuffer([]byte{0x00, 0x03})
	}
}

func (s *ResultSuite) TestUnmarshal_Ok_UInt32() {
	var v types.Result[uint32, string]
	s.SetBuffer([]byte{0x00, 0x03})
	if s.NoError(v.UnmarshalWube(s.Decoder)) {
		okValue, okSet := v.Ok()
		s.Equal(uint32(3), okValue)
		s.True(okSet)
		s.False(v.IsErr())
	}
}

func (s *ResultSuite) TestMarshal_Err_String() {
	v := types.NewErrResult[uint32, string]("John")
	if s.NoError(v.MarshalWube(s.Encoder)) {
		s.AssertBuffer([]byte{
			0x01, // discriminant
			0x04, // len("John")
			0x4a, // 'J'
			0x6f, // 'o'
			0x68, // 'h'
			0x6e, // 'n'
		})
	}
}

func (s *ResultSuite) TestUnmarshal_Err_String() {
	var v types.Result[uint32, string]
	s.SetBuffer([]byte{
		0x01, // discriminant
		0x04, // len("John")
		0x4a, // 'J'
		0x6f, // 'o'
		0x68, // 'h'
		0x6e, // 'n'
	})
	if s.NoError(v.UnmarshalWube(s.Decoder)) {
		errValue, errSet := v.Err()
		s.Equal("John", errValue)
		s.True(errSet)
		s.False(v.IsOk())
	}
}

func TestResult(t *testing.T) {
	suite.Run(t, new(ResultSuite))
}
