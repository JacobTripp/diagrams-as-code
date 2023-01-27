// Package builder proivdes a simplfied way to generate the code for diagrams
//
// A diagram is just an ordered list of nodes, each node has it's own
// attributes and the edges connecting the nodes have their own attributes.
// Bassically all you need to do is define the node/edge names for a diagram
// and their corresponding attributes.
package builder

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

var nodeTemplate, _ = template.New("").Parse(`{{range .Nodes}}
var {{.NodeName}} = diagram.NodeType{
	Name: "{{.Name}}",
	Description: "{{.Description}}",{{if .Attributes}}
	Attributes: {
	{{range $k, $v := .Attributes}}	"{{$k}}": "{{$v}}",
	{{end}}},
{{end}}}
{{end}}`)

var edgeTemplate, _ = template.New("").Parse(`{{range .Edges}}
var {{.EdgeName}} = diagram.EdgeType{
	Name: "{{.Name}}",
	Description: "{{.Description}}",{{if .Attributes}}
	Attributes: {
	{{range $k, $v := .Attributes}}	"{{$k}}": "{{$v}}",
	{{end}}},
{{end}}}
{{end}}`)

var diagramTemplate, _ = template.New("").Parse(`
var {{.DiagramName}} = diagram.New(
	DiagramType{
		Name: "{{.Name}}",
		Description: "{{.Description}}",{{if .Attributes}}
		Attributes: {
		{{range $k, $v := .Attributes}}	"{{$k}}": "{{$v}}",
		{{end}}},{{end}}
	},
	WithNodeTypes({{.NodeNames}}),
	WithEdgeTypes({{.EdgeNames}}),
)`)

var packageTemplate, _ = template.New("").Parse(`//
// Auto-generated code, DO NOT EDIT
//
package {{.PackageName}}

import "github.com/JacobTripp/diagrams-as-code/diagram"
`)

type Entity struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Attributes  map[string]string `yaml:"attributes"`
}
type NodeData struct {
	NodeName string
	Entity   `yaml:",inline"`
}

type EdgeData struct {
	EdgeName string
	Entity   `yaml:",inline"`
}

type DiagramData struct {
	DiagramName string
	NodeNames   string
	EdgeNames   string
	Entity      `yaml:",inline"`
}

type YamlData struct {
	Diagram DiagramData `yaml:"diagram"`
	Nodes   []NodeData  `yaml:"nodes"`
	Edges   []EdgeData  `yaml:"edges"`
}

func ImportYaml(y []byte) (YamlData, error) {
	var data YamlData
	err := yaml.Unmarshal(y, &data)
	if err != nil {
		return data, err
	}
	var nodes strings.Builder
	for i, node := range data.Nodes {
		n := strings.ReplaceAll(strings.Title(node.Name), " ", "")
		data.Nodes[i].NodeName = n
		nodes.WriteString(fmt.Sprintf("%s,", n))
	}
	var edges strings.Builder
	for i, edge := range data.Edges {
		n := strings.ReplaceAll(strings.Title(edge.Name), " ", "")
		data.Edges[i].EdgeName = n
		edges.WriteString(fmt.Sprintf("%s,", n))
	}
	data.Diagram.DiagramName = strings.ReplaceAll(
		strings.Title(data.Diagram.Name), " ", "",
	)
	data.Diagram.NodeNames = strings.TrimSuffix(nodes.String(), ",")
	data.Diagram.EdgeNames = strings.TrimSuffix(edges.String(), ",")

	return data, nil
}

func WritePackage(out io.Writer, data YamlData) {
	packageData := struct {
		PackageName string
	}{
		PackageName: strings.ToLower(data.Diagram.DiagramName),
	}
	packageTemplate.Execute(out, packageData)
	nodeTemplate.Execute(out, data)
	edgeTemplate.Execute(out, data)
	diagramTemplate.Execute(out, data.Diagram)
}
