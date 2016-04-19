package main

import (
	"testing"
)

func TestSave(t *testing.T) {
	m := NewURL()
	result := m.Save("http://example.com")
	expect := "aaaaaaaaaa"

	if result != expect {
		t.Errorf("got %v\nwant %v", result, expect)
	}
}
