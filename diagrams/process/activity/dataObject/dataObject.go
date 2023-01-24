package dataobject

import (
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type DataObject struct {
	n node.Node
}

func New(name string) DataObject {
	return DataObject{
		n: node.New(
			name,
			node.WithAttribute(
				attribute.New(
					"shape",
					"rect",
				),
			),
		),
	}
}

func (d DataObject) Name() string {
	return d.n.Name
}

func (d DataObject) String() string {
	return d.n.String()
}

func (d DataObject) EdgeAttrs() string {
	return ""
}

func (d DataObject) GetAttrs() []attribute.Attribute {
	return d.n.Attributes
}

func (d DataObject) AddAttribute(name, value string) {
}
func (d DataObject) Node() node.Node {
	return d.n
}
