package timeevent

import (
	"testing"
	"time"

	"github.com/JacobTripp/diagrams-as-code/diagram/path"
	"github.com/stretchr/testify/assert"
)

func TestNewWait(t *testing.T) {
	te := New(Wait, 3*time.Hour)
	assert.IsType(t, TimeEvent{}, te)
	assert.Implements(t, (*path.Node)(nil), te)
	assert.Contains(t, te.String(), "shape=\"invtriangle\"")
	assert.Contains(t, te.String(), "wait 3h0m0s")
	assert.Contains(t, te.String(), "xlabel=\"wait 3h0m0s\"")
}

func TestEdgeAttrs(t *testing.T) {
	te := New(Wait, 3*time.Hour)
	assert.Empty(t, te.EdgeAttrs())
}
