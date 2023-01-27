package merge

import (
	"github.com/JacobTripp/diagrams-as-code/process/activity/decision"
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Merge struct {
	n node.Node
	d *decision.Decision
}

func New() Merge {
	m := Merge{
		n: node.New("Merge"),
	}
	m.AddAttribute("shape", "diamond")
	m.AddAttribute("orientation", "90")
	return m
}

func (m *Merge) MergeDecision(d *decision.Decision) {
	for _, path := range d.Paths {
		path.AddNode(m)
	}
	m.d = d
}

func (m Merge) String() string {
	return m.n.String()
}

func (m Merge) Name() string {
	return m.n.Name
}

func (m Merge) GetAttrs() []attribute.Attribute {
	return m.n.Attributes
}

func (m *Merge) AddAttribute(name, value string) {
	m.n.AddAttr(
		attribute.New(
			name,
			value,
		),
	)
}
func (m Merge) Node() node.Node {
	return m.n
}
