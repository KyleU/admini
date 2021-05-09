package schematypes

import (
	"fmt"
	"github.com/kyleu/admini/app/util"
)

const KeyList = "list"

var _ Type = (*List)(nil)

type List struct {
	T *Wrapped `json:"t"`
}

func (x *List) Key() string {
	return KeyList
}

func (x *List) String() string {
	return fmt.Sprintf("[]%v", x.T.String())
}

func (x *List) Sortable() bool {
	return x.T.Sortable()
}

func (x *List) From(v interface{}) interface{} {
	switch t := v.(type) {
	case string:
		lt := util.SplitAndTrim(t, ",")
		return lt
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewList(t *Wrapped) *Wrapped {
	return Wrap(&List{T: t})
}
