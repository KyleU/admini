package project

import (
	"github.com/kyleu/admini/app/project/action"
	"sort"
)

type Project struct {
	Key     string         `json:"-"`
	Title   string         `json:"title,omitempty"`
	Sources []string       `json:"sources,omitempty"`
	Actions action.Actions `json:"actions,omitempty"`
}

func (p *Project) Name() string {
	if p.Title == "" {
		return p.Key
	}
	return p.Title
}

type Projects []*Project

func (p Projects) Get(key string) *Project {
	for _, x := range p {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func (p Projects) Replace(n *Project) {
	for idx, x := range p {
		if x.Key == n.Key {
			p[idx] = n
			return
		}
	}
	p = append(p, n)
}

func (p Projects) Sort() {
	sort.Slice(p, func(l int, r int) bool {
		return p[l].Key < p[r].Key
	})
}
