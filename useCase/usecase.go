package usecase

import (
	"errors"
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/graph"
)

type UseCase struct {
	*graph.Graph
}

func New() UseCase {
	rt := UseCase{graph.NewGraph("Use Case")}
	rt.AddAttribute("description", "Show's the customer's requirements")
	return rt
}

func (uc UseCase) AddCase(name string) {
	uc.AddPoint(newCase(name))
}

func newCase(name string) graph.Point {
	_case := graph.NewPoint(name)
	_case.AddAttribute("type", "Case")
	_case.AddAttribute("shape", "oval")
	_case.AddAttribute("cluster", "true")
	_case.AddAttribute("color", "gray")
	_case.AddAttribute("description", "Specific usecase that solves requirement")
	return _case
}

func (uc UseCase) AddActor(name string) {
	uc.AddPoint(newActor(name))
}

func newActor(name string) graph.Point {
	_actor := graph.NewPoint(name)
	_actor.AddAttribute("type", "Actor")
	_actor.AddAttribute("description", "an entity that you can't control")
	_actor.AddAttribute("shape", "rect")
	return _actor
}

type node interface {
	GetAttributeValue(string) (any, error)
}

func typeString(n node) string {
	v, _ := n.GetAttributeValue("type")
	return v.(string)
}

type connectRule struct {
	from []string
	to   []string
}

var ConnectionError = errors.New("invalid connection")

func connectionValidation(isFromNode bool, nodeType, lineType string) error {
	m := map[string]connectRule{
		"Communication": connectRule{
			from: []string{},
			to:   []string{"Actor"},
		},
		"Generalization": connectRule{
			from: []string{"Actor"},
			to:   []string{"Actor"},
		},
		"Includes": connectRule{
			from: []string{"Actor"},
			to:   []string{"Actor"},
		},
	}
	var list []string
	if isFromNode {
		list = m[lineType].from
	} else {
		list = m[lineType].to
	}
	for _, v := range list {
		if v == nodeType {
			return ConnectionError
		}
	}
	return nil
}

var CommunicationError = errors.New("invalid communication")

func (uc UseCase) Communication(fromName, toName string) error {
	foundTo, err := uc.GetPoint(toName)
	if err == nil {
		err = connectionValidation(false, typeString(foundTo), "Communication")
	}
	if err != nil {
		return fmt.Errorf("%w: %s", CommunicationError, err)
	}

	foundFrom, err := uc.GetPoint(fromName)
	if err == nil {
		err = connectionValidation(true, typeString(foundFrom), "Communication")
	}
	if err != nil {
		return fmt.Errorf("%w: %s", CommunicationError, err)
	}

	line := newCommunication(foundFrom.Name + "communication" + foundTo.Name)
	line.From = foundFrom
	line.To = foundTo
	uc.AddLine(line)
	return nil
}

func newCommunication(name string) graph.Line {
	_line := graph.NewLine(name)
	_line.AddAttribute("type", "Communication")
	_line.AddAttribute("description", "Represents a communication between items.")
	_line.AddAttribute("arrowhead", "none")
	_line.AddAttribute("style", "solid")
	_line.AddAttribute("style", "solid")
	return _line
}

var nodes = []string{
	"Actor",
	"Case",
}

func (uc UseCase) ListNodes() []string {
	return nodes
}

var edges = []string{
	"Communication",
	"Includes",
	"Generalization",
}

func (uc UseCase) ListEdges() []string {
	return edges
}
