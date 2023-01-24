package edge

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	edge := New()
	assert.IsType(t, Edge{}, edge)
	assert.Equal(t, "->", edge.EdgeOp)
}

func TestNewWithAttributes(t *testing.T) {
	t.Parallel()
	edge := New(WithAttribute(attribute.Attribute{}))
	assert.Len(t, edge.Attributes, 1)
}

func TestAddAttributeToEdge(t *testing.T) {
	t.Parallel()
	edge := New()
	edge.AddAttr(attribute.Attribute{})
	assert.Len(t, edge.Attributes, 1)
}

func TestNewDirected(t *testing.T) {
	t.Parallel()
	edge := New(WithUnDirected())
	assert.Equal(t, "--", edge.EdgeOp)
}

func TestEdgeFrom(t *testing.T) {
	t.Parallel()
	edge := New()
	edge.From(node.New("Foo"))
	assert.NotEmpty(t, edge.Origin)
}

func TestEdgeTo(t *testing.T) {
	t.Parallel()
	edge := New()
	edge.To(node.New("Foo"))
	assert.NotEmpty(t, edge.Destination)
}

func TestString(t *testing.T) {
	t.Parallel()
	edge := New()
	edge.From(node.New("Foo"))
	edge.To(node.New("Bar"))
	assert.Equal(t, "\"Foo\"->\"Bar\"\n", edge.String())
}

func TestStringWithAttributes(t *testing.T) {
	t.Parallel()
	edge := New(WithAttribute(attribute.New("Attr", 0)))
	edge.From(node.New("Foo"))
	edge.To(node.New("Bar"))
	assert.Equal(t, "\"Foo\"->\"Bar\" [Attr=\"0\"]\n", edge.String())
}
