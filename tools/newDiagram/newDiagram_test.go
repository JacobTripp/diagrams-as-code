package main

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestLoadYaml(t *testing.T) {
	y, sum := loadYaml("./test.yaml")
	assert.NotEmpty(t, sum)
	assert.Len(t, sum, 32)
	assert.IsType(t, &data{}, y)
	assert.NotEmpty(t, y)
}

func TestRender(t *testing.T) {
	tmpl, _ := loadTemplate("./diagram.go.tmpl")
	assert.NotEmpty(t, tmpl)
	y, _ := loadYaml("./test.yaml")
	code := render(tmpl, y)
	t.Log(code)
	assert.Contains(t, code, "// Auto-generated; DO NOT EDIT")
}
