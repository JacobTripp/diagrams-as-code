package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLine(t *testing.T) {
	line := NewLine("foo")
	assert.IsType(t, Line{}, line)
}

func TestLineAddAttribute(t *testing.T) {
	line := NewLine("foo")
	line.AddAttribute("style", "solid")
	value, err := line.GetAttributeValue("style")
	assert.NoError(t, err)
	assert.Equal(t, "solid", value)
}

func TestNewFrom(t *testing.T) {
	line := NewLine("foo", From(&Point{}))
	assert.NotNil(t, line.From)
}

func TestNewTo(t *testing.T) {
	line := NewLine("foo", To(&Point{}))
	assert.NotNil(t, line.To)
}

func TestAttributes(t *testing.T) {
	line := NewLine("foo")
	assert.Len(t, line.Attributes(), 0)
	line.AddAttribute("style", "solid")
	assert.Len(t, line.Attributes(), 1)
	line.AddAttribute("bar", "baz")
	assert.Len(t, line.Attributes(), 2)
}
