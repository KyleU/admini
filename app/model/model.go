package model

import (
	"sort"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/field"

	"github.com/kyleu/admini/app/util"
)

type Model struct {
	Key           string        `json:"key"`
	Title         string        `json:"title,omitempty"`
	Pkg           util.Pkg      `json:"pkg,omitempty"`
	Type          Type          `json:"type"`
	Interfaces    []string      `json:"interfaces,omitempty"`
	Fields        field.Fields  `json:"fields,omitempty"`
	Indexes       Indexes       `json:"indexes,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
	Metadata      *Metadata     `json:"metadata,omitempty"`
	pk            []string
}

func (m *Model) String() string {
	if len(m.Pkg) == 0 {
		return m.Key
	}
	return m.Pkg.ToPath(m.Key)
}

func (m *Model) Name() string {
	if m.Title == "" {
		return m.Key
	}
	return m.Title
}

func (m *Model) Description() string {
	return m.Type.String() + " model [" + m.Name() + "]"
}

func (m *Model) Path() util.Pkg {
	return m.Pkg.With(m.Key)
}

func (m *Model) PathString() string {
	return m.Pkg.ToPath(m.Key)
}

type Models []*Model

func (m Models) Get(pkg util.Pkg, key string) *Model {
	for _, x := range m {
		if x.Pkg.Equals(pkg) && x.Key == key {
			return x
		}
	}
	return nil
}

func (m Models) Sort() {
	sort.Slice(m, func(l int, r int) bool {
		return m[l].Key < m[r].Key
	})
}

func (m *Model) AddField(f *field.Field) error {
	if f == nil {
		return errors.New("nil field")
	}
	if _, v := m.Fields.Get(f.Key); v != nil {
		return errors.New("field [" + f.Key + "] already exists")
	}
	m.Fields = append(m.Fields, f)
	return nil
}

func (m *Model) AddIndex(i *Index) error {
	if i == nil {
		return errors.New("nil index")
	}
	if m.Indexes.Get(i.Key) != nil {
		return errors.New("index [" + i.Key + "] already exists")
	}
	m.Indexes = append(m.Indexes, i)
	return nil
}

func (m *Model) AddRelationship(r *Relationship) error {
	if r == nil {
		return errors.New("nil relation")
	}
	if m.Relationships.Get(r.Key) != nil {
		return errors.New("relation [" + r.Key + "] already exists")
	}
	m.Relationships = append(m.Relationships, r)
	return nil
}
