//
// Auto-generated code, DO NOT EDIT
//
package usecase

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/diagram"
)

var UseCase = diagram.NodeType{
	Name: "Use Case",
	Description: "",
	Attributes: map[string]string{
		"cluster": "true",
		"color": "gray",
		"shape": "oval",
	},
}

func (d UseCase) AddUseCase(name string) error {
	if err := d.d.AddNode(UseCase, name); err != nil {
		return err
	}
	return nil
}

var Actor = diagram.NodeType{
	Name: "Actor",
	Description: "",
	Attributes: map[string]string{
		"shape": "rect",
	},
}

func (d UseCase) AddActor(name string) error {
	if err := d.d.AddNode(Actor, name); err != nil {
		return err
	}
	return nil
}
var adders = map[string]func(UseCase, string) error {
"UseCase": UseCase.AddUseCase,
"Actor": UseCase.AddActor,
}
var Communication = diagram.EdgeType{
	Name: "Communication",
	Description: "A communication line connects and actor and a use case to show the actor participating in the use case. There is a potential to have any number of actors involved in a use case, there is no real limit. The purpose of a communication line is to show that an actor is simply involved in a use case, not to imply an information exchange in any particular direction or that the actor starts the use case. That type of information can be added to the use case's description.",
	Attributes: map[string]string{
		"arrowhead": "none",
		"style": "solid",
	},
}

var Includes = diagram.EdgeType{
	Name: "Includes",
	Description: "When there is repetitive behavior shared between two use cases then that behavior is best separated and captured within a new use case. This new use case can then be reused by other use cases using and include relationship. For example if two use cases need to check a credential database (that's an actor) then a new use case of 'Check Identity' should be created and then any use case that needs to check identity then includes that use case. The include relationship declares that the use case including the other completely reuses all of the steps from the use case being included.",
	Attributes: map[string]string{
		"arrowhead": "open",
		"label": "<<include>>",
		"style": "dashed",
	},
}

var Generalizaion = diagram.EdgeType{
	Name: "Generalizaion",
	Description: "Sometimes there are use cases whose behavior can be applied to several different cases, but with small changes. Unlike the include relationship that uses everything as is without changes, this relationship allows you to reuse a subset of behavior with small changes for a collection of specific situations. In object-oriented terms, you have a number of specialized cases of a generalized use case. For example, currently you have a 'Create Account' use case but what if you want to have different types of accounts that differ slightly from each other? You would want to describe the general behavior use case of 'Create Account' and then define specialized use cases in which the account being created is a specific type, such as a regular account or a management account. By using inheritance, you are saying that every step in the general use case must occur in the specialized use case.",
	Attributes: map[string]string{
		"arrowhead": "empty",
		"style": "solid",
	},
}

type UseCase struct {
	d *diagram.Diagram
	name string
}

func New(name string) UseCase {
	diagram := diagram.New(
		diagram.DiagramType{
			Name: "Use Case",
			Description: "A use case is a situation where your system is used to fulfill one or more of your customer's requirements; a use case captures a piece of functionality that the system provides. Use cases are the starting point of your model since they affect and guide all the other elements within your system's design. They describe a system's requirements strictly from the outside looking in; They specify the value that the system delivers to customers.",
		},
		diagram.WithNodeTypes(UseCase,Actor),
		diagram.WithEdgeTypes(Communication,Includes,Generalizaion),
	)
	return UseCase{name: name, d: &diagram}
}

func (d UseCase) Connect(t diagram.EdgeType, from, to string) error {
	if err := d.d.AddEdge(t, from, to); err != nil {
		return err
	}
	return nil
}

func (d UseCase) Add(nodeType, name string) error {
	fn, ok := adders[nodeType]
	if !ok {
		return fmt.Errorf("'%s' is not a valid node type", nodeType)
	}
	if err := fn(d, name); err != nil {
		return err
	}
	return nil
}