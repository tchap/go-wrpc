package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/ydnar/wasm-tools-go/wit"

	"github.com/tchap/go-wrpc/cmd/wrpc-bindgen-go/internal/generator"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()
	if len(flag.Args()) < 3 {
		return errors.New("at least 3 arguments required: <wit-json-file> <wit-package-name> <wit-interface> ...")
	}

	jsonFilename := flag.Arg(0)
	pkgName, err := wit.ParsePackageName(flag.Arg(1))
	if err != nil {
		return err
	}
	ifaceNames := flag.Args()[2:]

	f, err := os.Open(jsonFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	res, err := wit.DecodeJSON(f)
	if err != nil {
		return err
	}

	for _, pkg := range res.Packages {
		if pkg.Name.String() == pkgName.String() {
			return generator.New(pkg).GenerateInterfaces(ifaceNames...)
		}
	}
	return errors.New("package not found: " + pkgName.String())
}
