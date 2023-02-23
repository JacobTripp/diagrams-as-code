package node

import (
	"testing"

	"github.com/JacobTripp/diagrams-as-code/pkg/dot/attribute"
	"github.com/stretchr/testify/assert"
)

// New()
var newCases = map[string]struct {
	input    []attribute.Attribute
	expected int
}{
	"Zero Attributes": {
		input:    []attribute.Attribute{},
		expected: 0,
	},
	"One Attribute": {
		input: []attribute.Attribute{
			attribute.Attribute{},
		},
		expected: 1,
	},
	"Two Attributes": {
		input: []attribute.Attribute{
			attribute.Attribute{},
			attribute.Attribute{},
		},
		expected: 2,
	},
}

// WithAttribute()
func TestNew(t *testing.T) {
	for testName, testCase := range newCases {
		t.Run(testName, func(t *testing.T) {
			withFns := []optionalAttr{}
			for _, attr := range testCase.input {
				withFns = append(withFns, WithAttribute(attr))
			}
			node := New("Foo", withFns...)
			assert.Len(t, node.Attributes, testCase.expected)
		})
	}
}

// node.AddAttr()
func TestAddAttr(t *testing.T) {
	for testName, testCase := range newCases {
		t.Run(testName, func(t *testing.T) {
			node := New("Foo")
			for _, attr := range testCase.input {
				node.AddAttr(attr)
			}
			assert.Len(t, node.Attributes, testCase.expected)
		})
	}
}

// node.String()
var stringCases = map[string]struct {
	input    []attribute.Attribute
	expected string
}{
	"Zero attributes": {
		input:    []attribute.Attribute{},
		expected: "\"Foo\"\n",
	},
	"One attribute": {
		input: []attribute.Attribute{
			attribute.New("FooAttr", "FooValue"),
		},
		expected: "\"Foo\" [FooAttr=\"FooValue\"]\n",
	},
	"Two attributes": {
		input: []attribute.Attribute{
			attribute.New("FooAttr", "FooValue"),
			attribute.New("BarAttr", "BarValue"),
		},
		expected: "\"Foo\" [FooAttr=\"FooValue\"BarAttr=\"BarValue\"]\n",
	},
}

func TestString(t *testing.T) {
	for testName, testCase := range stringCases {
		t.Run(testName, func(t *testing.T) {
			node := New("Foo")
			for _, attr := range testCase.input {
				node.AddAttr(attr)
			}
			assert.Equal(t, testCase.expected, node.String())
		})
	}
}

func TestWithAttributes(t *testing.T) {
	attrs := []attribute.Attribute{
		attribute.New(
			"foo",
			"fooVal",
		),
		attribute.New(
			"bar",
			"barVal",
		),
		attribute.New(
			"baz",
			"bazVal",
		),
	}
	n := New("fooNode", WithAttributes(attrs))
	assert.Len(t, n.Attributes, 3)
}
