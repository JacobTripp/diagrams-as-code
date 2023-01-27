package activity

import (
	"github.com/JacobTripp/diagrams-as-code/dot/graph"
)

type Diagram struct {
	g graph.Graph
}

func New(name string) Diagram {
	return Diagram{}
}
