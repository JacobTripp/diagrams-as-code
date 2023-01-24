package signal

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Signal struct {
	n     node.Node
	attrs []attribute.Attribute
}

type SignalType string

const (
	Recieve = SignalType("recieve")
	Send    = SignalType("send")
)

var shape = map[SignalType][]attribute.Attribute{
	Recieve: {
		attribute.New(
			"shape",
			"cds",
		),
		attribute.New(
			"orientation",
			"180",
		),
	},
	Send: {
		attribute.New(
			"shape",
			"cds",
		),
	},
}

func New(t SignalType, name string) Signal {
	return Signal{
		n: node.New(
			fmt.Sprintf("%s <%s>", name, t),
			node.WithAttributes(
				shape[t],
			),
		),
	}
}
func (s Signal) Node() node.Node {
	return s.n
}
func (s Signal) Name() string {
	return s.n.Name
}
func (s Signal) String() string {
	return s.n.String()
}
func (s Signal) EdgeAttrs() string {
	return ""
}
func (s Signal) GetAttrs() []attribute.Attribute {
	return s.n.Attributes
}
func (s Signal) AddAttribute(name, value string) {

}
