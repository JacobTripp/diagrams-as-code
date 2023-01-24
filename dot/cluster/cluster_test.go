package cluster

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
	"github.com/stretchr/testify/assert"
)

// New()
type newInput struct {
	attrs []attribute.Attribute
	nodes []node.Node
}
type newExpected struct {
	attrs int
	nodes int
}

var newTestCases = map[string]struct {
	input    newInput
	expected newExpected
}{
	"Default attribute Zero nodes": {
		input: newInput{
			attrs: []attribute.Attribute{},
			nodes: []node.Node{},
		},
		expected: newExpected{
			attrs: 1,
			nodes: 0,
		},
	},
	"One more attribute Zero nodes": {
		input: newInput{
			attrs: []attribute.Attribute{
				attribute.Attribute{},
			},
			nodes: []node.Node{},
		},
		expected: newExpected{
			attrs: 2,
			nodes: 0,
		},
	},
	"Two more attributes Zero nodes": {
		input: newInput{
			attrs: []attribute.Attribute{
				attribute.Attribute{},
				attribute.Attribute{},
			},
			nodes: []node.Node{},
		},
		expected: newExpected{
			attrs: 3,
			nodes: 0,
		},
	},
	"Default attribute One node": {
		input: newInput{
			attrs: []attribute.Attribute{},
			nodes: []node.Node{
				node.Node{},
			},
		},
		expected: newExpected{
			attrs: 1,
			nodes: 1,
		},
	},
	"Default attribute Two nodes": {
		input: newInput{
			attrs: []attribute.Attribute{},
			nodes: []node.Node{
				node.Node{},
				node.Node{},
			},
		},
		expected: newExpected{
			attrs: 1,
			nodes: 2,
		},
	},
	"Two more attributes Two nodes": {
		input: newInput{
			attrs: []attribute.Attribute{
				attribute.Attribute{},
				attribute.Attribute{},
			},
			nodes: []node.Node{
				node.Node{},
				node.Node{},
			},
		},
		expected: newExpected{
			attrs: 3,
			nodes: 2,
		},
	},
}

// cluster.AddAttr() cluster.AddNode()
func TestNewAdd(t *testing.T) {
	for testName, testCase := range newTestCases {
		t.Run(testName, func(t *testing.T) {
			c := New("Foo")
			for _, attr := range testCase.input.attrs {
				c.AddAttr(attr)
			}
			for _, node := range testCase.input.nodes {
				c.AddNode(node)
			}
			assert.IsType(t, Cluster{}, c)
			assert.Len(t, c.Attributes, testCase.expected.attrs)
			assert.Len(t, c.Nodes, testCase.expected.nodes)
		})
	}
}

// WithAttribute()
func TestNewWithAttribute(t *testing.T) {
	for testName, testCase := range newTestCases {
		t.Run(testName, func(t *testing.T) {
			withAttrs := []clusterOpt{}
			for _, attr := range testCase.input.attrs {
				withAttrs = append(withAttrs, WithAttribute(attr))
			}
			c := New("Foo", withAttrs...)
			assert.IsType(t, Cluster{}, c)
			assert.Len(t, c.Attributes, testCase.expected.attrs)
		})
	}
}

// WithNode()
func TestNewWithNode(t *testing.T) {
	for testName, testCase := range newTestCases {
		t.Run(testName, func(t *testing.T) {
			withNodes := []clusterOpt{}
			for _, node := range testCase.input.nodes {
				withNodes = append(withNodes, WithNode(node))
			}
			c := New("Foo", withNodes...)
			assert.IsType(t, Cluster{}, c)
			assert.Len(t, c.Nodes, testCase.expected.nodes)
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()
	cluster := New("Foo")
	expected := `subgraph "Foo"{
cluster="true"
}`
	assert.Equal(t, expected, cluster.String())
}

func TestStringWithNode(t *testing.T) {
	t.Parallel()
	cluster := New("Foo", WithNode(node.New("Bar")))
	expected := `subgraph "Foo"{
cluster="true"
"Bar"
}`
	assert.Equal(t, expected, cluster.String())
}

func TestAddEdge(t *testing.T) {
	c := New("foo")
	c.AddEdge("bar->baz")
	assert.Len(t, c.Edges, 1)
}
