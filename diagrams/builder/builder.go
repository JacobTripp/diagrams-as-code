// Package builder proivdes a simplfied way to generate the code for diagrams
//
// A diagram is just an ordered list of nodes, each node has it's own
// attributes and the edges connecting the nodes have their own attributes.
// Bassically all you need to do is define the node/edge names for a diagram
// and their corresponding attributes.

package builder

import (
	"bytes"
	"io"
	"log"
	"text/template"
)

type spec struct {
	Name        string
	Description string
	Attributes  map[string]string
	Template    string
	StructName  string
}

type NodeSpec spec
type EdgeSpec spec

type DiagramSpec struct {
	Name        string
	PackageName string
	StructName  string
	Description string
	Template    string
	Attributes  map[string]string
	Nodes       []NodeSpec
	Edges       []EdgeSpec
}

const (
	codetmpl = "code.go.tmpl"
)

func (d DiagramSpec) Generate(out io.Writer) {
	d.PackageName = string(
		bytes.ReplaceAll(
			bytes.ToLower(bytes.NewBufferString(d.Name).Bytes()),
			[]byte(" "),
			[]byte(""),
		),
	)
	d.StructName = string(
		bytes.ReplaceAll(
			bytes.Title(bytes.NewBufferString(d.Name).Bytes()),
			[]byte(" "),
			[]byte(""),
		),
	)
	tmpl, err := template.ParseFiles(codetmpl)
	if err != nil {
		log.Panic(err)
	}
	tmpl.Execute(out, d)
}

func (d *DiagramSpec) AddEdgeSpec(e EdgeSpec) {
	e.StructName = string(bytes.ReplaceAll(
		bytes.Title(bytes.NewBufferString(e.Name).Bytes()),
		[]byte(" "),
		[]byte(""),
	))
	d.Edges = append(d.Edges, e)
}

func (d *DiagramSpec) AddNodeSpec(n NodeSpec) {
	n.StructName = string(bytes.ReplaceAll(
		bytes.Title(bytes.NewBufferString(n.Name).Bytes()),
		[]byte(" "),
		[]byte(""),
	))
	d.Nodes = append(d.Nodes, n)
}
