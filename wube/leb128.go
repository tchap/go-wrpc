package wube

import (
	"io"

	"github.com/go-delve/delve/pkg/dwarf/leb128"
)

type leb128Writer struct {
	next io.ByteWriter
	err  error
}

func newLEB128Writer(next io.ByteWriter) leb128Writer {
	return leb128Writer{next: next}
}

func (w *leb128Writer) WriteByte(c byte) error {
	err := w.next.WriteByte(c)
	if err != nil && w.err == nil {
		w.err = err
	}
	return err
}

func (w *leb128Writer) Err() error {
	return w.err
}

type leb128Reader struct {
	next leb128.Reader
	err  error
}

func newLEB128Reader(next leb128.Reader) leb128Reader {
	return leb128Reader{next: next}
}

func (r *leb128Reader) Read(p []byte) (int, error) {
	n, err := r.next.Read(p)
	if err != nil && r.err == nil {
		r.err = err
	}
	return n, err
}

func (r *leb128Reader) ReadByte() (byte, error) {
	c, err := r.next.ReadByte()
	if err != nil && r.err == nil {
		r.err = err
	}
	return c, err
}

func (r *leb128Reader) Len() int {
	return r.next.Len()
}

func (r *leb128Reader) Err() error {
	return r.err
}
