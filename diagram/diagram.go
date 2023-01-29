package diagram

import "fmt"

type entityType struct {
	Name        string
	Description string
	Attributes  map[string]string
}

type DiagramType entityType
type NodeType entityType
type EdgeType entityType
type NodeName string
type EdgeName string
type Connections map[NodeName]EdgeType
type NodeTypes map[NodeName]NodeType
type EdgeTypes map[EdgeName]EdgeType
type Nodes map[NodeName]Node
type diagramOption func(*Diagram)

type Node struct {
	t           NodeType
	value       string
	connections Connections
}

func (n *Node) connectTo(e EdgeType, t NodeName) {
	n.connections[t] = e
}

type Diagram struct {
	t         DiagramType
	Name      string
	nodeTypes NodeTypes
	edgeTypes EdgeTypes
	nodes     Nodes
}

func New(dt DiagramType, opts ...diagramOption) Diagram {
	d := Diagram{
		t:         dt,
		nodeTypes: NodeTypes{},
		edgeTypes: EdgeTypes{},
		nodes:     Nodes{},
	}
	for _, opt := range opts {
		opt(&d)
	}
	return d
}

func WithNodeTypes(types ...NodeType) diagramOption {
	return func(d *Diagram) {
		for _, t := range types {
			d.nodeTypes[NodeName(t.Name)] = t
		}
	}
}

func WithEdgeTypes(types ...EdgeType) diagramOption {
	return func(d *Diagram) {
		for _, t := range types {
			d.edgeTypes[EdgeName(t.Name)] = t
		}
	}
}

func (d Diagram) New(name string) Diagram {
	d.Name = name
	return d
}

func (d *Diagram) AddNode(t NodeType, value string) error {
	_, ok := d.nodeTypes[NodeName(t.Name)]
	if !ok {
		return fmt.Errorf(
			"node type '%s' does not exist in the diagram '%s'",
			t.Name,
			d.t.Name,
		)
	}
	d.nodes[NodeName(value)] = Node{
		t:           t,
		value:       value,
		connections: Connections{},
	}
	return nil
}

func (d *Diagram) AddEdge(t EdgeType, from, to string) error {
	f, ok := d.nodes[NodeName(from)]
	_, ok = d.nodes[NodeName(to)]
	if !ok {
		return fmt.Errorf("node '%s' does not exist", from)
	}
	_, ok = d.edgeTypes[EdgeName(t.Name)]
	if !ok {
		return fmt.Errorf(
			"edge type '%s' doesn not exist in the diagram '%s'",
			t.Name,
			d.t.Name,
		)
	}
	f.connectTo(t, NodeName(to))
	return nil
}
