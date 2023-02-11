package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	graph := NewGraph("Foo")
	assert.IsType(t, &Graph{}, graph)
}

func TestAddPoint(t *testing.T) {
	graph := NewGraph("Foo")
	graph.AddPoint(NewPoint("point"))
}

func TestAddLine(t *testing.T) {
	graph := NewGraph("Foo")
	graph.AddLine(NewLine("line"))
}

func TestGraphAddAttribute(t *testing.T) {
	graph := NewGraph("Foo")
	graph.AddAttribute("cluster", "true")
}
