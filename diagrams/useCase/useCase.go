package usecase

import (
	"fmt"
	"strings"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type noder interface {
	AddAttr(attribute.Attribute)
	String() string
}

type UseCase struct {
	Name        string
	node        node.Node
	description *Description
}

type Description struct {
	fields map[string]string
}

func NewUseCase(name string) UseCase {
	n := node.New(
		name,
		node.WithAttribute(
			attribute.New(
				"shape",
				"oval",
			),
		),
		node.WithAttribute(
			attribute.New(
				"group",
				"usecase",
			),
		),
	)
	return UseCase{Name: n.Name, node: n}
}

func (uc *UseCase) NewDescription() {
	d := Description{
		fields: map[string]string{},
	}
	uc.description = &d
}
func (uc UseCase) DescriptionString() string {
	var fields strings.Builder
	for name, description := range uc.description.fields {
		fields.WriteString(fmt.Sprintf("%s: %s\n", name, description))
	}
	return fmt.Sprintf("Use case name: %s\n%s", uc.node.Name, fields.String())
}

func (uc *UseCase) AddField(name, details string) {
	uc.description.fields[name] = details
}
