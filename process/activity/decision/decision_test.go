package decision

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/process/activity/action"
	"github.com/JacobTripp/diagrams-as-code/diagram/path"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := New()
	assert.IsType(t, Decision{}, d)
	assert.Implements(t, (*path.Node)(nil), &d)
}

func TestString(t *testing.T) {
	d := New()
	assert.Contains(t, d.String(), "shape=\"utr\"")
}

func TestAddPath(t *testing.T) {
	d := New()
	d.AddPath("isFoo")
	d.AddPath("isBar")
	assert.Len(t, d.Paths, 2)
}

func TestAddAction(t *testing.T) {
	d := New()
	d.AddPath("isFoo")
	fooAction := action.New("Foo action")
	d.AddAction("isFoo", &fooAction)
	assert.Equal(t, 2, d.Paths["isFoo"].Len())
	d.AddPath("isBar")
	barAction := action.New("Bar action")
	d.AddAction("isBar", &barAction)
	assert.Equal(t, 2, d.Paths["isBar"].Len())
}
