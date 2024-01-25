package wube

type Unmarshaler interface {
	UnmarshalWube(Decoder) error
}
