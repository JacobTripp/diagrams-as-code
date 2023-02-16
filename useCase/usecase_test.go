package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
* Example
 */
var yamlData = `---
name: Blog Creator
Actors:
	- Admin
	- Author
Cases:
	- Create Account
	- Delete Blog
	- New Blog
	- Authorize
Inherits:
	Create Account:
		- Authorize
	Delete Blog:
		- Authorize
	New Blog:
		- Authorize
Communicates:
	Admin:
		- Delete Blog
	Author:
		- New Blog
		- Create Account
`

func ExampleMethod() {

	// Output:
}

/*
* End Example
 */

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

func TestCommunication(t *testing.T) {
	uc := New()
	uc.AddActor("Washer")
	uc.AddCase("Wash Car")
	err := uc.Communication("Washer", "Wash Car")
	assert.NoError(t, err)
	assert.Len(t, uc.Lines(), 1)
	assert.Equal(t, "Wash Car", uc.Lines()[0].To.Name)
	assert.Equal(t, "Washer", uc.Lines()[0].From.Name)
	assert.Equal(t, "Washer", uc.Lines()[0].From.Name)
	assert.Len(t, uc.Points(), 2)
	err = uc.Communication("Washer", "Collect Money")
	assert.ErrorIs(t, err, CommunicationError)
	err = uc.Communication("Wash Car", "Washer")
	assert.ErrorIs(t, err, CommunicationError)
	err = uc.Communication("Wash Car", "Wash Car")
	assert.NoError(t, err)
}

func TestListNodes(t *testing.T) {
	uc := New()
	nodes := uc.ListNodes()
	expected := []string{
		"Actor",
		"Case",
	}
	assert.Equal(t, expected, nodes)
}

func TestListEdges(t *testing.T) {
	uc := New()
	edges := uc.ListEdges()
	expected := []string{
		"Communication",
		"Includes",
		"Generalization",
	}
	assert.Equal(t, expected, edges)
}
