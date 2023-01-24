package usecase

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Actor node.Node

func NewActor(name string) Actor {
	return Actor(
		node.New(
			name,
			node.WithAttribute(
				attribute.New(
					"label",
					fmt.Sprintf("%s <<actor>>", name),
				),
			),
			node.WithAttribute(
				attribute.New(
					"group",
					"actor",
				),
			),
			node.WithAttribute(
				attribute.New(
					"shape",
					"rect",
				),
			),
		))
}
