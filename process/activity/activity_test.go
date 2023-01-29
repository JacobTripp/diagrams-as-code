package activity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewActivity(t *testing.T) {
	a := New("Foo Activity")
	assert.NotEmpty(t, a)
}

func TestAddInit(t *testing.T) {
	a := New("Foo Activity")
	err := a.AddInit("")
	assert.Empty(t, err)
}

func TestAddAction(t *testing.T) {
	a := New("Foo Activity")
	err := a.AddAction("Foo Action")
	assert.Empty(t, err)
}

func TestAdd(t *testing.T) {
	a := New("Foo Activity")
	err := a.Add("Action", "Baz Action")
	assert.Empty(t, err)
}

func TestConnect(t *testing.T) {
	a := New("Foo Activity")
	_ = a.AddAction("Foo Action")
	_ = a.AddAction("Bar Action")
	err := a.Connect(Default, "Foo Action", "Bar Action")
	assert.Empty(t, err)
}
