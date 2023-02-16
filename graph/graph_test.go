package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	graph := NewGraph("Foo")
	assert.IsType(t, &Graph{}, graph)
}

func TestAddPoint(t *testing.T) {
	graph := NewGraph("Foo")
	graph.AddPoint(NewPoint("point"))
	assert.Equal(
		t,
		"point",
		graph.points.FindByValue("point").Value.(Point).Name, // so ugly
	)
}

func TestAddLine(t *testing.T) {
	graph := NewGraph("Foo")
	graph.AddLine(NewLine("line"))
	assert.Equal(
		t,
		"line",
		graph.lines.FindByValue("line").Value.(Line).Name,
	)
}

func TestGraphAddAttribute(t *testing.T) {
	graph := NewGraph("Foo")
	graph.AddAttribute("cluster", "true")
	assert.Equal(
		t,
		"true",
		graph.attrs.FindByValue("cluster").Value.(Attribute).Value,
	)
}

func TestGetPoint(t *testing.T) {
	graph := NewGraph("Foo")
	p := NewPoint("point")
	graph.AddPoint(p)
	shouldFind, err := graph.GetPoint("point")
	assert.NoError(t, err)
	assert.Equal(t, &p, shouldFind)
	notFind, err := graph.GetPoint("nope")
	assert.ErrorIs(t, err, PointNotFound)
	assert.Nil(t, notFind)
}

func TestGetLine(t *testing.T) {
	graph := NewGraph("Foo")
	l := NewLine("line")
	graph.AddLine(l)
	shouldFind, err := graph.GetLine("line")
	assert.NoError(t, err)
	assert.Equal(t, &l, shouldFind)
	notFind, err := graph.GetLine("nope")
	assert.ErrorIs(t, err, LineNotFound)
	assert.Nil(t, notFind)
}

func TestPoints(t *testing.T) {
	graph := NewGraph("Foo")
	assert.Len(t, graph.Points(), 0)
	p1 := NewPoint("point1")
	p2 := NewPoint("point2")
	graph.AddPoint(p1)
	assert.Len(t, graph.Points(), 1)
	graph.AddPoint(p2)
	assert.Len(t, graph.Points(), 2)
}

func TestLines(t *testing.T) {
	graph := NewGraph("Foo")
	assert.Len(t, graph.Lines(), 0)
	l1 := NewLine("line1")
	graph.AddLine(l1)
	assert.Len(t, graph.Lines(), 1)
	l2 := NewLine("line2")
	graph.AddLine(l2)
	assert.Len(t, graph.Lines(), 2)

}

func TestGetAttributeValue(t *testing.T) {
	graph := NewGraph("Foo")
	graph.AddAttribute("cluster", "true")
	assert.Equal(t, "true", graph.GetAttributeValue("cluster"))
	assert.Equal(t, "", graph.GetAttributeValue("nope"))
}

func TestGetAllAttributes(t *testing.T) {
	graph := NewGraph("Foo")
	assert.Len(t, graph.Attributes(), 0)
	graph.AddAttribute("cluster", "true")
	assert.Len(t, graph.Attributes(), 1)
	graph.AddAttribute("bar", "baz")
	assert.Len(t, graph.Attributes(), 2)
}
