package finalnode

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/stretchr/testify/assert"
)

func TestNewDefault(t *testing.T) {
	f := New(Default)
	assert.IsType(t, FinalNode{}, f)
	assert.Implements(t, (*path.Node)(nil), &f)
}

func TestNewPathEnd(t *testing.T) {
	f := New(PathEnd)
	assert.IsType(t, FinalNode{}, f)
	assert.Implements(t, (*path.Node)(nil), &f)
	assert.Contains(t, f.String(), "shape=\"circle\"")
}

func TestString(t *testing.T) {
	f := New(Default)
	expected := "\"FinalNode\""
	assert.Contains(t, f.String(), expected)
}

func TestName(t *testing.T) {
	f := New(Default)
	expected := "FinalNode"
	assert.Equal(t, expected, f.Name())
}

func TestGetAttrs(t *testing.T) {
	f := New(Default)
	assert.Len(t, f.GetAttrs(), 3)
}

func TestAddAttribute(t *testing.T) {
	f := New(Default)
	f.AddAttribute("foo", "bar")
	assert.Len(t, f.GetAttrs(), 4)
}
