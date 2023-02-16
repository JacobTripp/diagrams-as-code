package graph

import (
	"errors"
	"fmt"

	bst "github.com/JacobTripp/binarySearchTree"
)

type Graph struct {
	Name   string
	attrs  *bst.BinarySearchTree
	points *bst.BinarySearchTree
	lines  *bst.BinarySearchTree
}

func NewGraph(name string) *Graph {
	points := bst.NewBST(bst.WithSearchable("Name"))
	lines := bst.NewBST(bst.WithSearchable("Name"))
	attrs := bst.NewBST(bst.WithSearchable("Name"))
	return &Graph{
		Name:   name,
		attrs:  &attrs,
		points: &points,
		lines:  &lines,
	}
}

func (g Graph) AddPoint(p Point) {
	g.points.Insert(bst.NewLeaf(p))
}

func (g Graph) AddLine(l Line) {
	g.lines.Insert(bst.NewLeaf(l))
}

func (g Graph) AddAttribute(name, value string) {
	attr := Attribute{Name: name, Value: value}
	g.attrs.Insert(bst.NewLeaf(attr))
}

var PointNotFound = errors.New("point not found")

func (g Graph) GetPoint(search string) (*Point, error) {
	leaf := g.points.FindByValue(search)
	if leaf == nil {
		return nil, fmt.Errorf("%w: could not find '%s'", PointNotFound, search)
	}
	point := leaf.Value.(Point)
	return &point, nil
}

var LineNotFound = errors.New("line not found")

func (g Graph) GetLine(search string) (*Line, error) {
	leaf := g.lines.FindByValue(search)
	if leaf == nil {
		return nil, fmt.Errorf("%w: cound not find '%s'", LineNotFound, search)
	}
	line := leaf.Value.(Line)
	return &line, nil
}

func (g Graph) Points() []Point {
	leafs := g.points.GetAllLeafs()
	rt := make([]Point, len(leafs))
	for i, leaf := range leafs {
		rt[i] = leaf.Value.(Point)
	}
	return rt
}

func (g Graph) Lines() []Line {
	leafs := g.lines.GetAllLeafs()
	rt := make([]Line, len(leafs))
	for i, leaf := range leafs {
		rt[i] = leaf.Value.(Line)
	}
	return rt
}

func (g Graph) GetAttributeValue(name string) string {
	f := g.attrs.FindByValue(name)
	if f == nil {
		return ""
	}
	return f.Value.(Attribute).Value
}

func (g Graph) Attributes() []Attribute {
	leafs := g.attrs.GetAllLeafs()
	rt := make([]Attribute, len(leafs))
	for i, leaf := range leafs {
		rt[i] = leaf.Value.(Attribute)
	}
	return rt
}
