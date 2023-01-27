package initnode

import (
	"github.com/JacobTripp/diagrams-as-code/diagram/path"
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type InitNode struct {
	n node.Node
}

type InitType string

const (
	Default      = InitType("default")
	TimeEvent    = InitType("time event")
	DataObject   = InitType("data object")
	RecieveEvent = InitType("recieve event")
)

var shapes = map[InitType][]attribute.Attribute{
	Default: {
		attribute.New(
			"shape",
			"circle",
		),
		attribute.New(
			"label",
			"",
		),
		attribute.New(
			"fillcolor",
			"black",
		),
		attribute.New(
			"width",
			"0.5",
		),
	},
}

type InitNodeOption func(*InitNode)

func UseNode(n path.Node) InitNodeOption {
	return func(init *InitNode) {
		init.n = node.New(
			n.Name(),
			node.WithAttributes(n.GetAttrs()),
		)
	}
}

func New(t InitType, opts ...InitNodeOption) InitNode {
	node := InitNode{
		n: node.New(
			"init",
			node.WithAttributes(shapes[t]),
		),
	}
	if t != Default {
		for _, opt := range opts {
			opt(&node)
		}
	}
	return node
}

func (init InitNode) Name() string {
	return init.n.Name
}
func (init InitNode) String() string {
	return init.n.String()
}
func (init InitNode) EdgeAttrs() string {
	return ""
}
func (init InitNode) GetAttrs() []attribute.Attribute {
	return init.n.Attributes
}
func (init InitNode) AddAttribute(name, value string) {
}
func (init InitNode) Node() node.Node {
	return init.n
}
