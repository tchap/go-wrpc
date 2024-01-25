package examples

import (
	"bytes"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube"
)

type Suite struct {
	suite.Suite
	b       *bytes.Buffer
	Encoder wube.Encoder
	Decoder wube.Decoder
}

func (s *Suite) SetupTest() {
	s.b = bytes.NewBuffer(nil)
	s.Encoder = wube.NewEncoder(s.b)
	s.Decoder = wube.NewDecoder(s.b)
}

func (s *Suite) SetupSubTest() {
	s.SetupTest()
}

func (s *Suite) AssertBuffer(expected []byte) bool {
	return s.Equal(expected, s.b.Bytes())
}

func (s *Suite) SetBuffer(bs []byte) {
	s.b.Reset()
	if _, err := s.b.Write(bs); err != nil {
		s.FailNow("failed to set buffer content")
	}
}
