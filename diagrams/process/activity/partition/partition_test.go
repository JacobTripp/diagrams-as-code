package partition

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/diagrams/process/activity/action"
	"github.com/JacobTripp/diagrams-as-code/dot/attribute"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := New("foo", "blue")
	assert.IsType(t, Partition{}, p)
}

func TestAddNode(t *testing.T) {
	p := New("foo", "blue")
	a := action.New("bar activity")
	p.AddNode(&a)
	assert.Contains(t, a.GetAttrs(), attribute.New("color", "blue"))
}
