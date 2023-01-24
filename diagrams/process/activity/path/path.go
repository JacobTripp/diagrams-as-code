package path

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type Path struct {
	nodes   list.List
	initSet bool
}

type Node interface {
	String() string
	Name() string
	//EdgeAttrs() string
	AddAttribute(string, string)
	GetAttrs() []attribute.Attribute
	Node() node.Node
}

func New() Path {
	return Path{
		initSet: false,
	}
}

func (p *Path) AddInit(n Node) error {
	if p.initSet {
		return fmt.Errorf("path already has an init node set")
	}
	p.nodes.PushFront(n)
	p.initSet = true
	return nil
}

func (p Path) Len() int {
	return p.nodes.Len()
}

func (p *Path) AddNode(n Node) {
	p.nodes.PushBack(n)
}

func (p Path) PathString() string {
	var sb strings.Builder
	for e := p.nodes.Front(); e != nil; e = e.Next() {
		sb.WriteString(fmt.Sprintf("\"%s\"", e.Value.(Node).Name()))
		if e.Next() != nil {
			sb.WriteString("->")
		}
	}
	return sb.String()
}

func (p Path) NodeString() string {
	var sb strings.Builder
	for e := p.nodes.Front(); e != nil; e = e.Next() {
		sb.WriteString(fmt.Sprintf("%s", e.Value.(Node).String()))
	}
	return sb.String()
}

func (p Path) Head() Node {
	return p.nodes.Front().Value.(Node)
}
