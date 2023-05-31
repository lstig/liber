package types

import "github.com/swaggest/jsonschema-go"

func Uuid() jsonschema.Schema {
	def := jsonschema.Schema{}
	def.AddType(jsonschema.String)
	def.WithFormat("uuid")
	def.WithExamples("9b4e291c-4860-4b6a-aac6-cc11bd25e70f")
	return def
}