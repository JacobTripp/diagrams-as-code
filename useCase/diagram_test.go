package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUseCaseDiagram(t *testing.T) {
	diagram := NewDiagram("Foo")
	assert.IsType(t, Diagram{}, diagram)
}

func TestAddActor(t *testing.T) {
	diagram := NewDiagram("Foo")
	actor := NewActor("Foo")
	err := diagram.AddActor(actor)
	assert.Empty(t, err)
	assert.Len(t, diagram.graph.Clusters[0].Nodes, 1)
}

func TestAddDuplicateActor(t *testing.T) {
	diagram := NewDiagram("Foo")
	actor := NewActor("FooActor")
	err := diagram.AddActor(actor)
	assert.Empty(t, err)
	err = diagram.AddActor(actor)
	assert.Error(t, err)
}

func TestWithGeneralization(t *testing.T) {
	diagram := NewDiagram("Foo")
	parent := NewActor("ParentActor")
	child := NewActor("ChildActor")
	_ = diagram.AddActor(parent)
	_ = diagram.AddActor(child, WithGeneralization(parent))
	expected := "\"ChildActor\"->\"ParentActor\" [arrowhead=\"empty\"style=\"solid\"edgetooltip=\"generalization\"]"
	assert.Contains(t, diagram.String(), expected)
}

func TestAddGeneralization(t *testing.T) {
	diagram := NewDiagram("Foo")
	parent := NewActor("ParentActor")
	child := NewActor("ChildActor")
	diagram.AddGeneralization(parent, child)
	expected := "\"ChildActor\"->\"ParentActor\" [arrowhead=\"empty\"style=\"solid\"edgetooltip=\"generalization\"]"
	assert.Contains(t, diagram.String(), expected)
}

func TestAddUseCase(t *testing.T) {
	diagram := NewDiagram("Foo")
	usecase := NewUseCase("FooUseCase")
	err := diagram.AddUseCase(usecase)
	assert.Empty(t, err)
	assert.Len(t, diagram.graph.Clusters[1].Nodes, 1)
}

func TestAddDuplicateUseCase(t *testing.T) {
	diagram := NewDiagram("Foo")
	usecase := NewUseCase("FooUseCase")
	err := diagram.AddUseCase(usecase)
	assert.Empty(t, err)
	err = diagram.AddUseCase(usecase)
	assert.Error(t, err)
}

func TestActorUsesUseCase(t *testing.T) {
	diagram := NewDiagram("Foo")
	usecase := NewUseCase("FooUseCase")
	actor := NewActor("FooActor")
	_ = diagram.AddUseCase(usecase)
	_ = diagram.AddActor(actor)
	err := diagram.AddCommunication(actor, usecase)
	assert.Empty(t, err)
	assert.Len(t, diagram.graph.Edges, 1)
}

func TestAddInclude(t *testing.T) {
	diagram := NewDiagram("Foo")
	usecase1 := NewUseCase("FooUseCase1")
	usecase2 := NewUseCase("FooUseCase2")
	_ = diagram.AddUseCase(usecase1)
	_ = diagram.AddUseCase(usecase2)
	err := diagram.AddInclude(usecase1, usecase2)
	assert.Empty(t, err)
	assert.Len(t, diagram.graph.Edges, 1)
	expected := "[arrowhead=\"open\"style=\"dashed\"label=\"<<include>>\"edgetooltip=\"includes\"]"
	assert.Contains(t, diagram.String(), expected)
}

func TestAddActorStr(t *testing.T) {
	diagram := NewDiagram("Foo")
	diagram.AddActorStr("Bar")
	assert.Len(t, diagram.graph.Clusters[0].Nodes, 1)
	err := diagram.AddActorStr("Bar")
	assert.Error(t, err)
}

func TestAddUseCaseStr(t *testing.T) {
	diagram := NewDiagram("Foo")
	diagram.AddUseCaseStr("Bar")
	assert.Len(t, diagram.graph.Clusters[1].Nodes, 1)
	err := diagram.AddUseCaseStr("Bar")
	assert.Error(t, err)
}

func TestActor(t *testing.T) {
	diagram := NewDiagram("Foo")
	diagram.AddActorStr("Bar")
	assert.IsType(t, Actor{}, diagram.Actor("Bar"))
	_ = diagram.Actor("Baz")
	assert.Len(t, diagram.graph.Clusters[0].Nodes, 2)
}

func TestUseCase(t *testing.T) {
	diagram := NewDiagram("Foo")
	diagram.AddUseCaseStr("Bar")
	assert.IsType(t, UseCase{}, diagram.UseCase("Bar"))
	_ = diagram.UseCase("Baz")
	assert.Len(t, diagram.graph.Clusters[1].Nodes, 2)
}

func TestString(t *testing.T) {
	diagram := NewDiagram("Foo")
	usecase := NewUseCase("FooUseCase")
	actor := NewActor("FooActor")
	_ = diagram.AddUseCase(usecase)
	_ = diagram.AddActor(actor)
	_ = diagram.AddCommunication(actor, usecase)
	out := diagram.String()
	expected := `digraph {
label="Use Cases"
rankdir="LR"
layout="dot"
splines="true"
constraint="false"
ranksep="1.0 equally"
nodesep=".5"
subgraph "actors"{
cluster="true"
label=""
cluster="false"
pencolor="#ffffff"
"FooActor" [label="FooActor <<actor>>"group="actor"shape="rect"]
}
subgraph "usecases"{
cluster="true"
style="rounded"
label="Foo"
bgcolor="#eeeeee"
pencolor="#eeeeee"
"FooUseCase" [shape="oval"group="usecase"]
}
"FooActor"->"FooUseCase" [arrowhead="none"edgetooltip="communication"]
}`
	assert.Equal(t, expected, out)
}

func TestAddInheritance(t *testing.T) {
	diagram := NewDiagram("Foo")
	parent := diagram.UseCase("ParentUseCase")
	child := diagram.UseCase("ChildUseCase")
	err := diagram.AddInheritance(parent, child)
	expected := "[style=\"bold\"arrowhead=\"empty\"edgetooltip=\"inherits\"]"
	assert.Empty(t, err)
	assert.Contains(t, diagram.String(), expected)
}

func TestCommunicationBuilder(t *testing.T) {
	diagram := NewDiagram("Foo")
	comms := BuilderMap{
		"Foo": {
			"Bar",
			"Baz",
		},
		"Foo1": {
			"Bar1",
		},
	}
	diagram.CommunicationBuilder(comms)
	assert.Len(t, diagram.graph.Clusters[0].Nodes, 2)
	assert.Len(t, diagram.graph.Clusters[1].Nodes, 3)
}
func TestIncludeBuilder(t *testing.T) {
	diagram := NewDiagram("Foo")
	comms := BuilderMap{
		"Foo": {
			"Bar",
			"Baz",
		},
		"Foo1": {
			"Bar1",
		},
	}
	diagram.IncludeBuilder(comms)
	assert.Len(t, diagram.graph.Clusters[0].Nodes, 0)
	assert.Len(t, diagram.graph.Clusters[1].Nodes, 5)
}
func TestInheritanceBuilder(t *testing.T) {
	diagram := NewDiagram("Foo")
	comms := BuilderMap{
		"Foo": {
			"Bar",
			"Baz",
		},
		"Foo1": {
			"Bar1",
		},
	}
	diagram.InheritanceBuilder(comms)
	assert.Len(t, diagram.graph.Clusters[0].Nodes, 0)
	assert.Len(t, diagram.graph.Clusters[1].Nodes, 5)
}
