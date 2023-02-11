package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	uc := New()
	assert.Contains(t, uc.GetAttributeValue("description"), "customer's")
}

func TestAddCase(t *testing.T) {
	uc := New()
	uc.AddCase("Wash Car")
	assert.Len(t, uc.Points(), 1)
	uc.AddCase("Vacuum Car")
	assert.Len(t, uc.Points(), 2)
}

func TestAddActor(t *testing.T) {
	uc := New()
	uc.AddActor("Washer")
	assert.Len(t, uc.Points(), 1)
	uc.AddActor("Customer")
	assert.Len(t, uc.Points(), 2)
}
