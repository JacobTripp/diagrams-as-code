package graph

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/cluster"
	"github.com/JacobTripp/diagrams-as-code/dot/edge"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	graph := New()
	assert.IsType(t, Graph{}, graph)
}

func TestNewWithAttribute(t *testing.T) {
	t.Parallel()
	graph := New(WithAttribute(attribute.Attribute{}))
	assert.Len(t, graph.Attributes, 1)
}

func TestAddAttribute(t *testing.T) {
	t.Parallel()
	graph := New()
	graph.AddAttr(attribute.Attribute{})
	assert.Len(t, graph.Attributes, 1)
}

func TestNewWithCluster(t *testing.T) {
	t.Parallel()
	graph := New(WithCluster(cluster.Cluster{}))
	assert.Len(t, graph.Clusters, 1)
}

func TestAddCluster(t *testing.T) {
	t.Parallel()
	graph := New()
	graph.AddCluster(cluster.Cluster{})
	assert.Len(t, graph.Clusters, 1)
}

func TestNewWithNode(t *testing.T) {
	t.Parallel()
	graph := New(WithNode(node.Node{}))
	assert.Len(t, graph.Nodes, 1)
}

func TestAddNode(t *testing.T) {
	t.Parallel()
	graph := New()
	graph.AddNode(node.Node{})
	assert.Len(t, graph.Nodes, 1)
}

func TestNewWithEdge(t *testing.T) {
	t.Parallel()
	graph := New(WithEdge(edge.Edge{}))
	assert.Len(t, graph.Edges, 1)
}

func TestAddEdge(t *testing.T) {
	t.Parallel()
	graph := New()
	graph.AddEdge(edge.Edge{})
	assert.Len(t, graph.Edges, 1)
}

func TestString(t *testing.T) {
	t.Parallel()
	graph := New()
	expected := `digraph {
}`
	assert.Equal(t, expected, graph.String())
}

func TestStringWithAttribute(t *testing.T) {
	t.Parallel()
	attr := attribute.New("FooAttr", 0)
	graph := New(WithAttribute(attr))
	expected := `digraph {
FooAttr="0"
}`
	assert.Equal(t, expected, graph.String())
}

func TestStringWithCluster(t *testing.T) {
	t.Parallel()
	cluster := cluster.New("FooCluster")
	graph := New(WithCluster(cluster))
	expected := `digraph {
subgraph "FooCluster"{
cluster="true"
}
}`
	assert.Equal(t, expected, graph.String())
}

func TestStringWithEdge(t *testing.T) {
	t.Parallel()
	node1 := node.New("node1")
	node2 := node.New("node2")
	edge := edge.New()
	edge.From(node1)
	edge.To(node2)
	graph := New(WithEdge(edge))
	expected := `digraph {
"node1"->"node2"
}`
	assert.Equal(t, expected, graph.String())
}

func TestStringWithNode(t *testing.T) {
	t.Parallel()
	graph := New(WithNode(node.New("FooNode")))
	expected := `digraph {
"FooNode"
}`
	assert.Equal(t, expected, graph.String())
}

func TestStringWithAttrClusterEdge(t *testing.T) {
	t.Parallel()
	node1 := node.New("node1")
	node2 := node.New("node2")
	edge := edge.New()
	edge.From(node1)
	edge.To(node2)
	cluster := cluster.New("FooCluster")
	attr := attribute.New("FooAttr", 0)
	graph := New(WithAttribute(attr), WithCluster(cluster), WithEdge(edge))
	expected := `digraph {
FooAttr="0"
subgraph "FooCluster"{
cluster="true"
}
"node1"->"node2"
}`
	assert.Equal(t, expected, graph.String())
}
