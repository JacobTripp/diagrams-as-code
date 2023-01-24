package usecase

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/cluster"
	"github.com/JacobTripp/diagrams-as-code/dot/edge"
	"github.com/JacobTripp/diagrams-as-code/dot/graph"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Diagram struct {
	graph graph.Graph
}

func NewDiagram(systemName string) Diagram {
	label := attribute.New("label", "Use Cases")
	direction := attribute.New("rankdir", "LR")
	layout := attribute.New("layout", "dot")
	splines := attribute.New("splines", "true")
	constraint := attribute.New("constraint", "false")
	nodesep := attribute.New("nodesep", ".5")
	ranksep := attribute.New("ranksep", "1.0 equally")
	actors := cluster.New(
		"actors",
		cluster.WithAttribute(
			attribute.New(
				"label",
				"",
			),
		),
		cluster.WithAttribute(
			attribute.New(
				"cluster",
				"false",
			),
		),
		cluster.WithAttribute(
			attribute.New(
				"pencolor",
				"#ffffff",
			),
		),
	)
	usecases := cluster.New(
		"usecases",
		cluster.WithAttribute(
			attribute.New(
				"style",
				"rounded",
			),
		),
		cluster.WithAttribute(
			attribute.New(
				"label",
				systemName,
			),
		),
		cluster.WithAttribute(
			attribute.New(
				"bgcolor",
				"#eeeeee",
			),
		),
		cluster.WithAttribute(
			attribute.New(
				"pencolor",
				"#eeeeee",
			),
		),
	)
	return Diagram{
		graph: graph.New(
			graph.WithAttribute(label),
			graph.WithAttribute(direction),
			graph.WithAttribute(layout),
			graph.WithAttribute(splines),
			graph.WithAttribute(constraint),
			graph.WithAttribute(ranksep),
			graph.WithAttribute(nodesep),
			graph.WithCluster(actors),
			graph.WithCluster(usecases),
		),
	}
}

func nodeExists(nodes []node.Node, a node.Node) bool {
	for _, node := range nodes {
		if node.Name == a.Name {
			return true
		}
	}
	return false
}

type addActorOption func(Actor, *Diagram)

// this could make an infinite loop if the AddGeneralization
// function is ever updated to use WithGeneralizaion
func WithGeneralization(parent Actor) addActorOption {
	return func(a Actor, ucd *Diagram) {
		ucd.AddGeneralization(parent, a)
	}
}

func (ucd *Diagram) AddGeneralization(parent Actor, child Actor) {
	_ = ucd.AddActor(parent)
	_ = ucd.AddActor(child)
	edge := edge.New(
		edge.WithAttribute(
			attribute.New(
				"arrowhead",
				"empty",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"style",
				"solid",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"edgetooltip",
				"generalization",
			),
		),
	)
	edge.From(node.Node(child))
	edge.To(node.Node(parent))
	ucd.graph.AddEdge(edge)
}

func (ucd *Diagram) AddActorStr(name string) error {
	a := NewActor(name)
	err := ucd.AddActor(a)
	if err != nil {
		return err
	}
	return nil
}

func (ucd *Diagram) AddUseCaseStr(name string) error {
	a := NewUseCase(name)
	err := ucd.AddUseCase(a)
	if err != nil {
		return err
	}
	return nil
}

func (ucd *Diagram) AddActor(a Actor, opts ...addActorOption) error {
	actors := &ucd.graph.Clusters[0] // temporal coupling
	if nodeExists(actors.Nodes, node.Node(a)) {
		return fmt.Errorf("actor '%s' already exists", a.Name)
	}
	actors.AddNode(node.Node(a))
	for _, opt := range opts {
		opt(a, ucd)
	}
	return nil
}

func (ucd *Diagram) AddUseCase(u UseCase) error {
	usecases := &ucd.graph.Clusters[1] // temporal coupling
	if nodeExists(usecases.Nodes, u.node) {
		return fmt.Errorf("usecase '%s' already exists", u.node.Name)
	}
	usecases.AddNode(u.node)
	return nil
}

func (ucd *Diagram) AddCommunication(a Actor, u UseCase) error {
	_ = ucd.AddActor(a)
	_ = ucd.AddUseCase(u)
	edge := edge.New(
		edge.WithAttribute(
			attribute.New(
				"arrowhead", "none",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"edgetooltip", "communication",
			),
		),
	)
	edge.From(node.Node(a))
	edge.To(u.node)
	ucd.graph.AddEdge(edge)
	return nil
}

func (ucd *Diagram) AddInheritance(parent UseCase, child UseCase) error {
	_ = ucd.AddUseCase(parent)
	_ = ucd.AddUseCase(child)
	edge := edge.New(
		edge.WithAttribute(
			attribute.New(
				"style", "bold",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"arrowhead", "empty",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"edgetooltip", "inherits",
			),
		),
	)
	edge.From(child.node)
	edge.To(parent.node)
	ucd.graph.AddEdge(edge)
	return nil
}
func (ucd *Diagram) AddInclude(u1 UseCase, u2 UseCase) error {
	_ = ucd.AddUseCase(u1)
	_ = ucd.AddUseCase(u2)
	edge := edge.New(
		edge.WithAttribute(
			attribute.New(
				"arrowhead", "open",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"style", "dashed",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"label", "<<include>>",
			),
		),
		edge.WithAttribute(
			attribute.New(
				"edgetooltip", "includes",
			),
		),
	)
	edge.From(u1.node)
	edge.To(u2.node)
	ucd.graph.AddEdge(edge)
	return nil
}

func (ucd *Diagram) Actor(name string) Actor {
	for _, node := range ucd.graph.Clusters[0].Nodes {
		if node.Name == name {
			return Actor(node)
		}
	}
	a := NewActor(name)
	ucd.AddActor(a)
	return a
}

func (ucd *Diagram) UseCase(name string) UseCase {
	for _, node := range ucd.graph.Clusters[1].Nodes {
		if node.Name == name {
			return UseCase{node: node}
		}
	}
	u := NewUseCase(name)
	ucd.AddUseCase(u)
	return u
}

func (ucd Diagram) String() string {
	return ucd.graph.String()
}

type BuilderMap map[string][]string

type builder func(*Diagram, string, string)

func (d *Diagram) build(fn builder, set BuilderMap) {
	for x, ys := range set {
		for _, y := range ys {
			fn(d, x, y)
		}
	}
}

func commuFn(d *Diagram, a string, u string) {
	_ = d.AddCommunication(
		d.Actor(a),
		d.UseCase(u),
	)
}
func incluFn(d *Diagram, parent string, child string) {
	_ = d.AddInclude(
		d.UseCase(parent),
		d.UseCase(child),
	)
}
func inherFn(d *Diagram, parent string, child string) {
	_ = d.AddInheritance(
		d.UseCase(parent),
		d.UseCase(child),
	)
}
func (d *Diagram) CommunicationBuilder(bm BuilderMap) {
	d.build(commuFn, bm)
}
func (d *Diagram) IncludeBuilder(bm BuilderMap) {
	d.build(incluFn, bm)
}
func (d *Diagram) InheritanceBuilder(bm BuilderMap) {
	d.build(inherFn, bm)
}
