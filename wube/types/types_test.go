package types_test

import (
	"bytes"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube"
)

type TypeSuite struct {
	suite.Suite
	b       *bytes.Buffer
	Encoder wube.Encoder
	Decoder wube.Decoder
}

func (s *TypeSuite) SetupTest() {
	s.b = bytes.NewBuffer(nil)
	s.Encoder = wube.NewEncoder(s.b)
	s.Decoder = wube.NewDecoder(s.b)
}

func (s *TypeSuite) SetupSubTest() {
	s.SetupTest()
}

func (s *TypeSuite) AssertBuffer(expected []byte) bool {
	return s.Equal(expected, s.b.Bytes())
}

func (s *TypeSuite) SetBuffer(bs []byte) {
	s.b.Reset()
	if _, err := s.b.Write(bs); err != nil {
		s.FailNow("failed to set buffer content")
	}
}
