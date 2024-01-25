package wube

type WriteBuffer interface {
	Write([]byte) (int, error)
	WriteByte(byte) error
}
