package newDiagram

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoadTemplate(t *testing.T) {
	tmpl, sum := loadTemplate("./diagram.go.tmpl")
	assert.Len(t, sum, 32)
	assert.IsType(t, &template.Template{}, tmpl)
	assert.NotEmpty(t, tmpl)
	assert.Panics(t, func() { loadTemplate("./nope") })
}

func TestRecieverName(t *testing.T) {
	rn := recieverName("FooBar")
	assert.Equal(t, "_fo", rn)
	rn = recieverName("F")
	assert.Equal(t, "_f", rn)
	assert.Panics(t, func() { recieverName("") })
	assert.Panics(t, func() { recieverName("  ") })
	rn = recieverName("F bar")
	assert.Equal(t, "_fb", rn)
	rn = recieverName("F ")
	assert.Equal(t, "_f", rn)
	rn = recieverName("FO")
	assert.Equal(t, "_fo", rn)
}

type mockDiagram struct {
	mock.Mock
}

func (m *mockDiagram) GetString(s string) string {
	return "GetString"
}

func (m *mockDiagram) Get(s string) interface{} {
	if s == "nodes" || s == "edges" {
		return []map[string]string{
			{
				"name": "foo",
			},
			{
				"name": "bar",
			},
		}
	}
	return "Get"
}

func TestRender(t *testing.T) {
	code := bytes.NewBuffer([]byte{})
	dataMock := new(mockDiagram)
	dataMock.On("GetString").Return("get string")
	dataMock.On("Get").Return("get")
	err := Render(code, "./diagram.go.tmpl", new(mockDiagram))
	assert.NoError(t, err)
	assert.Contains(t, code.String(), "// Auto-generated; DO NOT EDIT")
}

func TestMain(t *testing.T) {

}
