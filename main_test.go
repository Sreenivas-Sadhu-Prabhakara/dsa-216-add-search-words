package main

import "testing"

func TestWordDictionary(t *testing.T) {
	d := NewWordDictionary()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")

	cases := []struct {
		pattern string
		want    bool
	}{
		{"pad", false},
		{"bad", true},
		{".ad", true},
		{"b..", true},
		{"...", true},
		{"....", false}, // '.' matches exactly one letter, never zero
		{"..", false},
		{"b.d", true},
		{".", false},
	}
	for _, c := range cases {
		if got := d.Search(c.pattern); got != c.want {
			t.Errorf("Search(%q) = %v, want %v", c.pattern, got, c.want)
		}
	}

	d.AddWord("a")
	if !d.Search(".") {
		t.Error("single dot should now match the single-letter word")
	}
	if d.Search("") {
		t.Error("empty pattern matches nothing stored")
	}
}
