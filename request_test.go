package main

import "testing"

const endpoint = `https://www.googleapis.com/books/v1/volumes`

func TestNewRequest(t *testing.T) {
	r := NewRequest()
	if r.String() != endpoint {
		t.Errorf("got %s, expected %s\n", r.String(), endpoint)
	}
}
