package graph

import (
	"fmt"

	bst "github.com/JacobTripp/binarySearchTree"
)

type Point struct {
	Name  string
	attrs *bst.BinarySearchTree
}

func NewPoint(name string) Point {
	bst := bst.NewBST(bst.WithSearchable("Name"))
	return Point{
		Name:  name,
		attrs: &bst,
	}
}

func (p Point) AddAttribute(name, value string) {
	p.attrs.Insert(bst.NewLeaf(Attribute{Name: name, Value: value}))
}

func (p Point) GetAttrValue(name string) (any, error) {
	found := p.attrs.FindByValue(name)
	if found == nil {
		return nil, fmt.Errorf("could not find attribute: %s", name)
	}

	return found.Value.(Attribute).Value, nil
}

func (p Point) Attributes() []Attribute {
	leafs := p.attrs.GetAllLeafs()
	rt := make([]Attribute, len(leafs))
	for i, leaf := range leafs {
		rt[i] = leaf.Value.(Attribute)
	}
	return rt
}
