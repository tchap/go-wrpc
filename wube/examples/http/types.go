package http

import (
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
