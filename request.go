package main

import "net/url"

const endVolums = `https://www.googleapis.com/books/v1/volumes`

type Request struct {
	Query url.Values
	url   *url.URL
}

func NewRequest() *Request {
	return &Request{
		Query: make(url.Values),
		url: &url.URL{
			Scheme: "https",
			Host:   "www.googleapis.com",
			Path:   "/books/v1/volumes",
		},
	}
}

func (r *Request) Keywords(kw string) *Request {
	r.Query.Add("q", kw)
	return r
}

func (r *Request) Author(kw string) *Request {
	r.Query.Add("q", "inauthor:"+kw)
	return r
}

func (r *Request) Tile(kw string) *Request {
	r.Query.Add("q", "intitle:"+kw)
	return r
}

func (r *Request) String() string {
	r.url.RawQuery = r.Query.Encode()
	return r.url.String()
}
