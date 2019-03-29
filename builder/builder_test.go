package builder

import (
	"testing"

	"github.com/rancher/mapper"
	"github.com/stretchr/testify/assert"
)

func TestEmptyStringWithDefault(t *testing.T) {
	schema := &mapper.Schema{
		ResourceFields: map[string]mapper.Field{
			"foo": {
				Default: "foo",
				Type:    "string",
				Create:  true,
			},
		},
	}
	schemas := mapper.NewSchemas()
	schemas.AddSchema(*schema)

	builder := NewBuilder(schemas)

	// Test if no field we set to "foo"
	result, err := builder.Construct(schema, nil, Create)
	if err != nil {
		t.Fatal(err)
	}
	value, ok := result["foo"]
	assert.True(t, ok)
	assert.Equal(t, "foo", value)

	// Test if field is "" we set to "foo"
	result, err = builder.Construct(schema, map[string]interface{}{
		"foo": "",
	}, Create)
	if err != nil {
		t.Fatal(err)
	}
	value, ok = result["foo"]
	assert.True(t, ok)
	assert.Equal(t, "foo", value)
}
