package path

import (
	"fmt"
	"testing"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
	"github.com/stretchr/testify/assert"
)

type nodeMock struct {
	name string
}

func (n nodeMock) String() string { return fmt.Sprintf("node %s\n", n.name) }

func (n nodeMock) Name() string { return n.name }
func (n nodeMock) GetAttrs() []attribute.Attribute {
	return []attribute.Attribute{}
}
func (n nodeMock) AddAttribute(string, string) {}
func (n nodeMock) Node() node.Node             { return node.New("mock") }

func TestNew(t *testing.T) {
	p := New()
	assert.IsType(t, Path{}, p)
}

func TestAddInit(t *testing.T) {
	p := New()
	err := p.AddInit(nodeMock{})
	assert.Empty(t, err)
	assert.Equal(t, 1, p.nodes.Len())
	err = p.AddInit(nodeMock{})
	assert.Error(t, err)
}

func TestAddNode(t *testing.T) {
	p := New()
	p.AddNode(nodeMock{})
	assert.Equal(t, 1, p.nodes.Len())
}

func TestLen(t *testing.T) {
	p := New()
	assert.Equal(t, 0, p.Len())
	p.AddNode(nodeMock{})
	assert.Equal(t, 1, p.Len())
}

func TestString(t *testing.T) {
	p := New()
	p.AddNode(nodeMock{name: "one"})
	p.AddNode(nodeMock{name: "two"})
	p.AddNode(nodeMock{name: "three"})
	expected := "\"one\"->\"two\"->\"three\""
	assert.Equal(t, expected, p.PathString())
}

func TestNodeString(t *testing.T) {
	p := New()
	p.AddNode(nodeMock{name: "one"})
	p.AddNode(nodeMock{name: "two"})
	p.AddNode(nodeMock{name: "three"})
	expected := `node one
node two
node three
`
	assert.Equal(t, expected, p.NodeString())
}
