package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint("foo")
	assert.IsType(t, &Point{}, p)
}

func TestAddAttribute(t *testing.T) {
	p := NewPoint("foo")
	p.AddAttribute("color", "blue")
	attr, err := p.GetAttrValue("color")
	assert.NoError(t, err)
	assert.Equal(t, "blue", attr)
}
