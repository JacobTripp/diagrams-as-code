package initnode

import (
	"testing"
	"time"

	"github.com/JacobTripp/diagrams-as-code/diagram/path"
	"github.com/JacobTripp/diagrams-as-code/process/activity/signal"
	te "github.com/JacobTripp/diagrams-as-code/process/activity/timeEvent"
	"github.com/stretchr/testify/assert"
)

func TestNewDefault(t *testing.T) {
	initNode := New(Default)
	assert.IsType(t, InitNode{}, initNode)
	assert.Implements(t, (*path.Node)(nil), initNode)
}

func TestNewTimeEvent(t *testing.T) {
	initNode := New(TimeEvent, UseNode(te.New(te.Timeout, 3*24*time.Hour)))
	assert.Contains(t, initNode.String(), "shape=\"invtriangle\"")
}

func TestNewRecieve(t *testing.T) {
	initNode := New(RecieveEvent, UseNode(signal.New(signal.Recieve, "Foo")))
	assert.Contains(t, initNode.String(), "shape=\"cds\"")
}
