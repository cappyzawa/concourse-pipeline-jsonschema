package main

import (
	"fmt"
	"os"

	"github.com/concourse/concourse/atc"
	"github.com/invopop/jsonschema"
)

func main() {
	var pipelineConfig atc.Config
	schema := jsonschema.Reflect(&pipelineConfig)
	// stepSchema(schema)
	reflected, err := schema.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, string(reflected))
}
