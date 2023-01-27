package activity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := New("foo activity diagram")
	assert.IsType(t, Diagram{}, d)
}
