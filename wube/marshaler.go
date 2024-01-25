package wube

type Marshaler interface {
	MarshalWube(Encoder) error
}
