package gbooks

import (
	"net/url"
)

const endVolums = `https://www.googleapis.com/books/v1/volumes`

type Request struct {
	Query   url.Values
	url     *url.URL
	kw      string
	authors string
	title   string
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
	r.kw = kw
	return r
}

func (r *Request) Author(kw string) *Request {
	r.authors = kw
	return r
}

func (r *Request) Title(kw string) *Request {
	r.title = kw
	return r
}

func (r *Request) parseQuery() string {
	var q string
	if r.kw != "" {
		q += r.kw
	}
	if r.authors != "" {
		q += " inauthors:"
		q += r.authors
	}
	if r.title != "" {
		q += " intitle:"
		q += r.title
	}
	return q
}

func (r *Request) String() string {
	if q := r.parseQuery(); q != "" {
		r.Query.Set("q", q)
	}
	r.url.RawQuery = r.Query.Encode()
	return r.url.String()
}
