package interrupt

import (
	"testing"

	initnode "github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/initNode"
	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	i := New()
	assert.IsType(t, Interrupt{}, i)
	assert.Implements(t, (*path.Node)(nil), &i)
}

func TestString(t *testing.T) {
	i := New()
	expected := "\"Interrupt\""
	assert.Contains(t, i.String(), expected)
}

func TestName(t *testing.T) {
	i := New()
	expected := "Interrupt"
	assert.Equal(t, expected, i.Name())
}

func TestGetAttrs(t *testing.T) {
	i := New()
	assert.Len(t, i.GetAttrs(), 1)
}

func TestAddAttribute(t *testing.T) {
	i := New()
	i.AddAttribute("foo", "bar")
	assert.Len(t, i.GetAttrs(), 2)
}

func TestOn(t *testing.T) {
	i := New()
	p := path.New()
	p.AddInit(initnode.New(initnode.Default))
	i.On("foo", p)
	assert.Len(t, i.c.Edges, 1)
}
