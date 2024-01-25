package wube

type ReadBuffer interface {
	Read([]byte) (int, error)
	ReadByte() (byte, error)
	Len() int
}
