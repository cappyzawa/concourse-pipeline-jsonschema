# concourse-pipeline-jsonschema
Json Schema for Concourse Pipeline

## why?
I developed [cappyzawa/fly\-lint\.vim: Vim Plugin for fly](https://github.com/cappyzawa/fly-lint.vim).  
This vim plugin executes format or lint via vim command.

I was using it, but I wanted completition as well as linter for the Concourse pipeline.

[redhat\-developer/yaml\-language\-server: Language Server for Yaml Files](https://github.com/redhat-developer/yaml-language-server) seemed to fulfill my hope.  
This tool only needs Json Schema, so I created it for Concourse Pipeline.

## Usage
[redhat\-developer/yaml\-language\-server: Language Server for Yaml Files](https://github.com/redhat-developer/yaml-language-server) should been installed.

There are [clients](https://github.com/redhat-developer/yaml-language-server#clients) for this lsp. 

This json schema needs to be linked with the concourse pipeline file by the setting of lsp client.

```
yaml.schemas: {
  "https://raw.githubusercontent.com/cappyzawa/concourse-pipeline-jsonschema/master/concourse_jsonschema.json": "/pipeline.yml"
}
```
