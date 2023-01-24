package action

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	action := New("Foo Action")
	assert.IsType(t, Action{}, action)
	assert.Equal(t, "Foo Action", action.n.Name)
	assert.Implements(t, (*path.Node)(nil), &action)
}

func TestNewIsAction(t *testing.T) {
	action := New("foo action", IsActivity())
	assert.Contains(t, action.String(), "label=\"&#8916;")
}

func TestString(t *testing.T) {
	action := New("Foo Action")
	expected := "\"Foo Action\" [shape=\"rect\"style=\"rounded\"]\n"
	assert.Equal(t, expected, action.String())
}

func TestEdgeAttrs(t *testing.T) {
	a := New("foo")
	assert.Empty(t, a.EdgeAttrs())
}
