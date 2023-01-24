package expansion

import (
	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Expansion struct {
	n     node.Node
	paths []*path.Path
	end   node.Node
}

var style = map[string][]attribute.Attribute{
	"front": {
		attribute.New(
			"shape",
			"noverhang",
		),
	},
	"end": {
		attribute.New(
			"shape",
			"noverhang",
		),
	},
}

func New() Expansion {
	n := node.New("Expansion")
	n.AddAttrs(style["front"])
	end := node.New("Expansion end")
	end.AddAttrs(style["end"])
	return Expansion{
		n:   n,
		end: end,
	}
}

func (e *Expansion) AddPath(p *path.Path) {
	/*
		e.paths = append(e.paths, p)
		// need a generic node type to construct arbitrary nodes
		p.AddNode(e.end)
	*/
}

func (e Expansion) String() string {
	return e.n.String()
}

func (e Expansion) Name() string {
	return e.n.Name
}

func (e Expansion) GetAttrs() []attribute.Attribute {
	return e.n.Attributes
}

func (e *Expansion) AddAttribute(name, value string) {
	e.n.AddAttr(
		attribute.New(name, value),
	)
}

func (e Expansion) Node() node.Node {
	return e.n
}
