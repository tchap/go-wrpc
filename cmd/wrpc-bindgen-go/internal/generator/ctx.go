package generator

import (
	"fmt"
	"io"

	"github.com/ydnar/wasm-tools-go/wit"
)

type Ctx struct {
	Interface     *wit.Interface
	Methods       map[string]*wit.Function
	Records       []*wit.TypeDef
	ImportedTypes []*wit.TypeDef
}

func NewCtx(iface *wit.Interface) *Ctx {
	c := &Ctx{Interface: iface}
	c.walk()
	return c
}

func (c *Ctx) walk() {
	for _, fnc := range c.Interface.Functions {
		for _, param := range fnc.Params {
			switch p := param.Type.(type) {
			case *wit.TypeDef:
				switch t := p.Kind.(type) {
				case *wit.TypeDef:
					switch t.Kind.(type) {
					case *wit.Record:
						c.Records = append(c.Records, t)
					}
				}
			}
		}
	}
}

func (c *Ctx) Generate(w io.Writer) error {
	// Write the records first.
	for _, r := range c.Records {
		if err := c.writeRecord(w, *r.Name, r.Kind.(*wit.Record)); err != nil {
			return err
		}
	}

	// Write collected imported types.
	for _, t := range c.ImportedTypes {
		if err := c.writeType(w, t); err != nil {
			return err
		}
	}
	return nil
}

func (c *Ctx) writeRecord(w io.Writer, name string, t *wit.Record) error {
	fmt.Fprintf(w, "type %s struct {\n", c.generatedRecordName(name))
	for _, field := range t.Fields {
		if err := c.writeRecordField(w, &field); err != nil {
			return err
		}
	}
	fmt.Fprintln(w, "}\n")
	return nil
}

func (c *Ctx) writeRecordField(w io.Writer, field *wit.Field) error {
	fmt.Fprintf(w, "\t%s ", normalizeName(field.Name))
	if err := c.writeType(w, field.Type); err != nil {
		return err
	}
	fmt.Fprintln(w)
	return nil
}

func (c *Ctx) writeVariant(w io.Writer, t *wit.TypeDef, k *wit.Variant) error {
	wrapperName := normalizeTypeRefName(*t.Name)
	fmt.Fprintf(w, "type %s struct {\n", wrapperName)
	fmt.Fprintf(w, "\tVariant %s\n", variantInterfaceName(*t.Name))
	fmt.Fprintln(w, "}\n")

	fmt.Fprintf(w, "type %s interface {\n", variantInterfaceName(*t.Name))
	fmt.Fprintf(w, "\t%s()\n", variantInterfaceMethodName(*t.Name))
	fmt.Fprintln(w, "}\n")

	// MarshalWube
	fmt.Fprintf(w, `func (wrapper %s) MarshalWube(enc wube.Encoder) error {
		switch v := wrapper.Variant.(type) {
	`, wrapperName)
	for i, kc := range k.Cases {
		fmt.Fprintf(w, `case %s:
			return enc.WriteVariant(%d, v)

		`, variantCaseTypeName(*t.Name, kc.Name), i)
	}
	fmt.Fprintf(w, `
			default:
				return fmt.Errorf("invalid %s variant value: %%v", v)
			}
		}

	`, *t.Name)

	// UnmarshalWube
	fmt.Fprintf(w, `func (wrapper *%s) UnmarshalWube(dec wube.Decoder) error {
		// Read enum to know the discriminant.
		d, err := dec.ReadEnum()
		if err != nil {
			return err
		}

		switch d {`, wrapperName)
	for i, kc := range k.Cases {
		fmt.Fprintf(w, `
			case %d:
				var inner %s
				if err := dec.Decode(&inner); err != nil {
					return err
				}

				wrapper.Variant = inner
				return nil

			`, i, variantCaseTypeName(*t.Name, kc.Name))
	}
	fmt.Fprintf(w, `
			default:
				return fmt.Errorf("unknown %s variant discriminant: %%d", d)
			}
		}
	`, *t.Name)

	for _, kase := range k.Cases {
		if err := c.writeCase(w, *t.Name, &kase); err != nil {
			return err
		}
	}
	return nil
}

func (c *Ctx) writeCase(w io.Writer, variantName string, kc *wit.Case) error {
	tName := variantCaseTypeName(variantName, kc.Name)
	fmt.Fprintf(w, "type %s ", tName)
	if err := c.writeType(w, kc.Type); err != nil {
		return err
	}
	fmt.Fprintln(w)
	fmt.Fprintf(w, "func (_ %s) %s() {}\n", tName, variantInterfaceMethodName(variantName))
	return nil
}

func (c *Ctx) generatedRecordName(name string) string {
	return fmt.Sprintf("%s", normalizeName(name))
}

func (c *Ctx) writeType(w io.Writer, t wit.Type) error {
	if t == nil {
		fmt.Fprint(w, "struct {}")
		return nil
	}

	switch t := t.(type) {
	case *wit.TypeDef:
		return c.writeUnderlyingType(w, t)

	case wit.U8:
		fmt.Fprint(w, "byte")
		return nil

	case wit.String:
		fmt.Fprint(w, "string")
		return nil

	default:
		return fmt.Errorf("unsupported WIT type: %T", t)
	}
}

func (c *Ctx) writeUnderlyingType(w io.Writer, t *wit.TypeDef) error {
	switch k := t.Kind.(type) {
	case *wit.TypeDef:
		if _, ok := k.Owner.(*wit.Interface); ok {
			fmt.Fprint(w, normalizeTypeRefName(*t.Name))
			c.ImportedTypes = append(c.ImportedTypes, k)
		}
		return nil

	case *wit.Option:
		return c.writeOption(w, k)

	case *wit.Future:
		return c.writeFuture(w, k)

	case *wit.List:
		return c.writeList(w, k)

	case *wit.Tuple:
		return c.writeTuple(w, k)

	case *wit.Variant:
		return c.writeVariant(w, t, k)

	case *wit.Stream:
		return c.writeStream(w, k)

	default:
		return fmt.Errorf("unsupported WIT type: %T", k)
	}
}

func (c *Ctx) writeOption(w io.Writer, t *wit.Option) error {
	fmt.Fprint(w, "wubetypes.Option[")
	if err := c.writeType(w, t.Type); err != nil {
		return err
	}
	fmt.Fprint(w, "]")
	return nil
}

func (c *Ctx) writeFuture(w io.Writer, t *wit.Future) error {
	fmt.Fprint(w, "wubetypes.Future[")
	if err := c.writeType(w, t.Type); err != nil {
		return err
	}
	fmt.Fprint(w, "]")
	return nil
}

func (c *Ctx) writeList(w io.Writer, t *wit.List) error {
	fmt.Fprint(w, "[]")
	return c.writeType(w, t.Type)
}

func (c *Ctx) writeTuple(w io.Writer, t *wit.Tuple) error {
	switch {
	case len(t.Types) == 0:
		fmt.Fprint(w, "wubetypes.Unit")
		return nil

	case len(t.Types) > 5:
		return fmt.Errorf("tuple size not supported: %d", len(t.Types))
	}

	fmt.Fprintf(w, "wubetypes.T%d[", len(t.Types))

	writeComma := false
	for _, t := range t.Types {
		if writeComma {
			fmt.Fprint(w, ", ")
		}
		if err := c.writeType(w, t); err != nil {
			return err
		}
		writeComma = true
	}
	fmt.Fprint(w, "]")
	return nil
}

func (c *Ctx) writeStream(w io.Writer, t *wit.Stream) error {
	// stream<type> is represented as future<list<type>>
	fmt.Fprint(w, "wubetypes.Future[[]")
	if err := c.writeType(w, t.Element); err != nil {
		return err
	}
	fmt.Fprint(w, "]")
	return nil
}
