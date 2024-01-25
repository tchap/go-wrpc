package wube

var _ WriteBuffer = NewSizeBuffer()

type SizeBuffer struct {
	n int
}

func NewSizeBuffer() *SizeBuffer {
	return &SizeBuffer{}
}

func (enc *SizeBuffer) Size() int {
	return enc.n
}

func (enc *SizeBuffer) Write(bs []byte) (int, error) {
	enc.n += len(bs)
	return len(bs), nil
}

func (enc *SizeBuffer) WriteByte(v byte) error {
	enc.n += 1
	return nil
}
