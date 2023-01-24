package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDescription(t *testing.T) {
	uc := NewUseCase("Foo")
	uc.NewDescription()
	assert.IsType(t, &Description{}, uc.description)
}

func TestAddField(t *testing.T) {
	uc := NewUseCase("Foo")
	uc.NewDescription()
	uc.AddField("Bar", "Bar details go here.")
	assert.Len(t, uc.description.fields, 1)
}

func TestDescriptionString(t *testing.T) {
	uc := NewUseCase("Foo")
	uc.NewDescription()
	uc.AddField("Bar", "Bar details go here.")
	uc.AddField("Baz", "Baz details go here.")
	expectedName := "Use case name: Foo"
	expectedBar := "Bar details go here."
	expectedBaz := "Baz: Baz details go here."

	assert.Contains(t, uc.DescriptionString(), expectedName)
	assert.Contains(t, uc.DescriptionString(), expectedBar)
	assert.Contains(t, uc.DescriptionString(), expectedBaz)
}
