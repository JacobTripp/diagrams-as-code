package merge

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/process/activity/decision"
	"github.com/JacobTripp/diagrams-as-code/diagram/path"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	m := New()
	assert.IsType(t, Merge{}, m)
	assert.Implements(t, (*path.Node)(nil), &m)
}

func TestString(t *testing.T) {
	m := New()
	expected := "\"Merge\""
	assert.Contains(t, m.String(), expected)
}

func TestName(t *testing.T) {
	m := New()
	expected := "Merge"
	assert.Equal(t, expected, m.Name())
}

func TestGetAttrs(t *testing.T) {
	m := New()
	assert.Len(t, m.GetAttrs(), 2)
}

func TestAddAttribute(t *testing.T) {
	m := New()
	m.AddAttribute("foo", "bar")
	assert.Len(t, m.GetAttrs(), 3)
}

func TestMergeDecision(t *testing.T) {
	m := New()
	d := decision.New()
	d.AddPath("isFoo")
	m.MergeDecision(&d)
	assert.NotEmpty(t, m.d)
}
