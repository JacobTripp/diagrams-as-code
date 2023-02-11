package usecase

import "github.com/JacobTripp/diagrams-as-code/graph"

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
