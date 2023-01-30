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
	"time"

	"gopkg.in/yaml.v3"
)

var nodeTemplate, _ = template.New("").Parse(`{{range .Nodes}}
var {{.NodeName}} = diagram.NodeType{
	Name: "{{.Name}}",
	Description: "{{.Description}}",{{if .Attributes}}
	Attributes: map[string]string{
	{{range $k, $v := .Attributes}}	"{{$k}}": "{{$v}}",
	{{end}}},
{{end}}}

func (d {{$.Diagram.DiagramName}}) Add{{.NodeName}}(name string) error {
	if err := d.d.AddNode({{.NodeName}}, name); err != nil {
		return err
	}
	return nil
}
{{end}}var adders = map[string]func({{$.Diagram.DiagramName}}, string) error {
{{range .Nodes}}"{{.NodeName}}": {{$.Diagram.DiagramName}}.Add{{.NodeName}},
{{end}}}`)

var edgeTemplate, _ = template.New("").Parse(`{{range .Edges}}
var {{.EdgeName}} = diagram.EdgeType{
	Name: "{{.Name}}",
	Description: "{{.Description}}",{{if .Attributes}}
	Attributes: map[string]string{
	{{range $k, $v := .Attributes}}	"{{$k}}": "{{$v}}",
	{{end}}},
{{end}}}
{{end}}`)

var diagramTemplate, _ = template.New("").Parse(`
type {{.DiagramName}} struct {
	d *diagram.Diagram
	name string
}

func New(name string) {{.DiagramName}} {
	diagram := diagram.New(
		diagram.DiagramType{
			Name: "{{.Name}}",
			Description: "{{.Description}}",{{if .Attributes}}
			Attributes: map[string]string{
			{{range $k, $v := .Attributes}}	"{{$k}}": "{{$v}}",
			{{end}}},{{end}}
		},
		diagram.WithNodeTypes({{.NodeNames}}),
		diagram.WithEdgeTypes({{.EdgeNames}}),
	)
	return {{.DiagramName}}{name: name, d: &diagram}
}

func (d {{.DiagramName}}) Connect(t diagram.EdgeType, from, to string) error {
	if err := d.d.AddEdge(t, from, to); err != nil {
		return err
	}
	return nil
}

func (d {{.DiagramName}}) Add(nodeType, name string) error {
	fn, ok := adders[nodeType]
	if !ok {
		return fmt.Errorf("'%s' is not a valid node type", nodeType)
	}
	if err := fn(d, name); err != nil {
		return err
	}
	return nil
}`)

var fnMap = template.FuncMap{
	"timeNow": func() int64 {
		return time.Now().UTC().Unix()
	},
}
var packageTemplate, _ = template.New("").Funcs(fnMap).Parse(`//
// Auto-generated code, DO NOT EDIT
//
// generation Unix UTC time: {{timeNow}}
// edit the yaml file and then run 'generateDiagram.go'
//
package {{.PackageName}}

import (
	"fmt"

	"github.com/JacobTripp/diagrams-as-code/diagram"
)
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
