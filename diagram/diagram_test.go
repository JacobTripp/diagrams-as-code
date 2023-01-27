package diagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// this kinda represents a yaml file that defines a new diagram
var testDiagram = DiagramType{
	Name:        "Foo",
	Description: "To test that things are created as expected",
	Attributes: map[string]string{
		"splines": "ortho",
	},
}

var testNodes = map[NodeName]NodeType{
	"Init": {
		Description: "Init is a black circle representing the start.",
		Attributes: map[string]string{
			"shape": "circle",
			"color": "black",
		},
	},
	"Action": {
		Description: "Action represents one step in the activity",
		Attributes: map[string]string{
			"shape": "box",
			"color": "yellow",
		},
	},
}

var testEdges = map[EdgeName]EdgeType{
	"Default": {
		Description: "Defualt is a solid line with a vee tip",
		Attributes: map[string]string{
			"arrowhead": "vee",
			"style":     "solid",
		},
	},
}

// this section represents generated code for new diagram package
var Init = NodeType{
	Name:        "Init",
	Description: "Init is a black circle representing the start.",
	Attributes: map[string]string{
		"shape": "circle",
		"color": "black",
	},
}
var Action = NodeType{
	Name:        "Action",
	Description: "Action represents one step in the activity",
	Attributes: map[string]string{
		"shape": "box",
		"color": "yellow",
	},
}
var Default = EdgeType{
	Name:        "Default",
	Description: "Defualt is a solid line with a vee tip",
	Attributes: map[string]string{
		"arrowhead": "vee",
		"style":     "solid",
	},
}

var TestDiagram = New(
	DiagramType{
		Name:        "Foo",
		Description: "To test that things are created as expected",
		Attributes: map[string]string{
			"splines": "ortho",
		},
	},
	WithNodeTypes(Action, Init),
	WithEdgeTypes(Default),
)

// This section represents using the package
func TestAddNode(t *testing.T) {
	TestDiagram.AddNode(Init, "")
	err := TestDiagram.AddNode(Action, "Lather")
	assert.Empty(t, err)
	TestDiagram.AddNode(Action, "Rinse")
	TestDiagram.AddNode(Action, "Dry")
	err = TestDiagram.AddNode(NodeType{Name: "Bar"}, "should fail")
	assert.Error(t, err)
}

func TestAddEdge(t *testing.T) {
	TestDiagram.AddNode(Init, "")
	TestDiagram.AddNode(Action, "Foo")
	TestDiagram.AddNode(Action, "Bar")
	err := TestDiagram.AddEdge(Default, "Foo", "Bar")
	assert.Empty(t, err)
	err = TestDiagram.AddEdge(Default, "Foo", "Baz")
	assert.Error(t, err)
	err = TestDiagram.AddEdge(EdgeType{Name: "fail"}, "Foo", "Bar")
	assert.Error(t, err)
}
