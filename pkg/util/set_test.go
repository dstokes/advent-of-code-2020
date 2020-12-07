package util

import "testing"

func TestNewSet(t *testing.T) {
	s := NewSet('a', 'b')
	if s.Contains('a') == false || s.Contains('b') == false {
		t.Error("NewSet didn't set values")
	}
}

func TestSetContains(t *testing.T) {
	s := make(set)
	s['a'] = struct{}{}
	set := Set{s: s}
	if set.Contains('a') == false {
		t.Error("Expected 'a' to exist")
	}
}

func TestSetAdd(t *testing.T) {
	set := NewSet()
	set.Add('a')
	if set.Contains('a') == false {
		t.Error("Expected map to have field 'a'")
	}
}

func TestSetIntersect(t *testing.T) {
	a := NewSet()
	b := NewSet()
	a.Add('a')
	a.Add('b')
	b.Add('a')
	c := a.Intersect(b)
	if c.Contains('a') == false || c.Contains('b') == true {
		t.Error("Intersect failed")
	}
}
