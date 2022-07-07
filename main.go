package main

import (
	"fmt"
	"os"

	"github.com/concourse/concourse/atc"
	"github.com/invopop/jsonschema"
)

func stepSchema(schema *jsonschema.Schema) {
	stepDef := schema.Definitions["Step"]

	// thx: https://gist.github.com/4662752c4af75b3220d94b657683d090
	// https://concourse-ci.org/steps.html#schema.step
	stepConfigs := []atc.StepConfig{
		&atc.GetStep{},
		&atc.PutStep{},
		&atc.TaskStep{},
		&atc.SetPipelineStep{},
		&atc.LoadVarStep{},
		&atc.InParallelStep{},
		&atc.DoStep{},
		&atc.TryStep{},
	}

	for _, s := range stepConfigs {
		stepSchema := jsonschema.Reflect(s)
		stepDef.AnyOf = append(stepDef.AnyOf, stepSchema.ContentSchema)
		for subName, subDef := range stepSchema.Definitions {
			if _, present := schema.Definitions[subName]; !present {
				schema.Definitions[subName] = subDef
			}
		}
	}
	stepDef.Required = nil
	schema.Definitions["Step"].Properties = nil
	schema.Definitions["Step"].Type = ""
	schema.Definitions["Step"].AdditionalProperties = nil
}

func main() {
	var pipelineConfig atc.Config
	schema := jsonschema.Reflect(&pipelineConfig)
	stepSchema(schema)
	schema.AdditionalProperties = jsonschema.TrueSchema
	schema.Definitions["CheckEvery"].Type = "string"
	reflected, err := schema.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, string(reflected))
}
