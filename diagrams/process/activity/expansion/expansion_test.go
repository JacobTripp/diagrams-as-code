package expansion

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/JacobTripp/diagrams-as-code/dot/node"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	e := New()
	assert.IsType(t, Expansion{}, e)
	assert.Implements(t, (*path.Node)(nil), &e)
}

func TestString(t *testing.T) {
	e := New()
	expected := "\"Expansion\""
	assert.Contains(t, e.String(), expected)
}

func TestName(t *testing.T) {
	e := New()
	expected := "Expansion"
	assert.Equal(t, expected, e.Name())
}

func TestGetAttrs(t *testing.T) {
	e := New()
	assert.Len(t, e.GetAttrs(), 1)
}

func TestAddAttribute(t *testing.T) {
	e := New()
	e.AddAttribute("foo", "bar")
	assert.Len(t, e.GetAttrs(), 2)
}

func TestNode(t *testing.T) {
	e := New()
	n := e.Node()
	assert.IsType(t, node.Node{}, n)
}

func TestPath(t *testing.T) {
	e := New()
	p := path.New()
	e.AddPath(&p)
	assert.Len(t, e.paths, 0)
}
