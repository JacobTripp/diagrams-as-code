package graph

import (
	"fmt"

	bst "github.com/JacobTripp/binarySearchTree"
)

type Line struct {
	Name  string
	From  *Point
	To    *Point
	attrs *bst.BinarySearchTree
}

type lineOpts func(*Line)

func From(p *Point) lineOpts {
	return func(l *Line) {
		l.From = p
	}
}
func To(p *Point) lineOpts {
	return func(l *Line) {
		l.To = p
	}
}

func NewLine(name string, opts ...lineOpts) *Line {
	bst := bst.NewBST(bst.WithSearchable("Name"))
	rt := &Line{
		Name:  name,
		attrs: &bst,
	}
	for _, opt := range opts {
		opt(rt)
	}
	return rt
}

func (l Line) AddAttribute(name, value string) {
	l.attrs.Insert(bst.NewLeaf(Attribute{Name: name, Value: value}))
}

func (l Line) GetAttrValue(name string) (any, error) {
	found := l.attrs.FindByValue(name)
	if found == nil {
		return nil, fmt.Errorf("could not find attribute: %s", name)
	}

	return found.Value.(Attribute).Value, nil
}
