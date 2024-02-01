package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"

	"github.com/ydnar/wasm-tools-go/wit"
)

type Generator struct {
	pkg *wit.Package
}

func New(pkg *wit.Package) *Generator {
	return &Generator{pkg: pkg}
}

func (gen *Generator) GenerateInterfaces(ifaceNames ...string) error {
	// Gather interfaces to process.
	ifaces := make([]*wit.Interface, 0, len(ifaceNames))
	for _, ifaceName := range ifaceNames {
		iface, ok := gen.pkg.Interfaces[ifaceName]
		if !ok {
			return fmt.Errorf("interface not found: %s", ifaceName)
		}
		ifaces = append(ifaces, iface)
	}

	// Generate.
	var b bytes.Buffer
	if err := gen.generate(&b, ifaces...); err != nil {
		return err
	}

	out, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(out)
	//_, err := os.Stdout.Write(b.Bytes())
	return err
}

func (gen *Generator) generate(w io.Writer, ifaces ...*wit.Interface) error {
	for _, iface := range ifaces {
		if err := NewCtx(iface).Generate(w); err != nil {
			return err
		}
	}
	return nil
}
