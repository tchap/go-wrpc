package http_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tchap/go-wrpc/wube/examples"
	"github.com/tchap/go-wrpc/wube/examples/http"
	"github.com/tchap/go-wrpc/wube/types"
)

type Headers = []types.T2[string, []byte]

type TypesSuite struct {
	examples.Suite
}

func (s *TypesSuite) TestRequestOptions_Encode() {
	opts := http.RequestOptions{
		ConnectTimeout:      1000,
		FirstByteTimeout:    2000,
		BetweenBytesTimeout: 500,
	}

	if s.NoError(s.Encoder.Encode(&opts)) {
		s.AssertBuffer([]byte{
			0xe8, 0x7, // 1000
			0xd0, 0xf, // 2000
			0xf4, 0x3, // 500
		})
	}
}

func (s *TypesSuite) TestRequestOptions_Decode() {
	s.SetBuffer([]byte{
		0xe8, 0x7, // 1000
		0xd0, 0xf, // 2000
		0xf4, 0x3, // 500
	})

	var opts http.RequestOptions
	if s.NoError(s.Decoder.Decode(&opts)) {
		s.Equal(http.RequestOptions{
			ConnectTimeout:      1000,
			FirstByteTimeout:    2000,
			BetweenBytesTimeout: 500,
		}, opts)
	}
}

func (s *TypesSuite) TestOutgoingRequest_Encode() {
	req := http.OutgoingRequest{
		Headers: Headers{
			{"Content-Type", []byte("application/json")},
		},
		Method:        http.MethodPost,
		PathWithQuery: types.NewSomeOption[string]("/api/v1/accounts"),
		Scheme:        types.NewSomeOption[http.Scheme](http.NewScheme_Https()),
		Body:          []byte(`{"name": "John"}`),
	}

	if s.NoError(s.Encoder.Encode(&req)) {
		s.AssertBuffer([]byte{
			0x1,                                                                    // len(Headers)
			0xc,                                                                    // len(Headers[0].V1)
			0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79, 0x70, 0x65, // Headers[0].V1
			0x10,                                                                                           // len(Headers[0].V2)
			0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, // Headers[0].V2
			0x2,                                                                                            // Method discriminant
			0x1,                                                                                            // PathWithQuery discriminant (some)
			0x10,                                                                                           // len(PathWithQuery)
			0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, // PathWithQuery
			0x1,                                                                                            // Scheme option discriminant (some)
			0x1,                                                                                            // Scheme variant discriminant (https)
			0x0,                                                                                            // Authority discriminant (none)
			0x10,                                                                                           // len(Body)
			0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x68, 0x6e, 0x22, 0x7d, // Body
			0x0, // Trailers discriminant (none)
		})
	}
}

func (s *TypesSuite) TestOutgoingRequest_Decode() {
	s.SetBuffer([]byte{
		0x1,                                                                    // len(Headers)
		0xc,                                                                    // len(Headers[0].V1)
		0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79, 0x70, 0x65, // Headers[0].V1
		0x10,                                                                                           // len(Headers[0].V2)
		0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, // Headers[0].V2
		0x2,                                                                                            // Method discriminant
		0x1,                                                                                            // PathWithQuery discriminant (some)
		0x10,                                                                                           // len(PathWithQuery)
		0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, // PathWithQuery
		0x1,                                                                                            // Scheme option discriminant (some)
		0x1,                                                                                            // Scheme variant discriminant (https)
		0x0,                                                                                            // Authority discriminant (none)
		0x10,                                                                                           // len(Body)
		0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x68, 0x6e, 0x22, 0x7d, // Body
		0x0, // Trailers discriminant (none)
	})

	var req http.OutgoingRequest
	if s.NoError(s.Decoder.Decode(&req)) {
		s.Equal(http.OutgoingRequest{
			Headers: Headers{
				{"Content-Type", []byte("application/json")},
			},
			Method:        http.MethodPost,
			PathWithQuery: types.NewSomeOption[string]("/api/v1/accounts"),
			Scheme:        types.NewSomeOption[http.Scheme](http.NewScheme_Https()),
			Body:          []byte(`{"name": "John"}`),
		}, req)
	}
}

func (s *TypesSuite) TestIncomingResponse_Encode() {
	resp := http.IncomingResponse{
		Status: 201,
		Headers: []types.T2[string, []byte]{
			{"Content-Type", []byte("application/json")},
		},
		Body: []byte(`{"id": 1, "name": "John"}`),
	}

	if s.NoError(s.Encoder.Encode(&resp)) {
		s.AssertBuffer([]byte{
			0xc9, 0x1, // StatusCode
			0x1,                                                                    // len(Headers)
			0xc,                                                                    // len(Headers[0].V1)
			0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79, 0x70, 0x65, // Headers[0].V1
			0x10,                                                                                           // len(Headers[0].V2)
			0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, // Headers[0].V2
			0x19,                                                                                                                                                 // len(Body)
			0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x31, 0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x68, 0x6e, 0x22, 0x7d, // Body
			0x0, // Trailers discriminant (none)
		})
	}
}

func (s *TypesSuite) TestIncomingResponse_Decode() {
	s.SetBuffer([]byte{
		0xc9, 0x1, // StatusCode
		0x1,                                                                    // len(Headers)
		0xc,                                                                    // len(Headers[0].V1)
		0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x54, 0x79, 0x70, 0x65, // Headers[0].V1
		0x10,                                                                                           // len(Headers[0].V2)
		0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, // Headers[0].V2
		0x19,                                                                                                                                                 // len(Body)
		0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x31, 0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4a, 0x6f, 0x68, 0x6e, 0x22, 0x7d, // Body
		0x0, // Trailers discriminant (none)
	})

	var resp http.IncomingResponse
	if s.NoError(s.Decoder.Decode(&resp)) {
		s.Equal(http.IncomingResponse{
			Status: 201,
			Headers: Headers{
				{"Content-Type", []byte("application/json")},
			},
			Body: []byte(`{"id": 1, "name": "John"}`),
		}, resp)
	}
}

func TestHTTP(t *testing.T) {
	suite.Run(t, new(TypesSuite))
}
