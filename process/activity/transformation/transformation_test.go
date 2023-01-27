package transformation

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/diagram/path"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tr := New("foo.bar")
	assert.IsType(t, Transformation{}, tr)
	assert.Implements(t, (*path.Node)(nil), &tr)
}

func TestString(t *testing.T) {
	tr := New("foo.bar")
	expected := "\"Transformation\""
	assert.Contains(t, tr.String(), expected)
}

func TestName(t *testing.T) {
	tr := New("foo.bar")
	expected := "Transformation"
	assert.Equal(t, expected, tr.Name())
}

func TestGetAttrs(t *testing.T) {
	tr := New("foo.bar")
	assert.Len(t, tr.GetAttrs(), 2)
}

func TestAddAttribute(t *testing.T) {
	tr := New("foo.bar")
	tr.AddAttribute("foo", "bar")
	assert.Len(t, tr.GetAttrs(), 3)
}
