// Package attribute provides that base attribute for DOT entites.
//
// An attribute is a simple key value pair
package attribute

import "fmt"

// This will accept any Name for attribute but for a list of accepted values
// see the [graphviz documentation].
//
// [graphviz documentation]: https://graphviz.org/doc/info/attrs.html
type Attribute struct {
	Name  string
	Value any
}

// This is a little reduntant becuase the Attribute struct and its fields
// are exported but it helps to keep the API consistant
func New(name string, value any) Attribute {
	return Attribute{
		Name:  name,
		Value: value,
	}
}

// Some attributes are standalone on thier own line and this is a helper so
// developers don't need to handle the newline themselves.
func (a Attribute) Stringln() string {
	return fmt.Sprintf("%s\n", a.String())
}

// Simply return the name=value as a string
func (a Attribute) String() string {
	return fmt.Sprintf("%s=\"%v\"", a.Name, a.Value)
}
