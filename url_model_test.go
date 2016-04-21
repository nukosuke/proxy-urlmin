package main

import (
	"testing"
)

func TestSave(t *testing.T) {
	m := NewURL(nil)
	result, err := m.Save("http://example.com")
	expect := "aaaaaaaaaa"

	if err != nil {
		t.Errorf("error occured in URL.Save()")
	} else if result != expect {
		t.Errorf("got %v\nwant %v", result, expect)
	}
}
