package main

import (
	"flag"
	"fmt"
	"go/token"
	"os"
	"strings"
	"text/template"
)

var codeTmpl = `package {{.LowerName}}

import (
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
)

type {{.Name}} struct {
  n node.Node
}

func New() {{.Name}} {
  return {{.Name}}{
    n: node.New("{{.Name}}"),
  }
}

func ({{.Initial}} {{.Name}}) String() string {
  return {{.Initial}}.n.String()
}

func ({{.Initial}} {{.Name}}) Name() string {
  return {{.Initial}}.n.Name
}

func ({{.Initial}} {{.Name}}) GetAttrs() []attribute.Attribute {
  return {{.Initial}}.n.Attributes
}

func ({{.Initial}} *{{.Name}}) AddAttribute(name, value string) {
  {{.Initial}}.n.AddAttr(
    attribute.New(name, value),
  )
}

func ({{.Initial}} {{.Name}}) Node() node.Node{
  return {{.Initial}}.n
}`

var testTmpl = `package {{.LowerName}}

import (
  "testing"

	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
  {{.Initial}} := New()
  assert.IsType(t, {{.Name}}{}, {{.Initial}})
  assert.Implements(t, (*path.Node)(nil), &{{.Initial}})
}

func TestString(t *testing.T) {
  {{.Initial}} := New()
  expected := "\"{{.Name}}\""
  assert.Contains(t, {{.Initial}}.String(), expected)
}

func TestName(t *testing.T) {
  {{.Initial}} := New()
  expected := "{{.Name}}"
  assert.Equal(t, expected, {{.Initial}}.Name())
}

func TestGetAttrs(t *testing.T) {
  {{.Initial}} := New()
  assert.Len(t, {{.Initial}}.GetAttrs(), 0)
}

func TestAddAttribute(t *testing.T) {
  {{.Initial}} := New()
  {{.Initial}}.AddAttribute("foo", "bar")
  assert.Len(t, {{.Initial}}.GetAttrs(), 1)
}

func TestNode(t *testing.T) {
  {{.Initial}} := New()
  n := {{.Initial}}.Node()
  assert.IsType(t, node.Node{}, n)
}`

type node struct {
	Name      string
	Initial   string
	LowerName string
}

func main() {
	nodeName := flag.String("name", "", "The name of the activity node")
	flag.Parse()
	name := string(*nodeName)
	if name == "" {
		panic("must include a node name")
	}
	initial := strings.ToLower(string(name[0]))
	if initial == "t" {
		if token.IsKeyword(strings.ToLower(string(name[:2]))) {
			initial = strings.ToLower(name)
		}
		initial = strings.ToLower(string(name[:2]))
	}
	n := node{
		Name:      strings.Title(name),
		Initial:   initial,
		LowerName: strings.ToLower(name),
	}
	code, err := template.New("code").Parse(codeTmpl)
	tests, err := template.New("tests").Parse(testTmpl)
	if err != nil {
		panic(err)
	}
	codeOut, err := os.Create(fmt.Sprintf("%s.go", n.LowerName))
	testsOut, err := os.Create(fmt.Sprintf("%s_test.go", n.LowerName))
	if err != nil {
		panic(err)
	}
	defer codeOut.Close()
	defer testsOut.Close()
	code.Execute(codeOut, n)
	tests.Execute(testsOut, n)
}
