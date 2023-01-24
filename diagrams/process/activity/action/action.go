package action

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Action struct {
	n node.Node
}

type actionOption func(*Action)

func IsActivity() actionOption {
	return func(a *Action) {
		activitySymbol := "&#8916;"
		a.n.AddAttr(
			attribute.New(
				"label",
				fmt.Sprintf("%s %s", activitySymbol, a.n.Name),
			),
		)
	}
}

func New(name string, opts ...actionOption) Action {
	a := Action{
		n: node.New(
			name,
			node.WithAttribute(
				attribute.New(
					"shape",
					"rect",
				),
			),
			node.WithAttribute(
				attribute.New(
					"style",
					"rounded",
				),
			),
		),
	}
	for _, opt := range opts {
		opt(&a)
	}
	return a
}

func (a Action) Name() string {
	return a.n.Name
}
func (a Action) String() string {
	return a.n.String()
}

func (a Action) EdgeAttrs() string { return "" }
func (a Action) GetAttrs() []attribute.Attribute {
	return a.n.Attributes
}
func (a *Action) AddAttribute(name, value string) {
	a.n.AddAttr(attribute.New(name, value))
}

func (a Action) Node() node.Node {
	return a.n
}
