// Package cluster provides the subraph/cluster of graphviz
//
// I'm not really sure why there is a difference between graph and a subgraph/
// cluster because a cluster just looks like a nested graph to me but graphviz
// makes a distiction between them so I just make two different packages in
// case the reason becomes clear in the future.
package cluster

import (
	"fmt"
	"strings"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

// Cluster represents a subgraph that can have it's own attributes, and nodes
type Cluster struct {
	Name       string // Also is the subgraph ID
	Attributes []attribute.Attribute
	Nodes      []node.Node
	Edges      []string
}

type clusterOpt func(*Cluster)

// Optionally create clsuter with an attribute
func WithAttribute(attr attribute.Attribute) clusterOpt {
	return func(c *Cluster) {
		c.AddAttr(attr)
	}
}

// Optionally create a cluster with a node
func WithNode(node node.Node) clusterOpt {
	return func(c *Cluster) {
		c.AddNode(node)
	}
}

// If no attrs are provided then just return a defualt cluster with the
// attribute cluster set to true
func New(name string, attrs ...clusterOpt) Cluster {
	c := Cluster{
		Name: name,
		Attributes: []attribute.Attribute{
			attribute.New("cluster", true),
		},
	}
	for _, attr := range attrs {
		attr(&c)
	}
	return c
}

// helper to append to the Attributess
func (c *Cluster) AddAttr(attr attribute.Attribute) {
	c.Attributes = append(c.Attributes, attr)
}

// helper to append to the Nodes
func (c *Cluster) AddNode(node node.Node) {
	c.Nodes = append(c.Nodes, node)
}

// return a string in the DOT syntax representing a subgraph
func (c Cluster) String() string {
	var attrStr strings.Builder
	for _, attr := range c.Attributes {
		attrStr.WriteString(attr.String())
		attrStr.WriteString("\n")
	}
	var nodeStr strings.Builder
	for _, node := range c.Nodes {
		nodeStr.WriteString(node.String())
	}
	var edgeStr strings.Builder
	for _, edge := range c.Edges {
		edgeStr.WriteString(edge)
	}
	return fmt.Sprintf(
		"subgraph \"%s\"{\n%s%s%s}",
		c.Name,
		attrStr.String(),
		nodeStr.String(),
		edgeStr.String(),
	)
}

func (c *Cluster) AddEdge(edgeString string) {
	c.Edges = append(c.Edges, edgeString)
}
