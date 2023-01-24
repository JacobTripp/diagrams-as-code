package transformation

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Transformation struct {
	n node.Node
}

func New(name string) Transformation {
	tr := Transformation{
		n: node.New("Transformation"),
	}
	tr.AddAttribute("shape", "note")
	tr.AddAttribute("label", fmt.Sprintf("<<transformation>>\n%s", name))
	return tr
}

func (tr Transformation) String() string {
	return tr.n.String()
}

func (tr Transformation) Name() string {
	return tr.n.Name
}

func (tr Transformation) GetAttrs() []attribute.Attribute {
	return tr.n.Attributes
}

func (tr *Transformation) AddAttribute(name, value string) {
	tr.n.AddAttr(
		attribute.New(name, value),
	)
}

func (tr Transformation) Node() node.Node {
	return tr.n
}
