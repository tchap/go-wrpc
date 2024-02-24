package main

import (
	"fmt"
	"os"

	"github.com/kr/pretty"

	"github.com/tchap/go-wrpc/internal/hostdata"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Load host data.
	hostData, err := hostdata.Read(os.Stdin)
	if err != nil {
		return err
	}

	pretty.Println(hostData)
	return nil
}
