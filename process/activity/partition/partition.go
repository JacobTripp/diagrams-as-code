package partition

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/process/activity/action"
)

type Partition map[string]string

func New(name string, color string) Partition {
	p := Partition{
		"color":  color,
		"xlabel": fmt.Sprintf("(%s)", name),
	}
	return p
}

func (p Partition) AddNode(a *action.Action) {
	for name, value := range p {
		a.AddAttribute(name, value)
	}
}
