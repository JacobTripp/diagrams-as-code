package interrupt

import (
	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/signal"
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/cluster"
	"github.com/JacobTripp/diagrams-as-code/dot/edge"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Interrupt struct {
	// need node.Node, all items are nodes that can be chained in Path
	c cluster.Cluster
}

func New() Interrupt {
	return Interrupt{
		c: cluster.New("Interrupt"),
	}
}

func (i *Interrupt) On(name string, p path.Path) {
	n := signal.New(signal.Recieve, name)
	i.c.AddNode(n.Node())
	e := edge.New()
	e.From(n.Node())
	e.To(p.Head().Node())
	e.AddAttr(attribute.Attribute{
		"style",
		"dashed",
	})
	i.c.AddEdge(e.String())
}

func (i Interrupt) Node() node.Node { return node.Node{} }
func (i Interrupt) String() string {
	return i.c.String()
}

func (i Interrupt) Name() string {
	return i.c.Name
}

func (i Interrupt) GetAttrs() []attribute.Attribute {
	return i.c.Attributes
}

func (i *Interrupt) AddAttribute(name, value string) {
	i.c.AddAttr(
		attribute.New(name, value),
	)
}
