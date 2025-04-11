// In denkproto-go/json/inference_graph_recipe/schema.go
package inference_graph_recipe

import _ "embed" // Import the embed package

//go:embed inference_graph_recipe.schema.json
var schemaBytes []byte

//go:embed inference_graph_recipe.schema.json
var schemaString string

// GetSchemaBytes provides the content of inference_graph_recipe.schema.json as bytes.
func GetSchemaBytes() []byte {
	// Return a copy to prevent modification of the embedded data
	cpy := make([]byte, len(schemaBytes))
	copy(cpy, schemaBytes)
	return cpy
}

// GetSchemaString provides the content of inference_graph_recipe.schema.json as a string.
func GetSchemaString() string {
	return schemaString
}
