package signal

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/stretchr/testify/assert"
)

func TestNewRecieve(t *testing.T) {
	s := New(Recieve, "foo")
	assert.IsType(t, Signal{}, s)
	assert.Implements(t, (*path.Node)(nil), s)
	assert.Contains(t, s.String(), "shape=\"cds\"")
	assert.Contains(t, s.String(), "orientation=\"180\"")
	assert.Contains(t, s.String(), "\"foo <recieve>\"")
}

func TestNewWait(t *testing.T) {
	s := New(Send, "foo")
	assert.IsType(t, Signal{}, s)
	assert.Implements(t, (*path.Node)(nil), s)
	assert.Contains(t, s.String(), "shape=\"cds\"")
	assert.Contains(t, s.String(), "\"foo <send>\"")
}

func TestString(t *testing.T) {
	s := New(Send, "foo")
	expected := "\"foo <send>\" [shape=\"cds\"]\n"
	assert.Equal(t, expected, s.String())
}

func TestEdgeAttrs(t *testing.T) {
	s := New(Recieve, "foo")
	assert.Empty(t, s.EdgeAttrs())
}

func TestAddAttribute(t *testing.T) {
	s := New(Recieve, "foo")
	s.AddAttribute("bar", "baz")
	assert.Len(t, s.attrs, 0)
}
