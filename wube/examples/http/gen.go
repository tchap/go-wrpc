package http

// This file was basically generated using wrpc-bindgen-go.

import (
	"fmt"

	"github.com/tchap/go-wrpc/wube"
	wubetypes "github.com/tchap/go-wrpc/wube/types"
)

type IncomingRequest struct {
	Method        Method
	PathWithQuery wubetypes.Option[string]
	Scheme        wubetypes.Option[Scheme]
	Authority     wubetypes.Option[string]
	Headers       []wubetypes.T2[string, []byte]
	Body          wubetypes.Future[[]byte]
	Trailers      wubetypes.Future[[]wubetypes.T2[string, []byte]]
}

type Method struct {
	Variant variantValue_Method
}

type variantValue_Method interface {
	isVariant_Method()
}

func (wrapper Method) MarshalWube(enc wube.Encoder) error {
	switch v := wrapper.Variant.(type) {
	case MethodGet:
		return enc.WriteVariant(0, v)

	case MethodHead:
		return enc.WriteVariant(1, v)

	case MethodPost:
		return enc.WriteVariant(2, v)

	case MethodPut:
		return enc.WriteVariant(3, v)

	case MethodDelete:
		return enc.WriteVariant(4, v)

	case MethodConnect:
		return enc.WriteVariant(5, v)

	case MethodOptions:
		return enc.WriteVariant(6, v)

	case MethodTrace:
		return enc.WriteVariant(7, v)

	case MethodPatch:
		return enc.WriteVariant(8, v)

	case MethodOther:
		return enc.WriteVariant(9, v)

	default:
		return fmt.Errorf("invalid method variant value: %v", v)
	}
}

func (wrapper *Method) UnmarshalWube(dec wube.Decoder) error {
	// Read enum to know the discriminant.
	d, err := dec.ReadEnum()
	if err != nil {
		return err
	}

	switch d {
	case 0:
		var inner MethodGet
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 1:
		var inner MethodHead
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 2:
		var inner MethodPost
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 3:
		var inner MethodPut
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 4:
		var inner MethodDelete
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 5:
		var inner MethodConnect
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 6:
		var inner MethodOptions
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 7:
		var inner MethodTrace
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 8:
		var inner MethodPatch
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 9:
		var inner MethodOther
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	default:
		return fmt.Errorf("unknown method variant discriminant: %d", d)
	}
}

type MethodGet struct{}

func (_ MethodGet) isVariant_Method() {}

type MethodHead struct{}

func (_ MethodHead) isVariant_Method() {}

type MethodPost struct{}

func (_ MethodPost) isVariant_Method() {}

type MethodPut struct{}

func (_ MethodPut) isVariant_Method() {}

type MethodDelete struct{}

func (_ MethodDelete) isVariant_Method() {}

type MethodConnect struct{}

func (_ MethodConnect) isVariant_Method() {}

type MethodOptions struct{}

func (_ MethodOptions) isVariant_Method() {}

type MethodTrace struct{}

func (_ MethodTrace) isVariant_Method() {}

type MethodPatch struct{}

func (_ MethodPatch) isVariant_Method() {}

type MethodOther string

func (_ MethodOther) isVariant_Method() {}

type Scheme struct {
	Variant variantValue_Scheme
}

type variantValue_Scheme interface {
	isVariant_Scheme()
}

func (wrapper Scheme) MarshalWube(enc wube.Encoder) error {
	switch v := wrapper.Variant.(type) {
	case SchemeHTTP:
		return enc.WriteVariant(0, v)

	case SchemeHTTPS:
		return enc.WriteVariant(1, v)

	case SchemeOther:
		return enc.WriteVariant(2, v)

	default:
		return fmt.Errorf("invalid scheme variant value: %v", v)
	}
}

func (wrapper *Scheme) UnmarshalWube(dec wube.Decoder) error {
	// Read enum to know the discriminant.
	d, err := dec.ReadEnum()
	if err != nil {
		return err
	}

	switch d {
	case 0:
		var inner SchemeHTTP
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 1:
		var inner SchemeHTTPS
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	case 2:
		var inner SchemeOther
		if err := dec.Decode(&inner); err != nil {
			return err
		}

		wrapper.Variant = inner
		return nil

	default:
		return fmt.Errorf("unknown scheme variant discriminant: %d", d)
	}
}

type SchemeHTTP struct{}

func (_ SchemeHTTP) isVariant_Scheme() {}

type SchemeHTTPS struct{}

func (_ SchemeHTTPS) isVariant_Scheme() {}

type SchemeOther string

func (_ SchemeOther) isVariant_Scheme() {}
