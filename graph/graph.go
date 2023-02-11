package graph

import bst "github.com/JacobTripp/binarySearchTree"

type Graph struct {
	Name   string
	attrs  *bst.BinarySearchTree
	Points *bst.BinarySearchTree
	Lines  *bst.BinarySearchTree
}

/*
what if connections are kept in a matrix where connections are represented
with a line, for example:
point1 = NewPoint()
point2 = NewPoint()
point3 = NewPoint()
graph.AddPoint(point1, point2, point3)
graph = [
	[0,line1,0], //point1
	[line1,0,line2], //point2
	[0,line2,0], //point3
]
*/
func NewGraph(name string) *Graph {
	points := bst.NewBST(bst.WithSearchable("Name"))
	lines := bst.NewBST(bst.WithSearchable("Name"))
	attrs := bst.NewBST(bst.WithSearchable("Name"))
	return &Graph{
		Name:   name,
		attrs:  &attrs,
		Points: &points,
		Lines:  &lines,
	}
}

func (g Graph) AddPoint(p *Point) {
}

func (g Graph) AddLine(l *Line) {
}

func (g Graph) AddAttribute(name, value string) {
}
