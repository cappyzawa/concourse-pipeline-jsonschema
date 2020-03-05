package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/jsonschema"
	"github.com/concourse/concourse/atc"
)

func main() {
	var pipelineConfig atc.Config
	reflected, err := jsonschema.Reflect(&pipelineConfig).MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, string(reflected))
}
