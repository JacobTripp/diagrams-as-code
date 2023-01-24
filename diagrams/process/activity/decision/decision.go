package decision

import (
	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Decision struct {
	n     node.Node
	Paths map[string]path.Path
}

func New() Decision {
	return Decision{
		n: node.New(
			"decision",
			node.WithAttribute(
				attribute.New(
					"label",
					"",
				),
			),
			node.WithAttribute(
				attribute.New(
					"shape",
					"utr",
				),
			),
			node.WithAttribute(
				attribute.New(
					"orientation",
					"90",
				),
			),
		),
		Paths: make(map[string]path.Path),
	}
}

func (d *Decision) AddAction(condition string, n path.Node) {
	p := d.Paths[condition]
	p.AddNode(n)
	d.Paths[condition] = p
}

func (d *Decision) AddPath(condition string) {
	p := path.New()
	p.AddInit(d)
	d.Paths[condition] = p
}

// path.Node interface implementations
func (d Decision) String() string {
	return d.n.String()
}
func (d Decision) Name() string {
	return d.n.Name
}
func (d Decision) GetAttrs() []attribute.Attribute {
	return d.n.Attributes
}
func (d Decision) AddAttribute(name, value string) {
}
func (d Decision) Node() node.Node {
	return d.n
}
