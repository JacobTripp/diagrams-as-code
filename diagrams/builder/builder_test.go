package builder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDiagramSpec = DiagramSpec{
	Name:        "Foo Diagram",
	Description: "This diagram shows testing",
	Attributes: map[string]string{
		"color": "blue",
	},
	Template: `digraph "Foo Diagram" {
{{.nodes}}
{{.edges}}
}`,
}
var testNodeSpec = NodeSpec{
	Name:        "foo node",
	Description: "node used for tests.",
	Attributes: map[string]string{
		"bar1": "baz1",
		"bar2": "baz2",
		"bar3": "baz3",
	},
	Template: "\"{{.name}}\" {{.attributes}}",
}

var testEdgeSpec = EdgeSpec{
	Name:        "foo edge",
	Description: "edge used for tests.",
	Attributes: map[string]string{
		"bar1": "baz1",
		"bar2": "baz2",
		"bar3": "baz3",
	},
	Template: "\"{{.from}}\"->\"{{.to}}\" {{.attributes}}",
}

func TestGenerate(t *testing.T) {
	b := testDiagramSpec
	b.AddNodeSpec(testNodeSpec)
	b.AddEdgeSpec(testEdgeSpec)
	buffer := bytes.NewBuffer([]byte{})
	b.Generate(buffer)
	assert.Contains(t, buffer.String(), "String() string {")
}

func TestAddEdgeSpec(t *testing.T) {
	b := testDiagramSpec
	b.AddEdgeSpec(testEdgeSpec)
	assert.Len(t, b.Edges, 1)
}

func TestAddNodeSpec(t *testing.T) {
	b := testDiagramSpec
	b.AddNodeSpec(testNodeSpec)
	assert.Len(t, b.Nodes, 1)
}

func TestDiagramSpec(t *testing.T) {
	b := testDiagramSpec
	assert.Equal(t, "Foo Diagram", b.Name)
}
