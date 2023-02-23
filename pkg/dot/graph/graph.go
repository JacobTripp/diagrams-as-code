package graph

import (
	"fmt"
	"strings"

	"github.com/JacobTripp/diagrams-as-code/pkg/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/pkg/dot/cluster"
	"github.com/JacobTripp/diagrams-as-code/pkg/dot/edge"
	"github.com/JacobTripp/diagrams-as-code/pkg/dot/node"
)

type Graph struct {
	Attributes []attribute.Attribute
	Clusters   []cluster.Cluster
	Nodes      []node.Node
	Edges      []edge.Edge
}

type GraphOption func(g *Graph)

func WithAttribute(attr attribute.Attribute) GraphOption {
	return func(g *Graph) {
		g.AddAttr(attr)
	}
}

func WithCluster(cl cluster.Cluster) GraphOption {
	return func(g *Graph) {
		g.AddCluster(cl)
	}
}

func WithEdge(ed edge.Edge) GraphOption {
	return func(g *Graph) {
		g.AddEdge(ed)
	}
}

func WithNode(n node.Node) GraphOption {
	return func(g *Graph) {
		g.AddNode(n)
	}
}

func New(opts ...GraphOption) Graph {
	g := Graph{}
	for _, opt := range opts {
		opt(&g)
	}
	return g
}

func (g *Graph) AddEdge(ed edge.Edge) {
	g.Edges = append(g.Edges, ed)
}

func (g *Graph) AddAttr(attr attribute.Attribute) {
	g.Attributes = append(g.Attributes, attr)
}

func (g *Graph) AddCluster(cl cluster.Cluster) {
	g.Clusters = append(g.Clusters, cl)
}

func (g *Graph) AddNode(n node.Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g Graph) String() string {
	var attrStr strings.Builder
	for _, attr := range g.Attributes {
		attrStr.WriteString(attr.String())
		attrStr.WriteString("\n")
	}
	var clusterStr strings.Builder
	for _, cluster := range g.Clusters {
		clusterStr.WriteString(cluster.String())
		clusterStr.WriteString("\n")
	}
	var nodeStr strings.Builder
	for _, n := range g.Nodes {
		nodeStr.WriteString(n.String())
	}
	var edgeStr strings.Builder
	for _, edge := range g.Edges {
		edgeStr.WriteString(edge.String())
	}
	return fmt.Sprintf(
		"digraph {\n%s%s%s%s}",
		attrStr.String(),
		nodeStr.String(),
		clusterStr.String(),
		edgeStr.String(),
	)
}
