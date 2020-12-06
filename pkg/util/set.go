package util

type set map[interface{}]struct{}

type Set struct {
	s set
}

func NewSet(v ...interface{}) Set {
	s := Set{s: make(set)}
	for _, item := range v {
		s.Add(item)
	}
	return s
}

func (set *Set) Add(v interface{}) {
	set.s[v] = struct{}{}
}

func (set *Set) AddBytes(v []byte) {
	for _, b := range v {
		set.s[b] = struct{}{}
	}
}

func (set *Set) Cardinality() int {
	return len(set.s)
}

func (set *Set) Contains(v interface{}) bool {
	_, ok := set.s[v]
	return ok
}

func (set *Set) Intersect(other Set) Set {
	inter := NewSet()
	if set.Cardinality() < other.Cardinality() {
		for v := range set.s {
			if set.Contains(v) {
				inter.Add(v)
			}
		}
	} else {
		for v := range other.s {
			if set.Contains(v) {
				inter.Add(v)
			}
		}
	}
	return inter
}
