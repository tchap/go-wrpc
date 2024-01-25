package wube_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EncoderSuite struct {
	WubeSuite
}

func (s *EncoderSuite) TestEncode_Bool_CustomType() {
	type MyBool bool
	v := MyBool(true)
	if s.NoError(s.Encoder.Encode(v)) {
		s.Equal([]byte{0x01}, s.b.Bytes())
	}
}

func (s *EncoderSuite) TestEncode_Slice_Bool() {
	if s.NoError(s.Encoder.Encode([]bool{true, false})) {
		s.AssertBuffer([]byte{0x02, 0x01, 0x00})
	}
}

func (s *EncoderSuite) TestWriteUInt32() {
	if s.NoError(s.Encoder.WriteUInt32(1)) {
		s.Equal([]byte{0x01}, s.b.Bytes())
	}
}

func (s *EncoderSuite) TestWriteUInt64() {
	if s.NoError(s.Encoder.WriteUInt64(1)) {
		s.Equal([]byte{0x01}, s.b.Bytes())
	}
}

func (s *EncoderSuite) TestWriteRune() {
	err := s.Encoder.WriteChar('a')
	if s.NoError(err) {
		s.Equal([]byte{0x61}, s.b.Bytes())
	}
}

func (s *EncoderSuite) TestWriteBool() {
	s.Run("true", func() {
		if s.NoError(s.Encoder.WriteBool(true)) {
			s.Equal([]byte{1}, s.b.Bytes())
		}
	})

	s.Run("false", func() {
		if s.NoError(s.Encoder.WriteBool(false)) {
			s.Equal([]byte{0}, s.b.Bytes())
		}
	})
}

func TestEncoder(t *testing.T) {
	suite.Run(t, new(EncoderSuite))
}
