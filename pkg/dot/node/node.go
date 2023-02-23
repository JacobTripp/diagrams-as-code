// Package node provides graphviz node functionality.
//
// This package allows creation of node with or without attributes at
// creation time or allows attributes to be added after a node has been
// created. A node can be returned as a string.
package node

import (
	"fmt"
	"strings"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
)

// The representation of a node where the Name becomes the node ID.
type Node struct {
	Name       string
	Attributes []attribute.Attribute
}

type optionalAttr func(n *Node)

// To add an optional attribute to create a node with
func WithAttribute(a attribute.Attribute) optionalAttr {
	return func(n *Node) {
		n.AddAttr(a)
	}
}

func WithAttributes(attrs []attribute.Attribute) optionalAttr {
	return func(n *Node) {
		for _, attr := range attrs {
			n.AddAttr(attr)
		}
	}
}

// Create new attribute with name as it's ID, optially create with attributes
func New(name string, attrs ...optionalAttr) Node {
	n := Node{
		Name: name,
	}
	for _, attr := range attrs {
		attr(&n)
	}
	return n
}

// Appends the attribute the node's attribute list by mutation (side effect)
// of the node.
func (n *Node) AddAttr(attr attribute.Attribute) {
	n.Attributes = append(n.Attributes, attr)
}

func (n *Node) AddAttrs(attrs []attribute.Attribute) {
	for _, attr := range attrs {
		n.Attributes = append(n.Attributes, attr)
	}
}

// returns a string in the dot language syntax
func (n *Node) String() string {
	var attrStr strings.Builder
	if len(n.Attributes) > 0 {
		attrStr.WriteString(" [")
		for _, attr := range n.Attributes {
			attrStr.WriteString(attr.String())
		}
		attrStr.WriteString("]")
	}
	return fmt.Sprintf("\"%s\"%s\n", n.Name, attrStr.String())
}
