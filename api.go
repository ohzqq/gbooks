package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	op "github.com/ohzqq/libopds2-go/opds2"
	"github.com/ohzqq/ur"
	"github.com/ohzqq/ur/pkg/book"
	gb "google.golang.org/api/books/v1"
)

const apiHost = `www.googleapis.com`
const apiPath = `/books/v1/volumes`
const gBooksVol = `https://www.googleapis.com/books/v1/volumes`
const gBooksBook = `https://www.googleapis.com/books/v1/volumes/`

var Client = &api{
	client: &http.Client{},
}

type api struct {
	api    *gb.Service
	client *http.Client
	Results
}

type Results struct {
	self  string
	Items []book.GBook `json:"items"`
	Total int          `json:"totalItems"`
}

func Search(q *ur.Query) ur.Feed {
	return Api(q)
}

func Api(q *ur.Query) *api {
	return Client.GetBooks(q)
}

func (client *api) Search(q *ur.Query) ur.Feed {
	return client.GetBooks(q)
}

func (p Results) Pubs() []op.Publication {
	var books []op.Publication
	for _, book := range p.Items {
		books = append(books, book.ToPub())
	}
	return books
}

func (s *api) GetBooks(q *ur.Query) (*api, error) {
	req := parseQuery(q)
	resp, err := s.client.Get(req.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//var vols Results
	err = json.Unmarshal(body, &s.Results)
	if err != nil {
		return nil, err
	}
	s.query = q
	s.self = req.String()

	return s, nil
}

func parseQuery(q *ur.Query) *url.URL {
	u := &url.URL{
		Host:   apiHost,
		Scheme: "https",
		Path:   apiPath,
	}
	vals := make(url.Values)
	if q.Search.Keywords != "" {
		vals.Set("q", q.Search.Keywords)
	}
	if q.Filter.Limit == 0 {
		vals.Set("maxResults", "40")
	}
	u.RawQuery = vals.Encode()
	return u
}

//func (s Scraper) Search(params map[string]string) book.Books {
//  var q []string

//  for k, v := range params {
//    switch k {
//    case "keywords":
//      q = append(q, v)
//    case "authors":
//      q = append(q, "inauthor:"+`"`+v+`"`)
//    case "title":
//      q = append(q, "intitle:"+`"`+v+`"`)
//    }
//  }

//  v := make(url.Values)
//  v.Set("q", strings.Join(q, " "))

//  return s.GetBooks(v.Encode())
//}
