// Package edge makes connections between nodes
//
package edge

import (
	"fmt"
	"strings"

	"github.com/JacobTripp/diagrams-as-code/pkg/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/pkg/dot/node"
)

type Edge struct {
	EdgeOp      string
	Origin      node.Node
	Destination node.Node
	Attributes  []attribute.Attribute
}

type EdgeOpt func(*Edge)

func WithUnDirected() EdgeOpt {
	return func(e *Edge) {
		e.EdgeOp = "--"
	}
}

func WithAttribute(attr attribute.Attribute) EdgeOpt {
	return func(e *Edge) {
		e.Attributes = append(e.Attributes, attr)
	}
}

func New(opts ...EdgeOpt) Edge {
	e := Edge{
		EdgeOp: "->",
	}
	for _, opt := range opts {
		opt(&e)
	}
	return e
}

func (e *Edge) From(n node.Node) {
	e.Origin = n
}

func (e *Edge) To(n node.Node) {
	e.Destination = n
}

func (e *Edge) AddAttr(attr attribute.Attribute) {
	e.Attributes = append(e.Attributes, attr)
}

func (e Edge) String() string {
	var attrStr strings.Builder
	if len(e.Attributes) > 0 {
		attrStr.WriteString(" [")
		for _, attr := range e.Attributes {
			attrStr.WriteString(attr.String())
		}
		attrStr.WriteString("]")
	}
	return fmt.Sprintf(
		"\"%s\"%s\"%s\"%s\n",
		e.Origin.Name,
		e.EdgeOp,
		e.Destination.Name,
		attrStr.String(),
	)
}
