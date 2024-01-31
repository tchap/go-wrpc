package http

import (
	"fmt"

	"github.com/tchap/go-wrpc/wube"
	"github.com/tchap/go-wrpc/wube/types"
)

// The following types as in sync with WASI-HTTP.
type (
	Duration   = uint64
	StatusCode = uint16
)

// ErrorCode represents an error code.
//
// WASI-HTTP uses a huge variant that would be insane to implement manually,
// so just falling back to using a string here.
//
// https://github.com/WebAssembly/wasi-http/blob/main/wit/types.wit
type ErrorCode = string

// Method is supposed to be generated from WIT,
//
// WASI-HTTP uses a variant, but we use enums just to demonstrate them.
//
// https://github.com/WebAssembly/wasi-http/blob/main/wit/types.wit
type Method uint64

const (
	MethodGet Method = iota
	MethodHead
	MethodPost
	MethodPut
	MethodDelete
	MethodConnect
	MethodOptions
	MethodTrace
	MethodPatch
)

func (m Method) MarshalWube(enc wube.Encoder) error {
	return enc.WriteEnum(uint64(m))
}

func (m *Method) UnmarshalWube(dec wube.Decoder) error {
	d, err := dec.ReadEnum()
	if err != nil {
		return err
	}
	*m = Method(d)
	return nil
}

// Scheme is supposed to be generated from WIT.
//
// https://github.com/WebAssembly/wasi-http/blob/main/wit/types.wit
type Scheme struct {
	Variant SchemeVariant
}

type SchemeVariant interface {
	variantGuard_Scheme()
}

func (s Scheme) MarshalWube(enc wube.Encoder) error {
	switch v := s.Variant.(type) {
	case SchemeHttp:
		return enc.WriteVariant(0, v)

	case SchemeHttps:
		return enc.WriteVariant(1, v)

	case SchemeOther:
		return enc.WriteVariant(2, v)

	default:
		return fmt.Errorf("invalid scheme variant value: %v", v)
	}
}

func (s *Scheme) UnmarshalWube(dec wube.Decoder) error {
	// Read enum to know the discriminant.
	d, err := dec.ReadEnum()
	if err != nil {
		return err
	}

	switch d {
	case 0:
		var inner SchemeHttp
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		s.Variant = inner
		return nil

	case 1:
		var inner SchemeHttps
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		s.Variant = inner
		return nil

	case 2:
		var inner SchemeOther
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		s.Variant = inner
		return nil

	default:
		return fmt.Errorf("unknown scheme variant discriminant: %d", d)
	}
}

type SchemeHttp struct{}

func NewScheme_Http() Scheme {
	inner := types.None{}
	return Scheme{Variant: SchemeHttp(inner)}
}

func (_ SchemeHttp) variantGuard_Scheme() {}

type SchemeHttps struct{}

func NewScheme_Https() Scheme {
	inner := types.None{}
	return Scheme{Variant: SchemeHttps(inner)}
}

func (_ SchemeHttps) variantGuard_Scheme() {}

type SchemeOther types.T1[string]

func NewScheme_Other(inner types.T1[string]) Scheme {
	return Scheme{Variant: SchemeOther(inner)}
}
func (_ SchemeOther) variantGuard_Scheme() {}

/*
	  record request-options {
		connect-timeout: duration,
		first-byte-timeout: duration,
		between-bytes-timeout: duration,
	  }
*/
type RequestOptions struct {
	ConnectTimeout      Duration
	FirstByteTimeout    Duration
	BetweenBytesTimeout Duration
}

/*
	  record outgoing-request {
		headers: list<tuple<string,list<u8>>>,
		method: method,
		path-with-query: option<string>,
		scheme: option<scheme>,
		authority: option<string>,
		body: list<u8>,
		trailers: option<list<tuple<string,list<u8>>>>,
	  }
*/
type OutgoingRequest struct {
	Headers       []types.T2[string, []byte]
	Method        Method
	PathWithQuery types.Option[string]
	Scheme        types.Option[Scheme]
	Authority     types.Option[string]
	Body          []byte
	Trailers      types.Option[[]types.T2[string, []byte]]
}

/*
	  record incoming-response {
		status: status-code,
		headers: list<tuple<string,list<u8>>>,
		body: list<u8>,
		trailers: option<list<tuple<string,list<u8>>>>,
	  }
*/
type IncomingResponse struct {
	Status   StatusCode
	Headers  []types.T2[string, []byte]
	Body     []byte
	Trailers types.Option[[]types.T2[string, []byte]]
}
