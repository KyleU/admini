package schematypes

const KeyDate = "date"

type Date struct{}

var _ Type = (*Date)(nil)

func (x *Date) Key() string {
	return KeyDate
}

func (x *Date) String() string {
	return x.Key()
}

func (x *Date) Sortable() bool {
	return true
}

func (x *Date) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewDate() *Wrapped {
	return Wrap(&Date{})
}
