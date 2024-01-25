package wube_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DecoderSuite struct {
	WubeSuite
}

func (s *DecoderSuite) TestDecode_Byte() {
	var v byte
	s.SetBuffer([]byte{0x01})
	if s.NoError(s.Decoder.Decode(&v)) {
		s.Equal(byte(1), v)
	}
}

func (s *DecoderSuite) TestDecode_Bool() {
	var v bool
	s.SetBuffer([]byte{0x01})
	if s.NoError(s.Decoder.Decode(&v)) {
		s.True(bool(v))
	}
}

func (s *DecoderSuite) TestDecode_Bool_CustomType() {
	type MyBool bool
	var v MyBool
	s.SetBuffer([]byte{0x01})
	if s.NoError(s.Decoder.Decode(&v)) {
		s.True(bool(v))
	}
}

func (s *DecoderSuite) TestDecode_Slice_Bool() {
	var v []bool
	s.SetBuffer([]byte{0x02, 0x01, 0x00})
	if s.NoError(s.Decoder.Decode(&v)) {
		s.Equal([]bool{true, false}, v)
	}
}

func TestDecoder(t *testing.T) {
	suite.Run(t, new(DecoderSuite))
}
