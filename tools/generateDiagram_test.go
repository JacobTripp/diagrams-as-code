package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var diagramData = []byte(`
diagram:
  name: Activity
  description: >-
    Activity diagrams show steps that a Use Case needs to
    accomplish it's purpose.
  attributes:
    color: green
    cluster: true
nodes:
  - name: Action
    description: One task the activity needs to accomplish
    attributes:
      shape: box
      style: rounded
  - name: Init
    description: Is the most basic way to start an Activity
    attributes:
      shape: circle
      color: black
edges:
  - name: Default
    description: Activity diagrams only have one type of connector
    attributes:
      arrowhead: vee
`)

func setUp(t *testing.T) *os.File {
	f, err := os.CreateTemp("", "testDiagramDef")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write(diagramData); err != nil {
		t.Fatal(err)
	}
	return f
}

func cleanUp(f *os.File) {
	os.Remove(f.Name())
}

func TestReadYaml(t *testing.T) {
	f := setUp(t)
	defer cleanUp(f)
	_, err := readYaml(f.Name())
	assert.Empty(t, err)
	_, err = readYaml("badName")
	assert.Error(t, err)
}

func TestGetDirs(t *testing.T) {
	index, err := getDirs(".")
	assert.Empty(t, err)
	assert.Len(t, index, 1)
}

func TestParse(t *testing.T) {
	f := setUp(t)
	defer cleanUp(f)
	data, _ := readYaml(f.Name())
	d, err := parse(data)
	assert.Empty(t, err)
	assert.Equal(t, "Activity", d.Diagram.Name)
}

func TestReadAndParse(t *testing.T) {
	f := setUp(t)
	defer cleanUp(f)
	_, err := readAndParse(f.Name())
	assert.Empty(t, err)
}

func TestMakeDir(t *testing.T) {
	err := makeDir("foo")
	assert.Empty(t, err)
	err = makeDir("foo")
	assert.Error(t, err)
	os.Remove("foo")
}
