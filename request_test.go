package main

import "testing"

const endpoint = `https://www.googleapis.com/books/v1/volumes`
const testQuery = `https://www.googleapis.com/books/v1/volumes?q=fish+inauthors%3Aamy+lane`

func TestNewRequest(t *testing.T) {
	r := NewRequest()
	if r.String() != endpoint {
		t.Errorf("got %s, expected %s\n", r.String(), endpoint)
	}
}

func TestSearchRequest(t *testing.T) {
	r := testSearchRequest()
	if q := r.String(); q != testQuery {
		t.Errorf("got %s, expected %s\n", q, testQuery)
	}
}

func testSearchRequest() *Request {
	return NewRequest().
		Keywords("fish").
		Author("amy lane")
}
