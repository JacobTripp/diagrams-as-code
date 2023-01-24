package finalnode

import (
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type FinalNode struct {
	n node.Node
}

type endType string

const (
	Default = endType("default")
	PathEnd = endType("path end")
)

var style = map[endType][]attribute.Attribute{
	Default: {
		attribute.New(
			"shape",
			"doublecircle",
		),
		attribute.New(
			"color",
			"black",
		),
		attribute.New(
			"label",
			"",
		),
	},
	PathEnd: {
		attribute.New(
			"shape",
			"circle",
		),
		attribute.New(
			"color",
			"white",
		),
		attribute.New(
			"label",
			"",
		),
	},
}

func New(t endType) FinalNode {
	n := node.New("FinalNode")
	n.AddAttrs(style[t])
	return FinalNode{
		n: n,
	}
}

func (f FinalNode) String() string {
	return f.n.String()
}

func (f FinalNode) Name() string {
	return f.n.Name
}

func (f FinalNode) GetAttrs() []attribute.Attribute {
	return f.n.Attributes
}

func (f *FinalNode) AddAttribute(name, value string) {
	f.n.AddAttr(
		attribute.New(name, value),
	)
}
func (f FinalNode) Node() node.Node {
	return f.n
}
