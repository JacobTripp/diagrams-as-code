package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLine(t *testing.T) {
	line := NewLine("foo")
	assert.IsType(t, &Line{}, line)
}

func TestLineAddAttribute(t *testing.T) {
	line := NewLine("foo")
	line.AddAttribute("style", "solid")
	value, err := line.GetAttrValue("style")
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
