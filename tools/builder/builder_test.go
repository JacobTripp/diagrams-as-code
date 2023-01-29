package builder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var diagramYaml = []byte(`
diagram:
  name: Activity
  description: >-
    Activity diagrams show steps that a Use Case needs to
    accomplish it's purpose.
  attributes:
    color: green
    cluster: true
nodes:
  - name: Action
    description: One task the activity needs to accomplish
    attributes:
      shape: box
      style: rounded
  - name: Init
    description: Is the most basic way to start an Activity
    attributes:
      shape: circle
      color: black
edges:
  - name: Default
    description: Activity diagrams only have one type of connector
    attributes:
      arrowhead: vee
`)

func TestImportYaml(t *testing.T) {
	yd, err := ImportYaml(diagramYaml)
	assert.Empty(t, err)
	assert.NotEmpty(t, yd)
	assert.Equal(t, "circle", yd.Nodes[1].Attributes["shape"])
	assert.Equal(t, "Activity", yd.Diagram.DiagramName)
}

func TestWritePackage(t *testing.T) {
	yd, _ := ImportYaml(diagramYaml)
	var out = bytes.NewBuffer([]byte{})
	WritePackage(out, yd)
	assert.Contains(t, out.String(), "diagram.New(")
}
