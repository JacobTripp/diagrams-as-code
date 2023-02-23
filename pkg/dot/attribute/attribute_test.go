package attribute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inOutForNew struct {
	name  string
	value any
}

var testNewCases = map[string]struct {
	input    inOutForNew
	expected inOutForNew
}{
	"String": {
		input: inOutForNew{
			name:  "Foo",
			value: "Bar",
		},
		expected: inOutForNew{
			name:  "Foo",
			value: "Bar",
		},
	},
	"Int": {
		input: inOutForNew{
			name:  "Foo",
			value: 0,
		},
		expected: inOutForNew{
			name:  "Foo",
			value: 0,
		},
	},
	"Bool": {
		input: inOutForNew{
			name:  "Foo",
			value: true,
		},
		expected: inOutForNew{
			name:  "Foo",
			value: true,
		},
	},
}

func TestNew(t *testing.T) {
	for testName, testCase := range testNewCases {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			myAttr := New(testCase.input.name, testCase.input.value)
			assert.Equal(t, testCase.expected.name, myAttr.Name)
			assert.Equal(t, testCase.expected.value, myAttr.Value)
		})
	}
}

func TestAttributeStringln(t *testing.T) {
	t.Parallel()
	myAttr := New("Foo", "Bar")
	out := myAttr.Stringln()
	assert.Equal(t, "Foo=\"Bar\"\n", out)
}

func TestAttributeString(t *testing.T) {
	t.Parallel()
	myAttr := New("Foo", "Bar")
	out := myAttr.String()
	assert.Equal(t, "Foo=\"Bar\"", out)
}
