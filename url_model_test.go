package main

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestSave(t *testing.T) {
	//TODO beforeにまとめる
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		t.Errorf("failed to connect redis")
	}

	m := NewURL(conn)
	result, err := m.Save("http://example.com")
	expect := "aaaaaaaaaa"

	if err != nil {
		t.Errorf("error occured in URL.Save()")
	} else if result != expect {
		t.Errorf("got %v\nwant %v", result, expect)
	}
}
