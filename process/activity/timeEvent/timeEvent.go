package timeevent

import (
	"fmt"
	"time"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type TimeEvent struct {
	n node.Node
}

type EventType string

const (
	Wait    = EventType("wait")
	Timeout = EventType("Timeout")
)

func New(t EventType, d time.Duration) TimeEvent {
	name := fmt.Sprintf("%s %s", t, d.String())
	return TimeEvent{
		n: node.New(
			name,
			node.WithAttribute(
				attribute.New(
					"shape",
					"invtriangle",
				),
			),
			node.WithAttribute(
				attribute.New(
					"xlabel",
					name,
				),
			),
			node.WithAttribute(
				attribute.New(
					"label",
					"",
				),
			),
		),
	}
}

func (t TimeEvent) Name() string {
	return t.n.Name
}
func (t TimeEvent) String() string {
	return t.n.String()
}
func (t TimeEvent) EdgeAttrs() string               { return "" }
func (t TimeEvent) AddAttribute(name, value string) {}
func (t TimeEvent) GetAttrs() []attribute.Attribute { return t.n.Attributes }
func (t TimeEvent) Node() node.Node {
	return t.n
}
