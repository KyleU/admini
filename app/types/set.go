// Code generated by projectforge.dev using code from the [types] module, which is under license [CC0]
package types

import (
	"fmt"
)

const KeySet = "set"

type Set struct {
	V *Wrapped `json:"v"`
}

var _ Type = (*Set)(nil)

func (x *Set) Key() string {
	return KeySet
}

func (x *Set) String() string {
	return fmt.Sprintf("%s[%s]", x.Key(), x.V.String())
}

func (x *Set) Sortable() bool {
	return x.V.Sortable()
}

func (x *Set) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func (x *Set) Default(string) interface{} {
	return emptyList
}

func NewSet() *Wrapped {
	return Wrap(&Set{})
}
