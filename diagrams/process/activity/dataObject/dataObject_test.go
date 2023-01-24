package dataobject

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/path"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := New("Foo")
	assert.IsType(t, DataObject{}, d)
	assert.Implements(t, (*path.Node)(nil), d)
	assert.Contains(t, d.String(), "shape=\"rect\"")
}

func TestName(t *testing.T) {
	d := New("foo")
	assert.Equal(t, "foo", d.Name())
}
